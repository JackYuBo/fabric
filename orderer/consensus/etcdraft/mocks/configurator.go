// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import cluster "github.com/hyperledger/fabric/orderer/common/cluster"

import mock "github.com/stretchr/testify/mock"

// Configurator is an autogenerated mock type for the Configurator type
type Configurator struct {
	mock.Mock
}

// Configure provides a mock function with given fields: channel, newNodes
func (_m *Configurator) Configure(channel string, newNodes []cluster.RemoteNode) {
	_m.Called(channel, newNodes)
}