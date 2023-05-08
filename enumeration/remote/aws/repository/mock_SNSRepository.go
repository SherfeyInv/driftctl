// Code generated by mockery v2.28.1. DO NOT EDIT.

package repository

import (
	sns "github.com/aws/aws-sdk-go/service/sns"
	mock "github.com/stretchr/testify/mock"
)

// MockSNSRepository is an autogenerated mock type for the SNSRepository type
type MockSNSRepository struct {
	mock.Mock
}

// ListAllSubscriptions provides a mock function with given fields:
func (_m *MockSNSRepository) ListAllSubscriptions() ([]*sns.Subscription, error) {
	ret := _m.Called()

	var r0 []*sns.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*sns.Subscription, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*sns.Subscription); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*sns.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllTopics provides a mock function with given fields:
func (_m *MockSNSRepository) ListAllTopics() ([]*sns.Topic, error) {
	ret := _m.Called()

	var r0 []*sns.Topic
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*sns.Topic, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*sns.Topic); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*sns.Topic)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockSNSRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockSNSRepository creates a new instance of MockSNSRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSNSRepository(t mockConstructorTestingTNewMockSNSRepository) *MockSNSRepository {
	mock := &MockSNSRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
