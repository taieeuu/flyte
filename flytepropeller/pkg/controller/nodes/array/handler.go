package array

import (
	"context"
	"fmt"
	"math"
	"slices"
	"strconv"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	idlcore "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/ioutils"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/plugins/array/errorcollector"
	"github.com/flyteorg/flyte/flytepropeller/events"
	eventsErr "github.com/flyteorg/flyte/flytepropeller/events/errors"
	"github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	"github.com/flyteorg/flyte/flytepropeller/pkg/compiler/validators"
	"github.com/flyteorg/flyte/flytepropeller/pkg/controller/config"
	"github.com/flyteorg/flyte/flytepropeller/pkg/controller/executors"
	"github.com/flyteorg/flyte/flytepropeller/pkg/controller/nodes/common"
	"github.com/flyteorg/flyte/flytepropeller/pkg/controller/nodes/errors"
	"github.com/flyteorg/flyte/flytepropeller/pkg/controller/nodes/handler"
	"github.com/flyteorg/flyte/flytepropeller/pkg/controller/nodes/interfaces"
	"github.com/flyteorg/flyte/flytepropeller/pkg/controller/nodes/task/k8s"
	"github.com/flyteorg/flyte/flytestdlib/bitarray"
	stdConfig "github.com/flyteorg/flyte/flytestdlib/config"
	"github.com/flyteorg/flyte/flytestdlib/logger"
	"github.com/flyteorg/flyte/flytestdlib/promutils"
	"github.com/flyteorg/flyte/flytestdlib/storage"
)

var (
	nilLiteral = &idlcore.Literal{
		Value: &idlcore.Literal_Scalar{
			Scalar: &idlcore.Scalar{
				Value: &idlcore.Scalar_NoneType{
					NoneType: &idlcore.Void{},
				},
			},
		},
	}
)

//go:generate mockery --all --case=underscore --with-expecter

// arrayNodeHandler is a handle implementation for processing array nodes
type arrayNodeHandler struct {
	eventConfig                 *config.EventConfig
	literalOffloadingConfig     config.LiteralOffloadingConfig
	gatherOutputsRequestChannel chan *gatherOutputsRequest
	metrics                     metrics
	nodeExecutionRequestChannel chan *nodeExecutionRequest
	nodeExecutor                interfaces.Node
	pluginStateBytesNotStarted  []byte
	pluginStateBytesStarted     []byte
}

// metrics encapsulates the prometheus metrics for this handler
type metrics struct {
	scope promutils.Scope
}

// newMetrics initializes a new metrics struct
func newMetrics(scope promutils.Scope) metrics {
	return metrics{
		scope: scope,
	}
}

// Abort stops the array node defined in the NodeExecutionContext
func (a *arrayNodeHandler) Abort(ctx context.Context, nCtx interfaces.NodeExecutionContext, reason string) error {
	arrayNode := nCtx.Node().GetArrayNode()
	arrayNodeState := nCtx.NodeStateReader().GetArrayNodeState()

	eventRecorder := newArrayEventRecorder(nCtx.EventsRecorder())
	messageCollector := errorcollector.NewErrorMessageCollector()

	taskPhase := idlcore.TaskExecution_ABORTED
	if arrayNodeState.Phase == v1alpha1.ArrayNodePhaseFailing {
		taskPhase = idlcore.TaskExecution_FAILED
	}
	switch arrayNodeState.Phase {
	case v1alpha1.ArrayNodePhaseExecuting, v1alpha1.ArrayNodePhaseFailing:
		for i, nodePhaseUint64 := range arrayNodeState.SubNodePhases.GetItems() {
			nodePhase := v1alpha1.NodePhase(nodePhaseUint64) // #nosec G115

			// do not process nodes that have not started or are in a terminal state
			if nodePhase == v1alpha1.NodePhaseNotYetStarted || isTerminalNodePhase(nodePhase) {
				continue
			}

			// create array contexts
			arrayNodeExecutor, arrayExecutionContext, arrayDAGStructure, arrayNodeLookup, subNodeSpec, _, err :=
				a.buildArrayNodeContext(ctx, nCtx, &arrayNodeState, arrayNode, i, eventRecorder)
			if err != nil {
				return err
			}

			// abort subNode
			err = arrayNodeExecutor.AbortHandler(ctx, arrayExecutionContext, arrayDAGStructure, arrayNodeLookup, subNodeSpec, reason)
			if err != nil {
				messageCollector.Collect(i, err.Error())
			} else {
				// record events transitioning subNodes to aborted
				retryAttempt := uint32(arrayNodeState.SubNodeRetryAttempts.GetItem(i)) // #nosec G115

				if err := sendEvents(ctx, nCtx, i, retryAttempt, idlcore.NodeExecution_ABORTED, idlcore.TaskExecution_ABORTED, eventRecorder, a.eventConfig); err != nil {
					logger.Warnf(ctx, "failed to record ArrayNode events: %v", err)
				}

				if err := eventRecorder.process(ctx, nCtx, i, retryAttempt); err != nil {
					logger.Warnf(ctx, "failed to record ArrayNode events: %v", err)
				}
			}
		}
	}

	if messageCollector.Length() > 0 {
		return fmt.Errorf(messageCollector.Summary(events.MaxErrorMessageLength)) //nolint:govet,staticcheck
	}

	// update state for subNodes
	if err := eventRecorder.finalize(ctx, nCtx, taskPhase, 0, a.eventConfig); err != nil {
		// a task event with abort phase is already emitted when handling ArrayNodePhaseFailing
		if !eventsErr.IsAlreadyExists(err) {
			logger.Errorf(ctx, "ArrayNode event recording failed: [%s]", err.Error())
			return err
		}
	}

	return nil
}

// Finalize completes the array node defined in the NodeExecutionContext
func (a *arrayNodeHandler) Finalize(ctx context.Context, nCtx interfaces.NodeExecutionContext) error {
	arrayNode := nCtx.Node().GetArrayNode()
	arrayNodeState := nCtx.NodeStateReader().GetArrayNodeState()

	eventRecorder := newArrayEventRecorder(nCtx.EventsRecorder())
	messageCollector := errorcollector.NewErrorMessageCollector()
	switch arrayNodeState.Phase {
	case v1alpha1.ArrayNodePhaseExecuting, v1alpha1.ArrayNodePhaseFailing, v1alpha1.ArrayNodePhaseSucceeding:
		for i, nodePhaseUint64 := range arrayNodeState.SubNodePhases.GetItems() {
			nodePhase := v1alpha1.NodePhase(nodePhaseUint64) // #nosec G115

			// do not process nodes that have not started or are in a terminal state
			if nodePhase == v1alpha1.NodePhaseNotYetStarted || isTerminalNodePhase(nodePhase) {
				continue
			}

			// create array contexts
			arrayNodeExecutor, arrayExecutionContext, arrayDAGStructure, arrayNodeLookup, subNodeSpec, _, err :=
				a.buildArrayNodeContext(ctx, nCtx, &arrayNodeState, arrayNode, i, eventRecorder)
			if err != nil {
				return err
			}

			// finalize subNode
			err = arrayNodeExecutor.FinalizeHandler(ctx, arrayExecutionContext, arrayDAGStructure, arrayNodeLookup, subNodeSpec)
			if err != nil {
				messageCollector.Collect(i, err.Error())
			}
		}
	}

	if messageCollector.Length() > 0 {
		return fmt.Errorf(messageCollector.Summary(events.MaxErrorMessageLength)) //nolint:govet,staticcheck
	}

	return nil
}

// FinalizeRequired defines whether or not this handler requires finalize to be called on node
// completion
func (a *arrayNodeHandler) FinalizeRequired() bool {
	// must return true because we can't determine if finalize is required for the subNode
	return true
}

// Handle is responsible for transitioning and reporting node state to complete the node defined
// by the NodeExecutionContext
func (a *arrayNodeHandler) Handle(ctx context.Context, nCtx interfaces.NodeExecutionContext) (handler.Transition, error) {
	arrayNode := nCtx.Node().GetArrayNode()
	arrayNodeState := nCtx.NodeStateReader().GetArrayNodeState()
	currentArrayNodePhase := arrayNodeState.Phase

	incrementTaskPhaseVersion := false
	eventRecorder := newArrayEventRecorder(nCtx.EventsRecorder())

	switch currentArrayNodePhase {
	case v1alpha1.ArrayNodePhaseNone:
		// identify and validate array node input value lengths
		literalMap, err := nCtx.InputReader().Get(ctx)
		if err != nil {
			return handler.UnknownTransition, err
		}

		size := -1

		for key, variable := range literalMap.GetLiterals() {
			literalType := validators.LiteralTypeForLiteral(variable)
			err := validators.ValidateLiteralType(literalType)

			if slices.Contains(arrayNode.GetBoundInputs(), key) {
				continue
			}

			if err != nil {
				errMsg := fmt.Sprintf("Failed to validate literal type for [%s] with err: %s", key, err)
				return handler.DoTransition(handler.TransitionTypeEphemeral,
					handler.PhaseInfoFailure(idlcore.ExecutionError_USER, errors.IDLNotFoundErr, errMsg, nil),
				), nil
			}
			if variable.GetOffloadedMetadata() != nil {
				// variable will be overwritten with the contents of the offloaded data which contains the actual large literal.
				// We need this for the map task to be able to create the subNodeSpec
				err := common.ReadLargeLiteral(ctx, nCtx.DataStore(), variable)
				if err != nil {
					return handler.DoTransition(handler.TransitionTypeEphemeral,
						handler.PhaseInfoFailure(idlcore.ExecutionError_SYSTEM, errors.RuntimeExecutionError, "couldn't read the offloaded literal", nil),
					), nil
				}
			}
			switch literalType.GetType().(type) {
			case *idlcore.LiteralType_CollectionType:
				collectionLength := len(variable.GetCollection().GetLiterals())
				if size == -1 {
					size = collectionLength
				} else if size != collectionLength {
					return handler.DoTransition(handler.TransitionTypeEphemeral,
						handler.PhaseInfoFailure(idlcore.ExecutionError_USER, errors.InvalidArrayLength,
							fmt.Sprintf("input arrays have different lengths: expecting '%d' found '%d'", size, collectionLength), nil),
					), nil
				}
			}
		}

		if size == -1 {
			// handles case where all inputs are bound
			if len(arrayNode.GetBoundInputs()) == len(literalMap.GetLiterals()) {
				size = 1
			} else {
				return handler.DoTransition(handler.TransitionTypeEphemeral,
					handler.PhaseInfoFailure(idlcore.ExecutionError_USER, errors.InvalidArrayLength, "no input array provided", nil),
				), nil
			}
		}

		// initialize ArrayNode state
		maxSystemFailuresValue := int(config.GetConfig().NodeConfig.MaxNodeRetriesOnSystemFailures)
		maxAttemptsValue := int(config.GetConfig().NodeConfig.DefaultMaxAttempts)
		if nCtx.Node().GetRetryStrategy() != nil && nCtx.Node().GetRetryStrategy().MinAttempts != nil && *nCtx.Node().GetRetryStrategy().MinAttempts != 1 {
			maxAttemptsValue = *nCtx.Node().GetRetryStrategy().MinAttempts
		}

		if config.GetConfig().NodeConfig.IgnoreRetryCause {
			maxSystemFailuresValue = maxAttemptsValue
		} else {
			maxAttemptsValue += maxSystemFailuresValue
		}

		for _, item := range []struct {
			arrayReference *bitarray.CompactArray
			maxValue       int
		}{
			// we use NodePhaseRecovered for the `maxValue` of `SubNodePhases` because `Phase` is
			// defined as an `iota` so it is impossible to programmatically get largest value
			{arrayReference: &arrayNodeState.SubNodePhases, maxValue: int(v1alpha1.NodePhaseRecovered)},
			{arrayReference: &arrayNodeState.SubNodeTaskPhases, maxValue: len(core.Phases) - 1},
			{arrayReference: &arrayNodeState.SubNodeRetryAttempts, maxValue: maxAttemptsValue},
			{arrayReference: &arrayNodeState.SubNodeSystemFailures, maxValue: maxSystemFailuresValue},
			{arrayReference: &arrayNodeState.SubNodeDeltaTimestamps, maxValue: int(config.GetConfig().ArrayNode.MaxDeltaTimestamp.Seconds())},
		} {

			*item.arrayReference, err = bitarray.NewCompactArray(uint(size), bitarray.Item(item.maxValue)) // #nosec G115
			if err != nil {
				return handler.UnknownTransition, err
			}
		}

		// initialize subNode status by faking events
		for i := 0; i < size; i++ {
			if err := sendEvents(ctx, nCtx, i, 0, idlcore.NodeExecution_QUEUED, idlcore.TaskExecution_UNDEFINED, eventRecorder, a.eventConfig); err != nil {
				logger.Warnf(ctx, "failed to record ArrayNode events: %v", err)
			}

			if err := eventRecorder.process(ctx, nCtx, i, 0); err != nil {
				logger.Warnf(ctx, "failed to record ArrayNode events: %v", err)
			}
		}

		// transition ArrayNode to `ArrayNodePhaseExecuting`
		arrayNodeState.Phase = v1alpha1.ArrayNodePhaseExecuting
	case v1alpha1.ArrayNodePhaseExecuting:
		// process array node subNodes
		remainingWorkflowParallelism := int(nCtx.ExecutionContext().GetExecutionConfig().MaxParallelism - nCtx.ExecutionContext().CurrentParallelism())
		incrementWorkflowParallelism, maxParallelism := inferParallelism(ctx, arrayNode.GetParallelism(),
			config.GetConfig().ArrayNode.DefaultParallelismBehavior, remainingWorkflowParallelism, len(arrayNodeState.SubNodePhases.GetItems()))

		nodeExecutionRequests := make([]*nodeExecutionRequest, 0, maxParallelism)
		currentParallelism := 0
		for i, nodePhaseUint64 := range arrayNodeState.SubNodePhases.GetItems() {
			if currentParallelism >= maxParallelism {
				break
			}

			nodePhase := v1alpha1.NodePhase(nodePhaseUint64)              // #nosec G115
			taskPhase := int(arrayNodeState.SubNodeTaskPhases.GetItem(i)) // #nosec G115

			// do not process nodes in terminal state
			if isTerminalNodePhase(nodePhase) {
				continue
			}

			// create array contexts
			subNodeEventRecorder := newArrayEventRecorder(nCtx.EventsRecorder())
			arrayNodeExecutor, arrayExecutionContext, arrayDAGStructure, arrayNodeLookup, subNodeSpec, subNodeStatus, err :=
				a.buildArrayNodeContext(ctx, nCtx, &arrayNodeState, arrayNode, i, subNodeEventRecorder)
			if err != nil {
				return handler.UnknownTransition, err
			}

			nodeExecutionRequest := &nodeExecutionRequest{
				ctx:                ctx,
				index:              i,
				nodePhase:          nodePhase,
				taskPhase:          taskPhase,
				nodeExecutor:       arrayNodeExecutor,
				executionContext:   arrayExecutionContext,
				dagStructure:       arrayDAGStructure,
				nodeLookup:         arrayNodeLookup,
				subNodeSpec:        subNodeSpec,
				subNodeStatus:      subNodeStatus,
				arrayEventRecorder: subNodeEventRecorder,
				responseChannel: make(chan struct {
					interfaces.NodeStatus
					error
				}, 1),
			}

			nodeExecutionRequests = append(nodeExecutionRequests, nodeExecutionRequest)
			a.nodeExecutionRequestChannel <- nodeExecutionRequest

			// TODO - this is a naive implementation of parallelism, if we want to support more
			// complex subNodes (ie. dynamics / subworkflows) we need to revisit this so that
			// parallelism is handled during subNode evaluations + avoid deadlocks
			if incrementWorkflowParallelism {
				nCtx.ExecutionContext().IncrementParallelism()
			}
			currentParallelism++
		}

		workerErrorCollector := errorcollector.NewErrorMessageCollector()
		subNodeFailureCollector := errorcollector.NewErrorMessageCollector()
		for i, nodeExecutionRequest := range nodeExecutionRequests {
			nodeExecutionResponse := <-nodeExecutionRequest.responseChannel
			if nodeExecutionResponse.error != nil {
				workerErrorCollector.Collect(i, nodeExecutionResponse.error.Error())
				continue
			}

			index := nodeExecutionRequest.index
			subNodeStatus := nodeExecutionRequest.subNodeStatus

			// capture subNode error if exists
			if nodeExecutionRequest.subNodeStatus.Error != nil {
				subNodeFailureCollector.Collect(index, subNodeStatus.Error.Message)
			}

			// process events by copying from internal event recorder
			if arrayEventRecorder, ok := nodeExecutionRequest.arrayEventRecorder.(*externalResourcesEventRecorder); ok {
				for _, event := range arrayEventRecorder.taskEvents {
					if err := eventRecorder.RecordTaskEvent(ctx, event, a.eventConfig); err != nil {
						return handler.UnknownTransition, err
					}
				}
				for _, event := range arrayEventRecorder.nodeEvents {
					if err := eventRecorder.RecordNodeEvent(ctx, event, a.eventConfig); err != nil {
						return handler.UnknownTransition, err
					}
				}
			}

			if err := eventRecorder.process(ctx, nCtx, index, subNodeStatus.GetAttempts()); err != nil {
				return handler.UnknownTransition, err
			}

			// update subNode state
			arrayNodeState.SubNodePhases.SetItem(index, uint64(subNodeStatus.GetPhase())) // #nosec G115
			if subNodeStatus.GetTaskNodeStatus() == nil {
				// resetting task phase because during retries we clear the GetTaskNodeStatus
				arrayNodeState.SubNodeTaskPhases.SetItem(index, uint64(0))
			} else {
				arrayNodeState.SubNodeTaskPhases.SetItem(index, uint64(subNodeStatus.GetTaskNodeStatus().GetPhase())) // #nosec G115
			}
			arrayNodeState.SubNodeRetryAttempts.SetItem(index, uint64(subNodeStatus.GetAttempts()))
			arrayNodeState.SubNodeSystemFailures.SetItem(index, uint64(subNodeStatus.GetSystemFailures()))

			if arrayNodeState.SubNodeDeltaTimestamps.BitSet != nil {
				startedAt := nCtx.NodeStatus().GetLastAttemptStartedAt()
				subNodeStartedAt := subNodeStatus.GetLastAttemptStartedAt()
				if subNodeStartedAt == nil {
					// subNodeStartedAt == nil indicates either (1) node has not started or (2) node status has
					// been reset (ex. retryable failure). in both cases we set the delta timestamp to 0
					arrayNodeState.SubNodeDeltaTimestamps.SetItem(index, 0)
				} else if startedAt != nil && arrayNodeState.SubNodeDeltaTimestamps.GetItem(index) == 0 {
					// otherwise if `SubNodeDeltaTimestamps` is unset, we compute the delta and set it
					deltaDuration := uint64(subNodeStartedAt.Time.Sub(startedAt.Time).Seconds())
					arrayNodeState.SubNodeDeltaTimestamps.SetItem(index, deltaDuration)
				}
			}

			// increment task phase version if subNode phase or task phase changed
			if subNodeStatus.GetPhase() != nodeExecutionRequest.nodePhase || subNodeStatus.GetTaskNodeStatus().GetPhase() != nodeExecutionRequest.taskPhase {
				incrementTaskPhaseVersion = true
			}
		}

		// if any workers failed then return the error
		if workerErrorCollector.Length() > 0 {
			return handler.UnknownTransition, fmt.Errorf("worker error(s) encountered: %s", workerErrorCollector.Summary(events.MaxErrorMessageLength))
		}

		// process phases of subNodes to determine overall `ArrayNode` phase
		successCount := 0
		failedCount := 0
		failingCount := 0
		runningCount := 0
		for _, nodePhaseUint64 := range arrayNodeState.SubNodePhases.GetItems() {
			nodePhase := v1alpha1.NodePhase(nodePhaseUint64) // #nosec G115
			switch nodePhase {
			case v1alpha1.NodePhaseSucceeded, v1alpha1.NodePhaseRecovered, v1alpha1.NodePhaseSkipped:
				successCount++
			case v1alpha1.NodePhaseFailing:
				failingCount++
			case v1alpha1.NodePhaseFailed, v1alpha1.NodePhaseTimedOut:
				failedCount++
			default:
				runningCount++
			}
		}

		// calculate minimum number of successes to succeed the ArrayNode
		minSuccesses := len(arrayNodeState.SubNodePhases.GetItems())
		if arrayNode.GetMinSuccesses() != nil {
			minSuccesses = int(*arrayNode.GetMinSuccesses())
		} else if minSuccessRatio := arrayNode.GetMinSuccessRatio(); minSuccessRatio != nil {
			minSuccesses = int(math.Ceil(float64(*minSuccessRatio) * float64(minSuccesses)))
		}

		// if there is a failing node set the error message if it has not been previous set
		if failingCount > 0 && arrayNodeState.Error == nil {
			arrayNodeState.Error = &idlcore.ExecutionError{
				Message: subNodeFailureCollector.Summary(events.MaxErrorMessageLength),
			}
		}

		if len(arrayNodeState.SubNodePhases.GetItems())-failedCount < minSuccesses {
			// no chance to reach the minimum number of successes
			arrayNodeState.Phase = v1alpha1.ArrayNodePhaseFailing
		} else if successCount >= minSuccesses && runningCount == 0 {
			// wait until all tasks have completed before declaring success
			arrayNodeState.Phase = v1alpha1.ArrayNodePhaseSucceeding
		}

		// if incrementWorkflowParallelism is not set then we need to increment the parallelism by one
		// to indicate that the overall ArrayNode is still running
		if !incrementWorkflowParallelism && arrayNodeState.Phase == v1alpha1.ArrayNodePhaseExecuting {
			nCtx.ExecutionContext().IncrementParallelism()
		}
	case v1alpha1.ArrayNodePhaseFailing:
		// note: sub node eventing handled during Abort
		if err := a.Abort(ctx, nCtx, "ArrayNodeFailing"); err != nil {
			return handler.UnknownTransition, err
		}

		// fail with reported error if one exists
		if arrayNodeState.Error != nil {
			return handler.DoTransition(handler.TransitionTypeEphemeral, handler.PhaseInfoFailureErr(arrayNodeState.Error, nil)), nil
		}

		return handler.DoTransition(handler.TransitionTypeEphemeral, handler.PhaseInfoFailure(
			idlcore.ExecutionError_UNKNOWN,
			"ArrayNodeFailing",
			"Unknown reason",
			nil,
		)), nil
	case v1alpha1.ArrayNodePhaseSucceeding:
		gatherOutputsRequests := make([]*gatherOutputsRequest, 0, len(arrayNodeState.SubNodePhases.GetItems()))
		for i, nodePhaseUint64 := range arrayNodeState.SubNodePhases.GetItems() {
			nodePhase := v1alpha1.NodePhase(nodePhaseUint64) // #nosec G115
			gatherOutputsRequest := &gatherOutputsRequest{
				ctx: ctx,
				responseChannel: make(chan struct {
					literalMap map[string]*idlcore.Literal
					error
				}, 1),
			}

			if nodePhase != v1alpha1.NodePhaseSucceeded {
				// retrieve output variables from task template
				outputLiterals := make(map[string]*idlcore.Literal)
				task, err := nCtx.ExecutionContext().GetTask(*arrayNode.GetSubNodeSpec().TaskRef)
				if err != nil {
					// Should never happen
					gatherOutputsRequest.responseChannel <- struct {
						literalMap map[string]*idlcore.Literal
						error
					}{nil, err}
					continue
				}

				if task.CoreTask() != nil && task.CoreTask().GetInterface() != nil && task.CoreTask().GetInterface().GetOutputs() != nil {
					for name := range task.CoreTask().GetInterface().GetOutputs().GetVariables() {
						outputLiterals[name] = nilLiteral
					}
				}

				gatherOutputsRequest.responseChannel <- struct {
					literalMap map[string]*idlcore.Literal
					error
				}{outputLiterals, nil}
			} else {
				// initialize subNode reader
				currentAttempt := int(arrayNodeState.SubNodeRetryAttempts.GetItem(i)) // #nosec G115
				subDataDir, subOutputDir, err := constructOutputReferences(ctx, nCtx,
					strconv.Itoa(i), strconv.Itoa(currentAttempt))
				if err != nil {
					gatherOutputsRequest.responseChannel <- struct {
						literalMap map[string]*idlcore.Literal
						error
					}{nil, err}
					continue
				}

				// checkpoint paths are not computed here because this function is only called when writing
				// existing cached outputs. if this functionality changes this will need to be revisited.
				outputPaths := ioutils.NewCheckpointRemoteFilePaths(ctx, nCtx.DataStore(), subOutputDir, ioutils.NewRawOutputPaths(ctx, subDataDir), "")
				reader := ioutils.NewRemoteFileOutputReader(ctx, nCtx.DataStore(), outputPaths, 0)

				gatherOutputsRequest.reader = &reader
				a.gatherOutputsRequestChannel <- gatherOutputsRequest
			}

			gatherOutputsRequests = append(gatherOutputsRequests, gatherOutputsRequest)
		}

		// attempt best effort at initializing outputLiterals with output variable names. currently
		// only TaskNode and WorkflowNode contain node interfaces.
		outputLiterals := make(map[string]*idlcore.Literal)
		switch arrayNode.GetSubNodeSpec().GetKind() {
		case v1alpha1.NodeKindTask:
			taskID := *arrayNode.GetSubNodeSpec().TaskRef
			taskNode, err := nCtx.ExecutionContext().GetTask(taskID)
			if err != nil {
				return handler.DoTransition(handler.TransitionTypeEphemeral, handler.PhaseInfoFailure(idlcore.ExecutionError_SYSTEM,
					errors.BadSpecificationError, fmt.Sprintf("failed to find ArrayNode subNode task with id: '%s'", taskID), nil)), nil
			}

			if outputs := taskNode.CoreTask().GetInterface().GetOutputs(); outputs != nil {
				for name := range outputs.GetVariables() {
					outputLiteral := &idlcore.Literal{
						Value: &idlcore.Literal_Collection{
							Collection: &idlcore.LiteralCollection{
								Literals: make([]*idlcore.Literal, 0, len(arrayNodeState.SubNodePhases.GetItems())),
							},
						},
					}

					outputLiterals[name] = outputLiteral
				}
			}
		case v1alpha1.NodeKindWorkflow:
			// TODO - to support launchplans we will need to process the output interface variables here
			fallthrough
		default:
			logger.Warnf(ctx, "ArrayNode does not support pre-populating outputLiteral collections for node kind '%s'", arrayNode.GetSubNodeSpec().GetKind())
		}

		workerErrorCollector := errorcollector.NewErrorMessageCollector()
		for i, gatherOutputsRequest := range gatherOutputsRequests {
			outputResponse := <-gatherOutputsRequest.responseChannel
			if outputResponse.error != nil {
				workerErrorCollector.Collect(i, outputResponse.error.Error())
				continue
			}

			// append literal for all output variables
			for name, literal := range outputResponse.literalMap {
				appendLiteral(name, literal, outputLiterals, len(arrayNodeState.SubNodePhases.GetItems()))
			}
		}

		// if any workers failed then return the error
		if workerErrorCollector.Length() > 0 {
			return handler.UnknownTransition, fmt.Errorf("worker error(s) encountered: %s", workerErrorCollector.Summary(events.MaxErrorMessageLength))
		}

		// only offload literal if config is enabled for this feature.
		if a.literalOffloadingConfig.Enabled {
			for outputLiteralKey, outputLiteral := range outputLiterals {
				// if the size of the output Literal is > threshold then we write the literal to the offloaded store and populate the literal with its zero value and update the offloaded url
				// use the OffloadLargeLiteralKey to create  {OffloadLargeLiteralKey}_offloaded_metadata.pb file in the datastore.
				// Update the url in the outputLiteral with the offloaded url and also update the size of the literal.
				offloadedOutputFile := v1alpha1.GetOutputsLiteralMetadataFile(outputLiteralKey, nCtx.NodeStatus().GetOutputDir())
				if err := common.OffloadLargeLiteral(ctx, nCtx.DataStore(), offloadedOutputFile, outputLiteral, a.literalOffloadingConfig); err != nil {
					return handler.UnknownTransition, err
				}
			}
		}
		outputLiteralMap := &idlcore.LiteralMap{
			Literals: outputLiterals,
		}

		outputFile := v1alpha1.GetOutputsFile(nCtx.NodeStatus().GetOutputDir())
		if err := nCtx.DataStore().WriteProtobuf(ctx, outputFile, storage.Options{}, outputLiteralMap); err != nil {
			return handler.UnknownTransition, err
		}

		// ensure task_execution set to succeeded
		if err := eventRecorder.finalize(ctx, nCtx, idlcore.TaskExecution_SUCCEEDED, 0, a.eventConfig); err != nil {
			if !eventsErr.IsAlreadyExists(err) {
				logger.Errorf(ctx, "ArrayNode event recording failed: [%s]", err.Error())
				return handler.UnknownTransition, err
			}
		}

		return handler.DoTransition(handler.TransitionTypeEphemeral, handler.PhaseInfoSuccess(
			&handler.ExecutionInfo{
				OutputInfo: &handler.OutputInfo{
					OutputURI: outputFile,
				},
			},
		)), nil
	default:
		return handler.UnknownTransition, errors.Errorf(errors.IllegalStateError, nCtx.NodeID(), "invalid ArrayNode phase %+v", arrayNodeState.Phase)
	}

	// if there were changes to subNode status then the eventRecorder will require finalizing to
	// report to admin through a TaskExecutionEvent.
	if eventRecorder.finalizeRequired(ctx) {
		// determine task phase from ArrayNodePhase
		taskPhase := idlcore.TaskExecution_UNDEFINED
		switch currentArrayNodePhase {
		case v1alpha1.ArrayNodePhaseNone:
			taskPhase = idlcore.TaskExecution_QUEUED
		case v1alpha1.ArrayNodePhaseExecuting:
			taskPhase = idlcore.TaskExecution_RUNNING
		case v1alpha1.ArrayNodePhaseSucceeding:
			taskPhase = idlcore.TaskExecution_SUCCEEDED
		case v1alpha1.ArrayNodePhaseFailing:
			taskPhase = idlcore.TaskExecution_FAILED
		}

		// increment taskPhaseVersion if we detect any changes in subNode state.
		if incrementTaskPhaseVersion {
			arrayNodeState.TaskPhaseVersion++
		}

		maxRetries := config.GetConfig().ArrayNode.MaxTaskPhaseVersionAttempts
		retries := 0
		for retries <= maxRetries {
			err := eventRecorder.finalize(ctx, nCtx, taskPhase, arrayNodeState.TaskPhaseVersion, a.eventConfig)

			if err == nil {
				break
			}

			// Handle potential race condition if FlyteWorkflow CRD fails to get synced
			if eventsErr.IsAlreadyExists(err) {
				if !incrementTaskPhaseVersion {
					break
				}
				logger.Warnf(ctx, "Event version already exists, bumping version and retrying (%d/%d): [%s]", retries+1, maxRetries, err.Error())
				arrayNodeState.TaskPhaseVersion++
			} else {
				logger.Errorf(ctx, "ArrayNode event recording failed: [%s]", err.Error())
				return handler.UnknownTransition, err
			}

			retries++
			if retries > maxRetries {
				logger.Errorf(ctx, "ArrayNode event recording failed after %d retries: [%s]", maxRetries, err.Error())
				return handler.UnknownTransition, err
			}
		}

		// if the ArrayNode phase has changed we need to reset the taskPhaseVersion to 0
		if currentArrayNodePhase != arrayNodeState.Phase {
			arrayNodeState.TaskPhaseVersion = 0
		}
	}

	// update array node status
	if err := nCtx.NodeStateWriter().PutArrayNodeState(arrayNodeState); err != nil {
		logger.Errorf(ctx, "failed to store ArrayNode state with err [%s]", err.Error())
		return handler.UnknownTransition, err
	}

	return handler.DoTransition(handler.TransitionTypeEphemeral, handler.PhaseInfoRunning(&handler.ExecutionInfo{})), nil
}

// Setup handles any initialization requirements for this handler
func (a *arrayNodeHandler) Setup(_ context.Context, _ interfaces.SetupContext) error {
	// start workers
	for i := 0; i < config.GetConfig().NodeExecutionWorkerCount; i++ {
		worker := worker{
			gatherOutputsRequestChannel: a.gatherOutputsRequestChannel,
			nodeExecutionRequestChannel: a.nodeExecutionRequestChannel,
		}

		go func() {
			worker.run()
		}()
	}

	return nil
}

// New initializes a new arrayNodeHandler
func New(nodeExecutor interfaces.Node, eventConfig *config.EventConfig, literalOffloadingConfig config.LiteralOffloadingConfig, scope promutils.Scope) (interfaces.NodeHandler, error) {
	// create k8s PluginState byte mocks to reuse instead of creating for each subNode evaluation
	pluginStateBytesNotStarted, err := bytesFromK8sPluginState(k8s.PluginState{Phase: k8s.PluginPhaseNotStarted})
	if err != nil {
		return nil, err
	}

	pluginStateBytesStarted, err := bytesFromK8sPluginState(k8s.PluginState{Phase: k8s.PluginPhaseStarted})
	if err != nil {
		return nil, err
	}

	eventConfigCopy, err := stdConfig.DeepCopyConfig(eventConfig)
	if err != nil {
		return nil, err
	}

	deepCopiedEventConfig, ok := eventConfigCopy.(*config.EventConfig)
	if !ok {
		return nil, fmt.Errorf("deep copy error: expected *config.EventConfig, but got %T", eventConfigCopy)
	}

	deepCopiedEventConfig.ErrorOnAlreadyExists = true

	arrayScope := scope.NewSubScope("array")
	return &arrayNodeHandler{
		eventConfig:                 deepCopiedEventConfig,
		literalOffloadingConfig:     literalOffloadingConfig,
		gatherOutputsRequestChannel: make(chan *gatherOutputsRequest),
		metrics:                     newMetrics(arrayScope),
		nodeExecutionRequestChannel: make(chan *nodeExecutionRequest),
		nodeExecutor:                nodeExecutor,
		pluginStateBytesNotStarted:  pluginStateBytesNotStarted,
		pluginStateBytesStarted:     pluginStateBytesStarted,
	}, nil
}

// buildArrayNodeContext creates a custom environment to execute the ArrayNode subnode. This is uniquely required for
// the arrayNodeHandler because we require the same node execution entrypoint (ie. recursiveNodeExecutor.RecursiveNodeHandler)
// but need many different execution details, for example setting input values as a singular item rather than a collection,
// injecting environment variables for flytekit maptask execution, aggregating eventing so that rather than tracking state for
// each subnode individually it sends a single event for the whole ArrayNode, and many more.
func (a *arrayNodeHandler) buildArrayNodeContext(ctx context.Context, nCtx interfaces.NodeExecutionContext, arrayNodeState *handler.ArrayNodeState, arrayNode v1alpha1.ExecutableArrayNode, subNodeIndex int, eventRecorder arrayEventRecorder) (
	interfaces.Node, executors.ExecutionContext, executors.DAGStructure, executors.NodeLookup, *v1alpha1.NodeSpec, *v1alpha1.NodeStatus, error) {

	nodePhase := v1alpha1.NodePhase(arrayNodeState.SubNodePhases.GetItem(subNodeIndex)) // #nosec G115
	taskPhase := int(arrayNodeState.SubNodeTaskPhases.GetItem(subNodeIndex))            // #nosec G115

	// need to initialize the inputReader every time to ensure TaskHandler can access for cache lookups / population
	inputs, err := nCtx.InputReader().Get(ctx)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	inputLiteralMap, err := constructLiteralMap(inputs, subNodeIndex, arrayNode)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	inputReader := newStaticInputReader(nCtx.InputReader(), inputLiteralMap)

	// wrap node lookup
	subNodeSpec := *arrayNode.GetSubNodeSpec()

	subNodeID := fmt.Sprintf("n%d", subNodeIndex)
	subNodeSpec.ID = subNodeID
	subNodeSpec.Name = subNodeID
	// mock the input bindings for the subNode to nil to bypass input resolution in the
	// `nodeExecutor.preExecute` function. this is required because this function is the entrypoint
	// for initial cache lookups. an alternative solution would be to mock the datastore to bypass
	// writing the inputFile.
	subNodeSpec.InputBindings = nil

	// TODO - if we want to support more plugin types we need to figure out the best way to store plugin state
	// currently just mocking based on node phase -> which works for all k8s plugins
	// we can not pre-allocated a bit array because max size is 256B and with 5k fanout node state = 1.28MB
	pluginStateBytes := a.pluginStateBytesStarted
	if taskPhase == int(core.PhaseUndefined) || taskPhase == int(core.PhaseRetryableFailure) || taskPhase == int(core.PhaseWaitingForResources) {
		pluginStateBytes = a.pluginStateBytesNotStarted
	}

	// construct output references
	currentAttempt := uint32(arrayNodeState.SubNodeRetryAttempts.GetItem(subNodeIndex)) // #nosec G115
	subDataDir, subOutputDir, err := constructOutputReferences(ctx, nCtx, strconv.Itoa(subNodeIndex), strconv.Itoa(int(currentAttempt)))
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	// compute start time for subNode using delta timestamp from ArrayNode NodeStatus
	var startedAt *metav1.Time
	if nCtx.NodeStatus().GetLastAttemptStartedAt() != nil && arrayNodeState.SubNodeDeltaTimestamps.BitSet != nil {
		if deltaSeconds := arrayNodeState.SubNodeDeltaTimestamps.GetItem(subNodeIndex); deltaSeconds != 0 {
			startedAt = &metav1.Time{Time: nCtx.NodeStatus().GetLastAttemptStartedAt().Add(time.Duration(deltaSeconds) * time.Second)} // #nosec G115
		}
	}

	subNodeStatus := &v1alpha1.NodeStatus{
		Phase:          nodePhase,
		DataDir:        subDataDir,
		OutputDir:      subOutputDir,
		Attempts:       currentAttempt,
		SystemFailures: uint32(arrayNodeState.SubNodeSystemFailures.GetItem(subNodeIndex)), // #nosec G115
		TaskNodeStatus: &v1alpha1.TaskNodeStatus{
			Phase:       taskPhase,
			PluginState: pluginStateBytes,
		},
		LastAttemptStartedAt: startedAt,
	}

	// initialize mocks
	arrayNodeLookup := newArrayNodeLookup(nCtx.ContextualNodeLookup(), subNodeID, &subNodeSpec, subNodeStatus)

	newParentInfo, err := common.CreateParentInfo(nCtx.ExecutionContext().GetParentInfo(), nCtx.NodeID(), nCtx.CurrentAttempt(), false)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	arrayExecutionContext := newArrayExecutionContext(executors.NewExecutionContextWithParentInfo(nCtx.ExecutionContext(), newParentInfo), subNodeIndex)

	arrayNodeExecutionContextBuilder := newArrayNodeExecutionContextBuilder(a.nodeExecutor.GetNodeExecutionContextBuilder(),
		subNodeID, subNodeIndex, subNodeStatus, inputReader, eventRecorder)
	arrayNodeExecutor := a.nodeExecutor.WithNodeExecutionContextBuilder(arrayNodeExecutionContextBuilder)

	return arrayNodeExecutor, arrayExecutionContext, &arrayNodeLookup, &arrayNodeLookup, &subNodeSpec, subNodeStatus, nil
}
