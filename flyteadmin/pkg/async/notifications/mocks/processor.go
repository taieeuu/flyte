// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Processor is an autogenerated mock type for the Processor type
type Processor struct {
	mock.Mock
}

type Processor_Expecter struct {
	mock *mock.Mock
}

func (_m *Processor) EXPECT() *Processor_Expecter {
	return &Processor_Expecter{mock: &_m.Mock}
}

// StartProcessing provides a mock function with given fields:
func (_m *Processor) StartProcessing() {
	_m.Called()
}

// Processor_StartProcessing_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartProcessing'
type Processor_StartProcessing_Call struct {
	*mock.Call
}

// StartProcessing is a helper method to define mock.On call
func (_e *Processor_Expecter) StartProcessing() *Processor_StartProcessing_Call {
	return &Processor_StartProcessing_Call{Call: _e.mock.On("StartProcessing")}
}

func (_c *Processor_StartProcessing_Call) Run(run func()) *Processor_StartProcessing_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Processor_StartProcessing_Call) Return() *Processor_StartProcessing_Call {
	_c.Call.Return()
	return _c
}

func (_c *Processor_StartProcessing_Call) RunAndReturn(run func()) *Processor_StartProcessing_Call {
	_c.Call.Return(run)
	return _c
}

// StopProcessing provides a mock function with given fields:
func (_m *Processor) StopProcessing() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StopProcessing")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Processor_StopProcessing_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopProcessing'
type Processor_StopProcessing_Call struct {
	*mock.Call
}

// StopProcessing is a helper method to define mock.On call
func (_e *Processor_Expecter) StopProcessing() *Processor_StopProcessing_Call {
	return &Processor_StopProcessing_Call{Call: _e.mock.On("StopProcessing")}
}

func (_c *Processor_StopProcessing_Call) Run(run func()) *Processor_StopProcessing_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Processor_StopProcessing_Call) Return(_a0 error) *Processor_StopProcessing_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Processor_StopProcessing_Call) RunAndReturn(run func() error) *Processor_StopProcessing_Call {
	_c.Call.Return(run)
	return _c
}

// NewProcessor creates a new instance of Processor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProcessor(t interface {
	mock.TestingT
	Cleanup(func())
}) *Processor {
	mock := &Processor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
