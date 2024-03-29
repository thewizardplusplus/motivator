// Code generated by mockery v2.14.0. DO NOT EDIT.

package entities

import mock "github.com/stretchr/testify/mock"

// MockTaskHandler is an autogenerated mock type for the TaskHandler type
type MockTaskHandler struct {
	mock.Mock
}

// HandleTask provides a mock function with given fields: task
func (_m *MockTaskHandler) HandleTask(task Task) {
	_m.Called(task)
}

type mockConstructorTestingTNewMockTaskHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTaskHandler creates a new instance of MockTaskHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTaskHandler(t mockConstructorTestingTNewMockTaskHandler) *MockTaskHandler {
	mock := &MockTaskHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
