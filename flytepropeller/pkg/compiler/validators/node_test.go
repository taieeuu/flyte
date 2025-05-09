package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyte/flytepropeller/pkg/compiler/common"
	"github.com/flyteorg/flyte/flytepropeller/pkg/compiler/common/mocks"
	"github.com/flyteorg/flyte/flytepropeller/pkg/compiler/errors"
)

func TestValidateBranchNode(t *testing.T) {
	t.Run("No Case No Default", func(t *testing.T) {
		n := &mocks.NodeBuilder{}
		n.EXPECT().GetBranchNode().Return(&core.BranchNode{
			IfElse: &core.IfElseBlock{},
		})
		n.EXPECT().GetId().Return("node1")

		wf := &mocks.WorkflowBuilder{}
		errs := errors.NewCompileErrors()
		_, ok := ValidateBranchNode(wf, n, false, errs)
		assert.False(t, ok)
		if !errs.HasErrors() {
			assert.Error(t, errs)
		} else {
			t.Log(errs)
			errsList := errs.Errors().List()
			assert.Len(t, errsList, 2)
			assert.Equal(t, errors.BranchNodeHasNoCondition, errsList[0].Code())
			assert.Equal(t, errors.BranchNodeHasNoDefault, errsList[1].Code())
		}
	})
}

func TestValidateNode(t *testing.T) {
	t.Run("Start-node", func(t *testing.T) {
		n := &mocks.NodeBuilder{}
		n.EXPECT().GetId().Return(common.StartNodeID)

		wf := &mocks.WorkflowBuilder{}
		errs := errors.NewCompileErrors()
		ValidateNode(wf, n, true, errs)
		if !assert.False(t, errs.HasErrors()) {
			assert.NoError(t, errs)
		}
	})

	t.Run("Sort upstream node ids", func(t *testing.T) {
		n := &mocks.NodeBuilder{}
		n.EXPECT().GetId().Return("my-node")
		n.EXPECT().GetInterface().Return(&core.TypedInterface{
			Outputs: &core.VariableMap{},
			Inputs:  &core.VariableMap{},
		})
		n.EXPECT().GetOutputAliases().Return(nil)
		n.EXPECT().GetBranchNode().Return(nil)
		n.EXPECT().GetWorkflowNode().Return(nil)
		n.EXPECT().GetTaskNode().Return(nil)

		coreN := &core.Node{}
		coreN.UpstreamNodeIds = []string{"n1", "n0"}
		n.EXPECT().GetCoreNode().Return(coreN)
		n.On("GetUpstreamNodeIds").Return(func() []string {
			return coreN.GetUpstreamNodeIds()
		})

		wf := &mocks.WorkflowBuilder{}
		errs := errors.NewCompileErrors()
		ValidateNode(wf, n, true, errs)
		if !assert.False(t, errs.HasErrors()) {
			assert.NoError(t, errs)
		}

		assert.Equal(t, []string{"n0", "n1"}, n.GetUpstreamNodeIds())
	})
}
