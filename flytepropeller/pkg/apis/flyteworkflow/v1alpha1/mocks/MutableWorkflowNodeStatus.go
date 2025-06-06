// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// MutableWorkflowNodeStatus is an autogenerated mock type for the MutableWorkflowNodeStatus type
type MutableWorkflowNodeStatus struct {
	mock.Mock
}

type MutableWorkflowNodeStatus_Expecter struct {
	mock *mock.Mock
}

func (_m *MutableWorkflowNodeStatus) EXPECT() *MutableWorkflowNodeStatus_Expecter {
	return &MutableWorkflowNodeStatus_Expecter{mock: &_m.Mock}
}

// GetExecutionError provides a mock function with given fields:
func (_m *MutableWorkflowNodeStatus) GetExecutionError() *core.ExecutionError {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetExecutionError")
	}

	var r0 *core.ExecutionError
	if rf, ok := ret.Get(0).(func() *core.ExecutionError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ExecutionError)
		}
	}

	return r0
}

// MutableWorkflowNodeStatus_GetExecutionError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExecutionError'
type MutableWorkflowNodeStatus_GetExecutionError_Call struct {
	*mock.Call
}

// GetExecutionError is a helper method to define mock.On call
func (_e *MutableWorkflowNodeStatus_Expecter) GetExecutionError() *MutableWorkflowNodeStatus_GetExecutionError_Call {
	return &MutableWorkflowNodeStatus_GetExecutionError_Call{Call: _e.mock.On("GetExecutionError")}
}

func (_c *MutableWorkflowNodeStatus_GetExecutionError_Call) Run(run func()) *MutableWorkflowNodeStatus_GetExecutionError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableWorkflowNodeStatus_GetExecutionError_Call) Return(_a0 *core.ExecutionError) *MutableWorkflowNodeStatus_GetExecutionError_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableWorkflowNodeStatus_GetExecutionError_Call) RunAndReturn(run func() *core.ExecutionError) *MutableWorkflowNodeStatus_GetExecutionError_Call {
	_c.Call.Return(run)
	return _c
}

// GetWorkflowNodePhase provides a mock function with given fields:
func (_m *MutableWorkflowNodeStatus) GetWorkflowNodePhase() v1alpha1.WorkflowNodePhase {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetWorkflowNodePhase")
	}

	var r0 v1alpha1.WorkflowNodePhase
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowNodePhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowNodePhase)
	}

	return r0
}

// MutableWorkflowNodeStatus_GetWorkflowNodePhase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWorkflowNodePhase'
type MutableWorkflowNodeStatus_GetWorkflowNodePhase_Call struct {
	*mock.Call
}

// GetWorkflowNodePhase is a helper method to define mock.On call
func (_e *MutableWorkflowNodeStatus_Expecter) GetWorkflowNodePhase() *MutableWorkflowNodeStatus_GetWorkflowNodePhase_Call {
	return &MutableWorkflowNodeStatus_GetWorkflowNodePhase_Call{Call: _e.mock.On("GetWorkflowNodePhase")}
}

func (_c *MutableWorkflowNodeStatus_GetWorkflowNodePhase_Call) Run(run func()) *MutableWorkflowNodeStatus_GetWorkflowNodePhase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableWorkflowNodeStatus_GetWorkflowNodePhase_Call) Return(_a0 v1alpha1.WorkflowNodePhase) *MutableWorkflowNodeStatus_GetWorkflowNodePhase_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableWorkflowNodeStatus_GetWorkflowNodePhase_Call) RunAndReturn(run func() v1alpha1.WorkflowNodePhase) *MutableWorkflowNodeStatus_GetWorkflowNodePhase_Call {
	_c.Call.Return(run)
	return _c
}

// IsDirty provides a mock function with given fields:
func (_m *MutableWorkflowNodeStatus) IsDirty() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsDirty")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MutableWorkflowNodeStatus_IsDirty_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsDirty'
type MutableWorkflowNodeStatus_IsDirty_Call struct {
	*mock.Call
}

// IsDirty is a helper method to define mock.On call
func (_e *MutableWorkflowNodeStatus_Expecter) IsDirty() *MutableWorkflowNodeStatus_IsDirty_Call {
	return &MutableWorkflowNodeStatus_IsDirty_Call{Call: _e.mock.On("IsDirty")}
}

func (_c *MutableWorkflowNodeStatus_IsDirty_Call) Run(run func()) *MutableWorkflowNodeStatus_IsDirty_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableWorkflowNodeStatus_IsDirty_Call) Return(_a0 bool) *MutableWorkflowNodeStatus_IsDirty_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableWorkflowNodeStatus_IsDirty_Call) RunAndReturn(run func() bool) *MutableWorkflowNodeStatus_IsDirty_Call {
	_c.Call.Return(run)
	return _c
}

// SetExecutionError provides a mock function with given fields: executionError
func (_m *MutableWorkflowNodeStatus) SetExecutionError(executionError *core.ExecutionError) {
	_m.Called(executionError)
}

// MutableWorkflowNodeStatus_SetExecutionError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetExecutionError'
type MutableWorkflowNodeStatus_SetExecutionError_Call struct {
	*mock.Call
}

// SetExecutionError is a helper method to define mock.On call
//   - executionError *core.ExecutionError
func (_e *MutableWorkflowNodeStatus_Expecter) SetExecutionError(executionError interface{}) *MutableWorkflowNodeStatus_SetExecutionError_Call {
	return &MutableWorkflowNodeStatus_SetExecutionError_Call{Call: _e.mock.On("SetExecutionError", executionError)}
}

func (_c *MutableWorkflowNodeStatus_SetExecutionError_Call) Run(run func(executionError *core.ExecutionError)) *MutableWorkflowNodeStatus_SetExecutionError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*core.ExecutionError))
	})
	return _c
}

func (_c *MutableWorkflowNodeStatus_SetExecutionError_Call) Return() *MutableWorkflowNodeStatus_SetExecutionError_Call {
	_c.Call.Return()
	return _c
}

func (_c *MutableWorkflowNodeStatus_SetExecutionError_Call) RunAndReturn(run func(*core.ExecutionError)) *MutableWorkflowNodeStatus_SetExecutionError_Call {
	_c.Call.Return(run)
	return _c
}

// SetWorkflowNodePhase provides a mock function with given fields: phase
func (_m *MutableWorkflowNodeStatus) SetWorkflowNodePhase(phase v1alpha1.WorkflowNodePhase) {
	_m.Called(phase)
}

// MutableWorkflowNodeStatus_SetWorkflowNodePhase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetWorkflowNodePhase'
type MutableWorkflowNodeStatus_SetWorkflowNodePhase_Call struct {
	*mock.Call
}

// SetWorkflowNodePhase is a helper method to define mock.On call
//   - phase v1alpha1.WorkflowNodePhase
func (_e *MutableWorkflowNodeStatus_Expecter) SetWorkflowNodePhase(phase interface{}) *MutableWorkflowNodeStatus_SetWorkflowNodePhase_Call {
	return &MutableWorkflowNodeStatus_SetWorkflowNodePhase_Call{Call: _e.mock.On("SetWorkflowNodePhase", phase)}
}

func (_c *MutableWorkflowNodeStatus_SetWorkflowNodePhase_Call) Run(run func(phase v1alpha1.WorkflowNodePhase)) *MutableWorkflowNodeStatus_SetWorkflowNodePhase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(v1alpha1.WorkflowNodePhase))
	})
	return _c
}

func (_c *MutableWorkflowNodeStatus_SetWorkflowNodePhase_Call) Return() *MutableWorkflowNodeStatus_SetWorkflowNodePhase_Call {
	_c.Call.Return()
	return _c
}

func (_c *MutableWorkflowNodeStatus_SetWorkflowNodePhase_Call) RunAndReturn(run func(v1alpha1.WorkflowNodePhase)) *MutableWorkflowNodeStatus_SetWorkflowNodePhase_Call {
	_c.Call.Return(run)
	return _c
}

// NewMutableWorkflowNodeStatus creates a new instance of MutableWorkflowNodeStatus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMutableWorkflowNodeStatus(t interface {
	mock.TestingT
	Cleanup(func())
}) *MutableWorkflowNodeStatus {
	mock := &MutableWorkflowNodeStatus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
