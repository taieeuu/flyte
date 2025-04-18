// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	time "time"

	storage "github.com/flyteorg/flyte/flytestdlib/storage"
	mock "github.com/stretchr/testify/mock"
)

// MutableTaskNodeStatus is an autogenerated mock type for the MutableTaskNodeStatus type
type MutableTaskNodeStatus struct {
	mock.Mock
}

type MutableTaskNodeStatus_Expecter struct {
	mock *mock.Mock
}

func (_m *MutableTaskNodeStatus) EXPECT() *MutableTaskNodeStatus_Expecter {
	return &MutableTaskNodeStatus_Expecter{mock: &_m.Mock}
}

// GetBarrierClockTick provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetBarrierClockTick() uint32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBarrierClockTick")
	}

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// MutableTaskNodeStatus_GetBarrierClockTick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBarrierClockTick'
type MutableTaskNodeStatus_GetBarrierClockTick_Call struct {
	*mock.Call
}

// GetBarrierClockTick is a helper method to define mock.On call
func (_e *MutableTaskNodeStatus_Expecter) GetBarrierClockTick() *MutableTaskNodeStatus_GetBarrierClockTick_Call {
	return &MutableTaskNodeStatus_GetBarrierClockTick_Call{Call: _e.mock.On("GetBarrierClockTick")}
}

func (_c *MutableTaskNodeStatus_GetBarrierClockTick_Call) Run(run func()) *MutableTaskNodeStatus_GetBarrierClockTick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableTaskNodeStatus_GetBarrierClockTick_Call) Return(_a0 uint32) *MutableTaskNodeStatus_GetBarrierClockTick_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableTaskNodeStatus_GetBarrierClockTick_Call) RunAndReturn(run func() uint32) *MutableTaskNodeStatus_GetBarrierClockTick_Call {
	_c.Call.Return(run)
	return _c
}

// GetCleanupOnFailure provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetCleanupOnFailure() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCleanupOnFailure")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MutableTaskNodeStatus_GetCleanupOnFailure_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCleanupOnFailure'
type MutableTaskNodeStatus_GetCleanupOnFailure_Call struct {
	*mock.Call
}

// GetCleanupOnFailure is a helper method to define mock.On call
func (_e *MutableTaskNodeStatus_Expecter) GetCleanupOnFailure() *MutableTaskNodeStatus_GetCleanupOnFailure_Call {
	return &MutableTaskNodeStatus_GetCleanupOnFailure_Call{Call: _e.mock.On("GetCleanupOnFailure")}
}

func (_c *MutableTaskNodeStatus_GetCleanupOnFailure_Call) Run(run func()) *MutableTaskNodeStatus_GetCleanupOnFailure_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableTaskNodeStatus_GetCleanupOnFailure_Call) Return(_a0 bool) *MutableTaskNodeStatus_GetCleanupOnFailure_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableTaskNodeStatus_GetCleanupOnFailure_Call) RunAndReturn(run func() bool) *MutableTaskNodeStatus_GetCleanupOnFailure_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastPhaseUpdatedAt provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetLastPhaseUpdatedAt() time.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLastPhaseUpdatedAt")
	}

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// MutableTaskNodeStatus_GetLastPhaseUpdatedAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastPhaseUpdatedAt'
type MutableTaskNodeStatus_GetLastPhaseUpdatedAt_Call struct {
	*mock.Call
}

// GetLastPhaseUpdatedAt is a helper method to define mock.On call
func (_e *MutableTaskNodeStatus_Expecter) GetLastPhaseUpdatedAt() *MutableTaskNodeStatus_GetLastPhaseUpdatedAt_Call {
	return &MutableTaskNodeStatus_GetLastPhaseUpdatedAt_Call{Call: _e.mock.On("GetLastPhaseUpdatedAt")}
}

func (_c *MutableTaskNodeStatus_GetLastPhaseUpdatedAt_Call) Run(run func()) *MutableTaskNodeStatus_GetLastPhaseUpdatedAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableTaskNodeStatus_GetLastPhaseUpdatedAt_Call) Return(_a0 time.Time) *MutableTaskNodeStatus_GetLastPhaseUpdatedAt_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableTaskNodeStatus_GetLastPhaseUpdatedAt_Call) RunAndReturn(run func() time.Time) *MutableTaskNodeStatus_GetLastPhaseUpdatedAt_Call {
	_c.Call.Return(run)
	return _c
}

// GetPhase provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetPhase() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPhase")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MutableTaskNodeStatus_GetPhase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPhase'
type MutableTaskNodeStatus_GetPhase_Call struct {
	*mock.Call
}

// GetPhase is a helper method to define mock.On call
func (_e *MutableTaskNodeStatus_Expecter) GetPhase() *MutableTaskNodeStatus_GetPhase_Call {
	return &MutableTaskNodeStatus_GetPhase_Call{Call: _e.mock.On("GetPhase")}
}

func (_c *MutableTaskNodeStatus_GetPhase_Call) Run(run func()) *MutableTaskNodeStatus_GetPhase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableTaskNodeStatus_GetPhase_Call) Return(_a0 int) *MutableTaskNodeStatus_GetPhase_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableTaskNodeStatus_GetPhase_Call) RunAndReturn(run func() int) *MutableTaskNodeStatus_GetPhase_Call {
	_c.Call.Return(run)
	return _c
}

// GetPhaseVersion provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetPhaseVersion() uint32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPhaseVersion")
	}

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// MutableTaskNodeStatus_GetPhaseVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPhaseVersion'
type MutableTaskNodeStatus_GetPhaseVersion_Call struct {
	*mock.Call
}

// GetPhaseVersion is a helper method to define mock.On call
func (_e *MutableTaskNodeStatus_Expecter) GetPhaseVersion() *MutableTaskNodeStatus_GetPhaseVersion_Call {
	return &MutableTaskNodeStatus_GetPhaseVersion_Call{Call: _e.mock.On("GetPhaseVersion")}
}

func (_c *MutableTaskNodeStatus_GetPhaseVersion_Call) Run(run func()) *MutableTaskNodeStatus_GetPhaseVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableTaskNodeStatus_GetPhaseVersion_Call) Return(_a0 uint32) *MutableTaskNodeStatus_GetPhaseVersion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableTaskNodeStatus_GetPhaseVersion_Call) RunAndReturn(run func() uint32) *MutableTaskNodeStatus_GetPhaseVersion_Call {
	_c.Call.Return(run)
	return _c
}

// GetPluginState provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetPluginState() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPluginState")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MutableTaskNodeStatus_GetPluginState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPluginState'
type MutableTaskNodeStatus_GetPluginState_Call struct {
	*mock.Call
}

// GetPluginState is a helper method to define mock.On call
func (_e *MutableTaskNodeStatus_Expecter) GetPluginState() *MutableTaskNodeStatus_GetPluginState_Call {
	return &MutableTaskNodeStatus_GetPluginState_Call{Call: _e.mock.On("GetPluginState")}
}

func (_c *MutableTaskNodeStatus_GetPluginState_Call) Run(run func()) *MutableTaskNodeStatus_GetPluginState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableTaskNodeStatus_GetPluginState_Call) Return(_a0 []byte) *MutableTaskNodeStatus_GetPluginState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableTaskNodeStatus_GetPluginState_Call) RunAndReturn(run func() []byte) *MutableTaskNodeStatus_GetPluginState_Call {
	_c.Call.Return(run)
	return _c
}

// GetPluginStateVersion provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetPluginStateVersion() uint32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPluginStateVersion")
	}

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// MutableTaskNodeStatus_GetPluginStateVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPluginStateVersion'
type MutableTaskNodeStatus_GetPluginStateVersion_Call struct {
	*mock.Call
}

// GetPluginStateVersion is a helper method to define mock.On call
func (_e *MutableTaskNodeStatus_Expecter) GetPluginStateVersion() *MutableTaskNodeStatus_GetPluginStateVersion_Call {
	return &MutableTaskNodeStatus_GetPluginStateVersion_Call{Call: _e.mock.On("GetPluginStateVersion")}
}

func (_c *MutableTaskNodeStatus_GetPluginStateVersion_Call) Run(run func()) *MutableTaskNodeStatus_GetPluginStateVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableTaskNodeStatus_GetPluginStateVersion_Call) Return(_a0 uint32) *MutableTaskNodeStatus_GetPluginStateVersion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableTaskNodeStatus_GetPluginStateVersion_Call) RunAndReturn(run func() uint32) *MutableTaskNodeStatus_GetPluginStateVersion_Call {
	_c.Call.Return(run)
	return _c
}

// GetPreviousNodeExecutionCheckpointPath provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetPreviousNodeExecutionCheckpointPath() storage.DataReference {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPreviousNodeExecutionCheckpointPath")
	}

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

// MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPreviousNodeExecutionCheckpointPath'
type MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath_Call struct {
	*mock.Call
}

// GetPreviousNodeExecutionCheckpointPath is a helper method to define mock.On call
func (_e *MutableTaskNodeStatus_Expecter) GetPreviousNodeExecutionCheckpointPath() *MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath_Call {
	return &MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath_Call{Call: _e.mock.On("GetPreviousNodeExecutionCheckpointPath")}
}

func (_c *MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath_Call) Run(run func()) *MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath_Call) Return(_a0 storage.DataReference) *MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath_Call) RunAndReturn(run func() storage.DataReference) *MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath_Call {
	_c.Call.Return(run)
	return _c
}

// IsDirty provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) IsDirty() bool {
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

// MutableTaskNodeStatus_IsDirty_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsDirty'
type MutableTaskNodeStatus_IsDirty_Call struct {
	*mock.Call
}

// IsDirty is a helper method to define mock.On call
func (_e *MutableTaskNodeStatus_Expecter) IsDirty() *MutableTaskNodeStatus_IsDirty_Call {
	return &MutableTaskNodeStatus_IsDirty_Call{Call: _e.mock.On("IsDirty")}
}

func (_c *MutableTaskNodeStatus_IsDirty_Call) Run(run func()) *MutableTaskNodeStatus_IsDirty_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableTaskNodeStatus_IsDirty_Call) Return(_a0 bool) *MutableTaskNodeStatus_IsDirty_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableTaskNodeStatus_IsDirty_Call) RunAndReturn(run func() bool) *MutableTaskNodeStatus_IsDirty_Call {
	_c.Call.Return(run)
	return _c
}

// SetBarrierClockTick provides a mock function with given fields: tick
func (_m *MutableTaskNodeStatus) SetBarrierClockTick(tick uint32) {
	_m.Called(tick)
}

// MutableTaskNodeStatus_SetBarrierClockTick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetBarrierClockTick'
type MutableTaskNodeStatus_SetBarrierClockTick_Call struct {
	*mock.Call
}

// SetBarrierClockTick is a helper method to define mock.On call
//   - tick uint32
func (_e *MutableTaskNodeStatus_Expecter) SetBarrierClockTick(tick interface{}) *MutableTaskNodeStatus_SetBarrierClockTick_Call {
	return &MutableTaskNodeStatus_SetBarrierClockTick_Call{Call: _e.mock.On("SetBarrierClockTick", tick)}
}

func (_c *MutableTaskNodeStatus_SetBarrierClockTick_Call) Run(run func(tick uint32)) *MutableTaskNodeStatus_SetBarrierClockTick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MutableTaskNodeStatus_SetBarrierClockTick_Call) Return() *MutableTaskNodeStatus_SetBarrierClockTick_Call {
	_c.Call.Return()
	return _c
}

func (_c *MutableTaskNodeStatus_SetBarrierClockTick_Call) RunAndReturn(run func(uint32)) *MutableTaskNodeStatus_SetBarrierClockTick_Call {
	_c.Call.Return(run)
	return _c
}

// SetCleanupOnFailure provides a mock function with given fields: _a0
func (_m *MutableTaskNodeStatus) SetCleanupOnFailure(_a0 bool) {
	_m.Called(_a0)
}

// MutableTaskNodeStatus_SetCleanupOnFailure_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetCleanupOnFailure'
type MutableTaskNodeStatus_SetCleanupOnFailure_Call struct {
	*mock.Call
}

// SetCleanupOnFailure is a helper method to define mock.On call
//   - _a0 bool
func (_e *MutableTaskNodeStatus_Expecter) SetCleanupOnFailure(_a0 interface{}) *MutableTaskNodeStatus_SetCleanupOnFailure_Call {
	return &MutableTaskNodeStatus_SetCleanupOnFailure_Call{Call: _e.mock.On("SetCleanupOnFailure", _a0)}
}

func (_c *MutableTaskNodeStatus_SetCleanupOnFailure_Call) Run(run func(_a0 bool)) *MutableTaskNodeStatus_SetCleanupOnFailure_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MutableTaskNodeStatus_SetCleanupOnFailure_Call) Return() *MutableTaskNodeStatus_SetCleanupOnFailure_Call {
	_c.Call.Return()
	return _c
}

func (_c *MutableTaskNodeStatus_SetCleanupOnFailure_Call) RunAndReturn(run func(bool)) *MutableTaskNodeStatus_SetCleanupOnFailure_Call {
	_c.Call.Return(run)
	return _c
}

// SetLastPhaseUpdatedAt provides a mock function with given fields: updatedAt
func (_m *MutableTaskNodeStatus) SetLastPhaseUpdatedAt(updatedAt time.Time) {
	_m.Called(updatedAt)
}

// MutableTaskNodeStatus_SetLastPhaseUpdatedAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetLastPhaseUpdatedAt'
type MutableTaskNodeStatus_SetLastPhaseUpdatedAt_Call struct {
	*mock.Call
}

// SetLastPhaseUpdatedAt is a helper method to define mock.On call
//   - updatedAt time.Time
func (_e *MutableTaskNodeStatus_Expecter) SetLastPhaseUpdatedAt(updatedAt interface{}) *MutableTaskNodeStatus_SetLastPhaseUpdatedAt_Call {
	return &MutableTaskNodeStatus_SetLastPhaseUpdatedAt_Call{Call: _e.mock.On("SetLastPhaseUpdatedAt", updatedAt)}
}

func (_c *MutableTaskNodeStatus_SetLastPhaseUpdatedAt_Call) Run(run func(updatedAt time.Time)) *MutableTaskNodeStatus_SetLastPhaseUpdatedAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Time))
	})
	return _c
}

func (_c *MutableTaskNodeStatus_SetLastPhaseUpdatedAt_Call) Return() *MutableTaskNodeStatus_SetLastPhaseUpdatedAt_Call {
	_c.Call.Return()
	return _c
}

func (_c *MutableTaskNodeStatus_SetLastPhaseUpdatedAt_Call) RunAndReturn(run func(time.Time)) *MutableTaskNodeStatus_SetLastPhaseUpdatedAt_Call {
	_c.Call.Return(run)
	return _c
}

// SetPhase provides a mock function with given fields: phase
func (_m *MutableTaskNodeStatus) SetPhase(phase int) {
	_m.Called(phase)
}

// MutableTaskNodeStatus_SetPhase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPhase'
type MutableTaskNodeStatus_SetPhase_Call struct {
	*mock.Call
}

// SetPhase is a helper method to define mock.On call
//   - phase int
func (_e *MutableTaskNodeStatus_Expecter) SetPhase(phase interface{}) *MutableTaskNodeStatus_SetPhase_Call {
	return &MutableTaskNodeStatus_SetPhase_Call{Call: _e.mock.On("SetPhase", phase)}
}

func (_c *MutableTaskNodeStatus_SetPhase_Call) Run(run func(phase int)) *MutableTaskNodeStatus_SetPhase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MutableTaskNodeStatus_SetPhase_Call) Return() *MutableTaskNodeStatus_SetPhase_Call {
	_c.Call.Return()
	return _c
}

func (_c *MutableTaskNodeStatus_SetPhase_Call) RunAndReturn(run func(int)) *MutableTaskNodeStatus_SetPhase_Call {
	_c.Call.Return(run)
	return _c
}

// SetPhaseVersion provides a mock function with given fields: version
func (_m *MutableTaskNodeStatus) SetPhaseVersion(version uint32) {
	_m.Called(version)
}

// MutableTaskNodeStatus_SetPhaseVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPhaseVersion'
type MutableTaskNodeStatus_SetPhaseVersion_Call struct {
	*mock.Call
}

// SetPhaseVersion is a helper method to define mock.On call
//   - version uint32
func (_e *MutableTaskNodeStatus_Expecter) SetPhaseVersion(version interface{}) *MutableTaskNodeStatus_SetPhaseVersion_Call {
	return &MutableTaskNodeStatus_SetPhaseVersion_Call{Call: _e.mock.On("SetPhaseVersion", version)}
}

func (_c *MutableTaskNodeStatus_SetPhaseVersion_Call) Run(run func(version uint32)) *MutableTaskNodeStatus_SetPhaseVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MutableTaskNodeStatus_SetPhaseVersion_Call) Return() *MutableTaskNodeStatus_SetPhaseVersion_Call {
	_c.Call.Return()
	return _c
}

func (_c *MutableTaskNodeStatus_SetPhaseVersion_Call) RunAndReturn(run func(uint32)) *MutableTaskNodeStatus_SetPhaseVersion_Call {
	_c.Call.Return(run)
	return _c
}

// SetPluginState provides a mock function with given fields: _a0
func (_m *MutableTaskNodeStatus) SetPluginState(_a0 []byte) {
	_m.Called(_a0)
}

// MutableTaskNodeStatus_SetPluginState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPluginState'
type MutableTaskNodeStatus_SetPluginState_Call struct {
	*mock.Call
}

// SetPluginState is a helper method to define mock.On call
//   - _a0 []byte
func (_e *MutableTaskNodeStatus_Expecter) SetPluginState(_a0 interface{}) *MutableTaskNodeStatus_SetPluginState_Call {
	return &MutableTaskNodeStatus_SetPluginState_Call{Call: _e.mock.On("SetPluginState", _a0)}
}

func (_c *MutableTaskNodeStatus_SetPluginState_Call) Run(run func(_a0 []byte)) *MutableTaskNodeStatus_SetPluginState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MutableTaskNodeStatus_SetPluginState_Call) Return() *MutableTaskNodeStatus_SetPluginState_Call {
	_c.Call.Return()
	return _c
}

func (_c *MutableTaskNodeStatus_SetPluginState_Call) RunAndReturn(run func([]byte)) *MutableTaskNodeStatus_SetPluginState_Call {
	_c.Call.Return(run)
	return _c
}

// SetPluginStateVersion provides a mock function with given fields: _a0
func (_m *MutableTaskNodeStatus) SetPluginStateVersion(_a0 uint32) {
	_m.Called(_a0)
}

// MutableTaskNodeStatus_SetPluginStateVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPluginStateVersion'
type MutableTaskNodeStatus_SetPluginStateVersion_Call struct {
	*mock.Call
}

// SetPluginStateVersion is a helper method to define mock.On call
//   - _a0 uint32
func (_e *MutableTaskNodeStatus_Expecter) SetPluginStateVersion(_a0 interface{}) *MutableTaskNodeStatus_SetPluginStateVersion_Call {
	return &MutableTaskNodeStatus_SetPluginStateVersion_Call{Call: _e.mock.On("SetPluginStateVersion", _a0)}
}

func (_c *MutableTaskNodeStatus_SetPluginStateVersion_Call) Run(run func(_a0 uint32)) *MutableTaskNodeStatus_SetPluginStateVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MutableTaskNodeStatus_SetPluginStateVersion_Call) Return() *MutableTaskNodeStatus_SetPluginStateVersion_Call {
	_c.Call.Return()
	return _c
}

func (_c *MutableTaskNodeStatus_SetPluginStateVersion_Call) RunAndReturn(run func(uint32)) *MutableTaskNodeStatus_SetPluginStateVersion_Call {
	_c.Call.Return(run)
	return _c
}

// SetPreviousNodeExecutionCheckpointPath provides a mock function with given fields: _a0
func (_m *MutableTaskNodeStatus) SetPreviousNodeExecutionCheckpointPath(_a0 storage.DataReference) {
	_m.Called(_a0)
}

// MutableTaskNodeStatus_SetPreviousNodeExecutionCheckpointPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPreviousNodeExecutionCheckpointPath'
type MutableTaskNodeStatus_SetPreviousNodeExecutionCheckpointPath_Call struct {
	*mock.Call
}

// SetPreviousNodeExecutionCheckpointPath is a helper method to define mock.On call
//   - _a0 storage.DataReference
func (_e *MutableTaskNodeStatus_Expecter) SetPreviousNodeExecutionCheckpointPath(_a0 interface{}) *MutableTaskNodeStatus_SetPreviousNodeExecutionCheckpointPath_Call {
	return &MutableTaskNodeStatus_SetPreviousNodeExecutionCheckpointPath_Call{Call: _e.mock.On("SetPreviousNodeExecutionCheckpointPath", _a0)}
}

func (_c *MutableTaskNodeStatus_SetPreviousNodeExecutionCheckpointPath_Call) Run(run func(_a0 storage.DataReference)) *MutableTaskNodeStatus_SetPreviousNodeExecutionCheckpointPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(storage.DataReference))
	})
	return _c
}

func (_c *MutableTaskNodeStatus_SetPreviousNodeExecutionCheckpointPath_Call) Return() *MutableTaskNodeStatus_SetPreviousNodeExecutionCheckpointPath_Call {
	_c.Call.Return()
	return _c
}

func (_c *MutableTaskNodeStatus_SetPreviousNodeExecutionCheckpointPath_Call) RunAndReturn(run func(storage.DataReference)) *MutableTaskNodeStatus_SetPreviousNodeExecutionCheckpointPath_Call {
	_c.Call.Return(run)
	return _c
}

// NewMutableTaskNodeStatus creates a new instance of MutableTaskNodeStatus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMutableTaskNodeStatus(t interface {
	mock.TestingT
	Cleanup(func())
}) *MutableTaskNodeStatus {
	mock := &MutableTaskNodeStatus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
