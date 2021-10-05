// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package repository

import (
	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/network/armnetwork"
	mock "github.com/stretchr/testify/mock"
)

// mockSubnetsClient is an autogenerated mock type for the subnetsClient type
type mockSubnetsClient struct {
	mock.Mock
}

// List provides a mock function with given fields: resourceGroupName, virtualNetworkName, options
func (_m *mockSubnetsClient) List(resourceGroupName string, virtualNetworkName string, options *armnetwork.SubnetsListOptions) subnetsListPager {
	ret := _m.Called(resourceGroupName, virtualNetworkName, options)

	var r0 subnetsListPager
	if rf, ok := ret.Get(0).(func(string, string, *armnetwork.SubnetsListOptions) subnetsListPager); ok {
		r0 = rf(resourceGroupName, virtualNetworkName, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(subnetsListPager)
		}
	}

	return r0
}
