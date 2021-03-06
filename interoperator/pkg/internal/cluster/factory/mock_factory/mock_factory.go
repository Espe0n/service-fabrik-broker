// Code generated by MockGen. DO NOT EDIT.
// Source: factory.go

// Package mock_factory is a generated GoMock package.
package mock_factory

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockClusterFactory is a mock of ClusterFactory interface
type MockClusterFactory struct {
	ctrl     *gomock.Controller
	recorder *MockClusterFactoryMockRecorder
}

// MockClusterFactoryMockRecorder is the mock recorder for MockClusterFactory
type MockClusterFactoryMockRecorder struct {
	mock *MockClusterFactory
}

// NewMockClusterFactory creates a new mock instance
func NewMockClusterFactory(ctrl *gomock.Controller) *MockClusterFactory {
	mock := &MockClusterFactory{ctrl: ctrl}
	mock.recorder = &MockClusterFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterFactory) EXPECT() *MockClusterFactoryMockRecorder {
	return m.recorder
}

// GetCluster mocks base method
func (m *MockClusterFactory) GetCluster(instanceID, bindingID, serviceID, planID string) (client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster", instanceID, bindingID, serviceID, planID)
	ret0, _ := ret[0].(client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCluster indicates an expected call of GetCluster
func (mr *MockClusterFactoryMockRecorder) GetCluster(instanceID, bindingID, serviceID, planID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockClusterFactory)(nil).GetCluster), instanceID, bindingID, serviceID, planID)
}
