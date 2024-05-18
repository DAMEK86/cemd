// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/spine-go/api"
	mock "github.com/stretchr/testify/mock"

	model "github.com/enbility/spine-go/model"
)

// UCMGCPInterface is an autogenerated mock type for the UCMGCPInterface type
type UCMGCPInterface struct {
	mock.Mock
}

type UCMGCPInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *UCMGCPInterface) EXPECT() *UCMGCPInterface_Expecter {
	return &UCMGCPInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *UCMGCPInterface) AddFeatures() {
	_m.Called()
}

// UCMGCPInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type UCMGCPInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *UCMGCPInterface_Expecter) AddFeatures() *UCMGCPInterface_AddFeatures_Call {
	return &UCMGCPInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *UCMGCPInterface_AddFeatures_Call) Run(run func()) *UCMGCPInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UCMGCPInterface_AddFeatures_Call) Return() *UCMGCPInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *UCMGCPInterface_AddFeatures_Call) RunAndReturn(run func()) *UCMGCPInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *UCMGCPInterface) AddUseCase() {
	_m.Called()
}

// UCMGCPInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type UCMGCPInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *UCMGCPInterface_Expecter) AddUseCase() *UCMGCPInterface_AddUseCase_Call {
	return &UCMGCPInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *UCMGCPInterface_AddUseCase_Call) Run(run func()) *UCMGCPInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UCMGCPInterface_AddUseCase_Call) Return() *UCMGCPInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *UCMGCPInterface_AddUseCase_Call) RunAndReturn(run func()) *UCMGCPInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// CurrentPerPhase provides a mock function with given fields: entity
func (_m *UCMGCPInterface) CurrentPerPhase(entity api.EntityRemoteInterface) ([]float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for CurrentPerPhase")
	}

	var r0 []float64
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) ([]float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) []float64); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float64)
		}
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCMGCPInterface_CurrentPerPhase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CurrentPerPhase'
type UCMGCPInterface_CurrentPerPhase_Call struct {
	*mock.Call
}

// CurrentPerPhase is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCMGCPInterface_Expecter) CurrentPerPhase(entity interface{}) *UCMGCPInterface_CurrentPerPhase_Call {
	return &UCMGCPInterface_CurrentPerPhase_Call{Call: _e.mock.On("CurrentPerPhase", entity)}
}

func (_c *UCMGCPInterface_CurrentPerPhase_Call) Run(run func(entity api.EntityRemoteInterface)) *UCMGCPInterface_CurrentPerPhase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCMGCPInterface_CurrentPerPhase_Call) Return(_a0 []float64, _a1 error) *UCMGCPInterface_CurrentPerPhase_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCMGCPInterface_CurrentPerPhase_Call) RunAndReturn(run func(api.EntityRemoteInterface) ([]float64, error)) *UCMGCPInterface_CurrentPerPhase_Call {
	_c.Call.Return(run)
	return _c
}

// EnergyConsumed provides a mock function with given fields: entity
func (_m *UCMGCPInterface) EnergyConsumed(entity api.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for EnergyConsumed")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCMGCPInterface_EnergyConsumed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnergyConsumed'
type UCMGCPInterface_EnergyConsumed_Call struct {
	*mock.Call
}

// EnergyConsumed is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCMGCPInterface_Expecter) EnergyConsumed(entity interface{}) *UCMGCPInterface_EnergyConsumed_Call {
	return &UCMGCPInterface_EnergyConsumed_Call{Call: _e.mock.On("EnergyConsumed", entity)}
}

func (_c *UCMGCPInterface_EnergyConsumed_Call) Run(run func(entity api.EntityRemoteInterface)) *UCMGCPInterface_EnergyConsumed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCMGCPInterface_EnergyConsumed_Call) Return(_a0 float64, _a1 error) *UCMGCPInterface_EnergyConsumed_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCMGCPInterface_EnergyConsumed_Call) RunAndReturn(run func(api.EntityRemoteInterface) (float64, error)) *UCMGCPInterface_EnergyConsumed_Call {
	_c.Call.Return(run)
	return _c
}

// EnergyFeedIn provides a mock function with given fields: entity
func (_m *UCMGCPInterface) EnergyFeedIn(entity api.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for EnergyFeedIn")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCMGCPInterface_EnergyFeedIn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnergyFeedIn'
type UCMGCPInterface_EnergyFeedIn_Call struct {
	*mock.Call
}

// EnergyFeedIn is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCMGCPInterface_Expecter) EnergyFeedIn(entity interface{}) *UCMGCPInterface_EnergyFeedIn_Call {
	return &UCMGCPInterface_EnergyFeedIn_Call{Call: _e.mock.On("EnergyFeedIn", entity)}
}

func (_c *UCMGCPInterface_EnergyFeedIn_Call) Run(run func(entity api.EntityRemoteInterface)) *UCMGCPInterface_EnergyFeedIn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCMGCPInterface_EnergyFeedIn_Call) Return(_a0 float64, _a1 error) *UCMGCPInterface_EnergyFeedIn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCMGCPInterface_EnergyFeedIn_Call) RunAndReturn(run func(api.EntityRemoteInterface) (float64, error)) *UCMGCPInterface_EnergyFeedIn_Call {
	_c.Call.Return(run)
	return _c
}

// Frequency provides a mock function with given fields: entity
func (_m *UCMGCPInterface) Frequency(entity api.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for Frequency")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCMGCPInterface_Frequency_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Frequency'
type UCMGCPInterface_Frequency_Call struct {
	*mock.Call
}

// Frequency is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCMGCPInterface_Expecter) Frequency(entity interface{}) *UCMGCPInterface_Frequency_Call {
	return &UCMGCPInterface_Frequency_Call{Call: _e.mock.On("Frequency", entity)}
}

func (_c *UCMGCPInterface_Frequency_Call) Run(run func(entity api.EntityRemoteInterface)) *UCMGCPInterface_Frequency_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCMGCPInterface_Frequency_Call) Return(_a0 float64, _a1 error) *UCMGCPInterface_Frequency_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCMGCPInterface_Frequency_Call) RunAndReturn(run func(api.EntityRemoteInterface) (float64, error)) *UCMGCPInterface_Frequency_Call {
	_c.Call.Return(run)
	return _c
}

// IsUseCaseSupported provides a mock function with given fields: remoteEntity
func (_m *UCMGCPInterface) IsUseCaseSupported(remoteEntity api.EntityRemoteInterface) (bool, error) {
	ret := _m.Called(remoteEntity)

	if len(ret) == 0 {
		panic("no return value specified for IsUseCaseSupported")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (bool, error)); ok {
		return rf(remoteEntity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) bool); ok {
		r0 = rf(remoteEntity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(remoteEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCMGCPInterface_IsUseCaseSupported_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsUseCaseSupported'
type UCMGCPInterface_IsUseCaseSupported_Call struct {
	*mock.Call
}

// IsUseCaseSupported is a helper method to define mock.On call
//   - remoteEntity api.EntityRemoteInterface
func (_e *UCMGCPInterface_Expecter) IsUseCaseSupported(remoteEntity interface{}) *UCMGCPInterface_IsUseCaseSupported_Call {
	return &UCMGCPInterface_IsUseCaseSupported_Call{Call: _e.mock.On("IsUseCaseSupported", remoteEntity)}
}

func (_c *UCMGCPInterface_IsUseCaseSupported_Call) Run(run func(remoteEntity api.EntityRemoteInterface)) *UCMGCPInterface_IsUseCaseSupported_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCMGCPInterface_IsUseCaseSupported_Call) Return(_a0 bool, _a1 error) *UCMGCPInterface_IsUseCaseSupported_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCMGCPInterface_IsUseCaseSupported_Call) RunAndReturn(run func(api.EntityRemoteInterface) (bool, error)) *UCMGCPInterface_IsUseCaseSupported_Call {
	_c.Call.Return(run)
	return _c
}

// Power provides a mock function with given fields: entity
func (_m *UCMGCPInterface) Power(entity api.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for Power")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCMGCPInterface_Power_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Power'
type UCMGCPInterface_Power_Call struct {
	*mock.Call
}

// Power is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCMGCPInterface_Expecter) Power(entity interface{}) *UCMGCPInterface_Power_Call {
	return &UCMGCPInterface_Power_Call{Call: _e.mock.On("Power", entity)}
}

func (_c *UCMGCPInterface_Power_Call) Run(run func(entity api.EntityRemoteInterface)) *UCMGCPInterface_Power_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCMGCPInterface_Power_Call) Return(_a0 float64, _a1 error) *UCMGCPInterface_Power_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCMGCPInterface_Power_Call) RunAndReturn(run func(api.EntityRemoteInterface) (float64, error)) *UCMGCPInterface_Power_Call {
	_c.Call.Return(run)
	return _c
}

// PowerLimitationFactor provides a mock function with given fields: entity
func (_m *UCMGCPInterface) PowerLimitationFactor(entity api.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for PowerLimitationFactor")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCMGCPInterface_PowerLimitationFactor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PowerLimitationFactor'
type UCMGCPInterface_PowerLimitationFactor_Call struct {
	*mock.Call
}

// PowerLimitationFactor is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCMGCPInterface_Expecter) PowerLimitationFactor(entity interface{}) *UCMGCPInterface_PowerLimitationFactor_Call {
	return &UCMGCPInterface_PowerLimitationFactor_Call{Call: _e.mock.On("PowerLimitationFactor", entity)}
}

func (_c *UCMGCPInterface_PowerLimitationFactor_Call) Run(run func(entity api.EntityRemoteInterface)) *UCMGCPInterface_PowerLimitationFactor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCMGCPInterface_PowerLimitationFactor_Call) Return(_a0 float64, _a1 error) *UCMGCPInterface_PowerLimitationFactor_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCMGCPInterface_PowerLimitationFactor_Call) RunAndReturn(run func(api.EntityRemoteInterface) (float64, error)) *UCMGCPInterface_PowerLimitationFactor_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *UCMGCPInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// UCMGCPInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type UCMGCPInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *UCMGCPInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *UCMGCPInterface_UpdateUseCaseAvailability_Call {
	return &UCMGCPInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *UCMGCPInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *UCMGCPInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *UCMGCPInterface_UpdateUseCaseAvailability_Call) Return() *UCMGCPInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *UCMGCPInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *UCMGCPInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// UseCaseName provides a mock function with given fields:
func (_m *UCMGCPInterface) UseCaseName() model.UseCaseNameType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UseCaseName")
	}

	var r0 model.UseCaseNameType
	if rf, ok := ret.Get(0).(func() model.UseCaseNameType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.UseCaseNameType)
	}

	return r0
}

// UCMGCPInterface_UseCaseName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UseCaseName'
type UCMGCPInterface_UseCaseName_Call struct {
	*mock.Call
}

// UseCaseName is a helper method to define mock.On call
func (_e *UCMGCPInterface_Expecter) UseCaseName() *UCMGCPInterface_UseCaseName_Call {
	return &UCMGCPInterface_UseCaseName_Call{Call: _e.mock.On("UseCaseName")}
}

func (_c *UCMGCPInterface_UseCaseName_Call) Run(run func()) *UCMGCPInterface_UseCaseName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UCMGCPInterface_UseCaseName_Call) Return(_a0 model.UseCaseNameType) *UCMGCPInterface_UseCaseName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UCMGCPInterface_UseCaseName_Call) RunAndReturn(run func() model.UseCaseNameType) *UCMGCPInterface_UseCaseName_Call {
	_c.Call.Return(run)
	return _c
}

// VoltagePerPhase provides a mock function with given fields: entity
func (_m *UCMGCPInterface) VoltagePerPhase(entity api.EntityRemoteInterface) ([]float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for VoltagePerPhase")
	}

	var r0 []float64
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) ([]float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) []float64); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float64)
		}
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCMGCPInterface_VoltagePerPhase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VoltagePerPhase'
type UCMGCPInterface_VoltagePerPhase_Call struct {
	*mock.Call
}

// VoltagePerPhase is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCMGCPInterface_Expecter) VoltagePerPhase(entity interface{}) *UCMGCPInterface_VoltagePerPhase_Call {
	return &UCMGCPInterface_VoltagePerPhase_Call{Call: _e.mock.On("VoltagePerPhase", entity)}
}

func (_c *UCMGCPInterface_VoltagePerPhase_Call) Run(run func(entity api.EntityRemoteInterface)) *UCMGCPInterface_VoltagePerPhase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCMGCPInterface_VoltagePerPhase_Call) Return(_a0 []float64, _a1 error) *UCMGCPInterface_VoltagePerPhase_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCMGCPInterface_VoltagePerPhase_Call) RunAndReturn(run func(api.EntityRemoteInterface) ([]float64, error)) *UCMGCPInterface_VoltagePerPhase_Call {
	_c.Call.Return(run)
	return _c
}

// NewUCMGCPInterface creates a new instance of UCMGCPInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUCMGCPInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UCMGCPInterface {
	mock := &UCMGCPInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
