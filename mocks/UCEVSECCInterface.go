// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	cemdapi "github.com/enbility/cemd/api"
	api "github.com/enbility/spine-go/api"

	mock "github.com/stretchr/testify/mock"

	model "github.com/enbility/spine-go/model"
)

// UCEVSECCInterface is an autogenerated mock type for the UCEVSECCInterface type
type UCEVSECCInterface struct {
	mock.Mock
}

type UCEVSECCInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *UCEVSECCInterface) EXPECT() *UCEVSECCInterface_Expecter {
	return &UCEVSECCInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *UCEVSECCInterface) AddFeatures() {
	_m.Called()
}

// UCEVSECCInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type UCEVSECCInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *UCEVSECCInterface_Expecter) AddFeatures() *UCEVSECCInterface_AddFeatures_Call {
	return &UCEVSECCInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *UCEVSECCInterface_AddFeatures_Call) Run(run func()) *UCEVSECCInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UCEVSECCInterface_AddFeatures_Call) Return() *UCEVSECCInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *UCEVSECCInterface_AddFeatures_Call) RunAndReturn(run func()) *UCEVSECCInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *UCEVSECCInterface) AddUseCase() {
	_m.Called()
}

// UCEVSECCInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type UCEVSECCInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *UCEVSECCInterface_Expecter) AddUseCase() *UCEVSECCInterface_AddUseCase_Call {
	return &UCEVSECCInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *UCEVSECCInterface_AddUseCase_Call) Run(run func()) *UCEVSECCInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UCEVSECCInterface_AddUseCase_Call) Return() *UCEVSECCInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *UCEVSECCInterface_AddUseCase_Call) RunAndReturn(run func()) *UCEVSECCInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// IsUseCaseSupported provides a mock function with given fields: remoteEntity
func (_m *UCEVSECCInterface) IsUseCaseSupported(remoteEntity api.EntityRemoteInterface) (bool, error) {
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

// UCEVSECCInterface_IsUseCaseSupported_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsUseCaseSupported'
type UCEVSECCInterface_IsUseCaseSupported_Call struct {
	*mock.Call
}

// IsUseCaseSupported is a helper method to define mock.On call
//   - remoteEntity api.EntityRemoteInterface
func (_e *UCEVSECCInterface_Expecter) IsUseCaseSupported(remoteEntity interface{}) *UCEVSECCInterface_IsUseCaseSupported_Call {
	return &UCEVSECCInterface_IsUseCaseSupported_Call{Call: _e.mock.On("IsUseCaseSupported", remoteEntity)}
}

func (_c *UCEVSECCInterface_IsUseCaseSupported_Call) Run(run func(remoteEntity api.EntityRemoteInterface)) *UCEVSECCInterface_IsUseCaseSupported_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCEVSECCInterface_IsUseCaseSupported_Call) Return(_a0 bool, _a1 error) *UCEVSECCInterface_IsUseCaseSupported_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCEVSECCInterface_IsUseCaseSupported_Call) RunAndReturn(run func(api.EntityRemoteInterface) (bool, error)) *UCEVSECCInterface_IsUseCaseSupported_Call {
	_c.Call.Return(run)
	return _c
}

// ManufacturerData provides a mock function with given fields: entity
func (_m *UCEVSECCInterface) ManufacturerData(entity api.EntityRemoteInterface) (cemdapi.ManufacturerData, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for ManufacturerData")
	}

	var r0 cemdapi.ManufacturerData
	var r1 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (cemdapi.ManufacturerData, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) cemdapi.ManufacturerData); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(cemdapi.ManufacturerData)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UCEVSECCInterface_ManufacturerData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ManufacturerData'
type UCEVSECCInterface_ManufacturerData_Call struct {
	*mock.Call
}

// ManufacturerData is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCEVSECCInterface_Expecter) ManufacturerData(entity interface{}) *UCEVSECCInterface_ManufacturerData_Call {
	return &UCEVSECCInterface_ManufacturerData_Call{Call: _e.mock.On("ManufacturerData", entity)}
}

func (_c *UCEVSECCInterface_ManufacturerData_Call) Run(run func(entity api.EntityRemoteInterface)) *UCEVSECCInterface_ManufacturerData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCEVSECCInterface_ManufacturerData_Call) Return(_a0 cemdapi.ManufacturerData, _a1 error) *UCEVSECCInterface_ManufacturerData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UCEVSECCInterface_ManufacturerData_Call) RunAndReturn(run func(api.EntityRemoteInterface) (cemdapi.ManufacturerData, error)) *UCEVSECCInterface_ManufacturerData_Call {
	_c.Call.Return(run)
	return _c
}

// OperatingState provides a mock function with given fields: entity
func (_m *UCEVSECCInterface) OperatingState(entity api.EntityRemoteInterface) (model.DeviceDiagnosisOperatingStateType, string, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for OperatingState")
	}

	var r0 model.DeviceDiagnosisOperatingStateType
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) (model.DeviceDiagnosisOperatingStateType, string, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface) model.DeviceDiagnosisOperatingStateType); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(model.DeviceDiagnosisOperatingStateType)
	}

	if rf, ok := ret.Get(1).(func(api.EntityRemoteInterface) string); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(api.EntityRemoteInterface) error); ok {
		r2 = rf(entity)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UCEVSECCInterface_OperatingState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OperatingState'
type UCEVSECCInterface_OperatingState_Call struct {
	*mock.Call
}

// OperatingState is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *UCEVSECCInterface_Expecter) OperatingState(entity interface{}) *UCEVSECCInterface_OperatingState_Call {
	return &UCEVSECCInterface_OperatingState_Call{Call: _e.mock.On("OperatingState", entity)}
}

func (_c *UCEVSECCInterface_OperatingState_Call) Run(run func(entity api.EntityRemoteInterface)) *UCEVSECCInterface_OperatingState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *UCEVSECCInterface_OperatingState_Call) Return(_a0 model.DeviceDiagnosisOperatingStateType, _a1 string, _a2 error) *UCEVSECCInterface_OperatingState_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *UCEVSECCInterface_OperatingState_Call) RunAndReturn(run func(api.EntityRemoteInterface) (model.DeviceDiagnosisOperatingStateType, string, error)) *UCEVSECCInterface_OperatingState_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *UCEVSECCInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// UCEVSECCInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type UCEVSECCInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *UCEVSECCInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *UCEVSECCInterface_UpdateUseCaseAvailability_Call {
	return &UCEVSECCInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *UCEVSECCInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *UCEVSECCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *UCEVSECCInterface_UpdateUseCaseAvailability_Call) Return() *UCEVSECCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *UCEVSECCInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *UCEVSECCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// UseCaseName provides a mock function with given fields:
func (_m *UCEVSECCInterface) UseCaseName() model.UseCaseNameType {
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

// UCEVSECCInterface_UseCaseName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UseCaseName'
type UCEVSECCInterface_UseCaseName_Call struct {
	*mock.Call
}

// UseCaseName is a helper method to define mock.On call
func (_e *UCEVSECCInterface_Expecter) UseCaseName() *UCEVSECCInterface_UseCaseName_Call {
	return &UCEVSECCInterface_UseCaseName_Call{Call: _e.mock.On("UseCaseName")}
}

func (_c *UCEVSECCInterface_UseCaseName_Call) Run(run func()) *UCEVSECCInterface_UseCaseName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UCEVSECCInterface_UseCaseName_Call) Return(_a0 model.UseCaseNameType) *UCEVSECCInterface_UseCaseName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UCEVSECCInterface_UseCaseName_Call) RunAndReturn(run func() model.UseCaseNameType) *UCEVSECCInterface_UseCaseName_Call {
	_c.Call.Return(run)
	return _c
}

// NewUCEVSECCInterface creates a new instance of UCEVSECCInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUCEVSECCInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UCEVSECCInterface {
	mock := &UCEVSECCInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
