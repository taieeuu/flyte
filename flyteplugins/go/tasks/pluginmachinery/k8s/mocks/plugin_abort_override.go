// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	client "sigs.k8s.io/controller-runtime/pkg/client"

	core "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core"

	k8s "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/k8s"

	mock "github.com/stretchr/testify/mock"
)

// PluginAbortOverride is an autogenerated mock type for the PluginAbortOverride type
type PluginAbortOverride struct {
	mock.Mock
}

type PluginAbortOverride_Expecter struct {
	mock *mock.Mock
}

func (_m *PluginAbortOverride) EXPECT() *PluginAbortOverride_Expecter {
	return &PluginAbortOverride_Expecter{mock: &_m.Mock}
}

// OnAbort provides a mock function with given fields: ctx, tCtx, resource
func (_m *PluginAbortOverride) OnAbort(ctx context.Context, tCtx core.TaskExecutionContext, resource client.Object) (k8s.AbortBehavior, error) {
	ret := _m.Called(ctx, tCtx, resource)

	if len(ret) == 0 {
		panic("no return value specified for OnAbort")
	}

	var r0 k8s.AbortBehavior
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, core.TaskExecutionContext, client.Object) (k8s.AbortBehavior, error)); ok {
		return rf(ctx, tCtx, resource)
	}
	if rf, ok := ret.Get(0).(func(context.Context, core.TaskExecutionContext, client.Object) k8s.AbortBehavior); ok {
		r0 = rf(ctx, tCtx, resource)
	} else {
		r0 = ret.Get(0).(k8s.AbortBehavior)
	}

	if rf, ok := ret.Get(1).(func(context.Context, core.TaskExecutionContext, client.Object) error); ok {
		r1 = rf(ctx, tCtx, resource)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PluginAbortOverride_OnAbort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnAbort'
type PluginAbortOverride_OnAbort_Call struct {
	*mock.Call
}

// OnAbort is a helper method to define mock.On call
//   - ctx context.Context
//   - tCtx core.TaskExecutionContext
//   - resource client.Object
func (_e *PluginAbortOverride_Expecter) OnAbort(ctx interface{}, tCtx interface{}, resource interface{}) *PluginAbortOverride_OnAbort_Call {
	return &PluginAbortOverride_OnAbort_Call{Call: _e.mock.On("OnAbort", ctx, tCtx, resource)}
}

func (_c *PluginAbortOverride_OnAbort_Call) Run(run func(ctx context.Context, tCtx core.TaskExecutionContext, resource client.Object)) *PluginAbortOverride_OnAbort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(core.TaskExecutionContext), args[2].(client.Object))
	})
	return _c
}

func (_c *PluginAbortOverride_OnAbort_Call) Return(behavior k8s.AbortBehavior, err error) *PluginAbortOverride_OnAbort_Call {
	_c.Call.Return(behavior, err)
	return _c
}

func (_c *PluginAbortOverride_OnAbort_Call) RunAndReturn(run func(context.Context, core.TaskExecutionContext, client.Object) (k8s.AbortBehavior, error)) *PluginAbortOverride_OnAbort_Call {
	_c.Call.Return(run)
	return _c
}

// NewPluginAbortOverride creates a new instance of PluginAbortOverride. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPluginAbortOverride(t interface {
	mock.TestingT
	Cleanup(func())
}) *PluginAbortOverride {
	mock := &PluginAbortOverride{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
