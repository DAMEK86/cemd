// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	cemdapi "github.com/enbility/cemd/api"
	api "github.com/enbility/spine-go/api"

	mock "github.com/stretchr/testify/mock"
)

// DeviceEventCallback is an autogenerated mock type for the DeviceEventCallback type
type DeviceEventCallback struct {
	mock.Mock
}

type DeviceEventCallback_Expecter struct {
	mock *mock.Mock
}

func (_m *DeviceEventCallback) EXPECT() *DeviceEventCallback_Expecter {
	return &DeviceEventCallback_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ski, device, event
func (_m *DeviceEventCallback) Execute(ski string, device api.DeviceRemoteInterface, event cemdapi.EventType) {
	_m.Called(ski, device, event)
}

// DeviceEventCallback_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type DeviceEventCallback_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ski string
//   - device api.DeviceRemoteInterface
//   - event cemdapi.EventType
func (_e *DeviceEventCallback_Expecter) Execute(ski interface{}, device interface{}, event interface{}) *DeviceEventCallback_Execute_Call {
	return &DeviceEventCallback_Execute_Call{Call: _e.mock.On("Execute", ski, device, event)}
}

func (_c *DeviceEventCallback_Execute_Call) Run(run func(ski string, device api.DeviceRemoteInterface, event cemdapi.EventType)) *DeviceEventCallback_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(api.DeviceRemoteInterface), args[2].(cemdapi.EventType))
	})
	return _c
}

func (_c *DeviceEventCallback_Execute_Call) Return() *DeviceEventCallback_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *DeviceEventCallback_Execute_Call) RunAndReturn(run func(string, api.DeviceRemoteInterface, cemdapi.EventType)) *DeviceEventCallback_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewDeviceEventCallback creates a new instance of DeviceEventCallback. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceEventCallback(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceEventCallback {
	mock := &DeviceEventCallback{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
