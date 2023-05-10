// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/rancher/pkg/generated/controllers/management.cattle.io/v3 (interfaces: ClusterRegistrationTokenCache,ClusterCache)

// Package mockmgmtcontrollers is a generated GoMock package.
package mockmgmtcontrollers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v30 "github.com/rancher/rancher/pkg/generated/controllers/management.cattle.io/v3"
	labels "k8s.io/apimachinery/pkg/labels"
)

// MockClusterRegistrationTokenCache is a mock of ClusterRegistrationTokenCache interface.
type MockClusterRegistrationTokenCache struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRegistrationTokenCacheMockRecorder
}

// MockClusterRegistrationTokenCacheMockRecorder is the mock recorder for MockClusterRegistrationTokenCache.
type MockClusterRegistrationTokenCacheMockRecorder struct {
	mock *MockClusterRegistrationTokenCache
}

// NewMockClusterRegistrationTokenCache creates a new mock instance.
func NewMockClusterRegistrationTokenCache(ctrl *gomock.Controller) *MockClusterRegistrationTokenCache {
	mock := &MockClusterRegistrationTokenCache{ctrl: ctrl}
	mock.recorder = &MockClusterRegistrationTokenCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRegistrationTokenCache) EXPECT() *MockClusterRegistrationTokenCacheMockRecorder {
	return m.recorder
}

// AddIndexer mocks base method.
func (m *MockClusterRegistrationTokenCache) AddIndexer(arg0 string, arg1 v30.ClusterRegistrationTokenIndexer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddIndexer", arg0, arg1)
}

// AddIndexer indicates an expected call of AddIndexer.
func (mr *MockClusterRegistrationTokenCacheMockRecorder) AddIndexer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIndexer", reflect.TypeOf((*MockClusterRegistrationTokenCache)(nil).AddIndexer), arg0, arg1)
}

// Get mocks base method.
func (m *MockClusterRegistrationTokenCache) Get(arg0, arg1 string) (*v3.ClusterRegistrationToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v3.ClusterRegistrationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClusterRegistrationTokenCacheMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterRegistrationTokenCache)(nil).Get), arg0, arg1)
}

// GetByIndex mocks base method.
func (m *MockClusterRegistrationTokenCache) GetByIndex(arg0, arg1 string) ([]*v3.ClusterRegistrationToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIndex", arg0, arg1)
	ret0, _ := ret[0].([]*v3.ClusterRegistrationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIndex indicates an expected call of GetByIndex.
func (mr *MockClusterRegistrationTokenCacheMockRecorder) GetByIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIndex", reflect.TypeOf((*MockClusterRegistrationTokenCache)(nil).GetByIndex), arg0, arg1)
}

// List mocks base method.
func (m *MockClusterRegistrationTokenCache) List(arg0 string, arg1 labels.Selector) ([]*v3.ClusterRegistrationToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*v3.ClusterRegistrationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockClusterRegistrationTokenCacheMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterRegistrationTokenCache)(nil).List), arg0, arg1)
}

// MockClusterCache is a mock of ClusterCache interface.
type MockClusterCache struct {
	ctrl     *gomock.Controller
	recorder *MockClusterCacheMockRecorder
}

// MockClusterCacheMockRecorder is the mock recorder for MockClusterCache.
type MockClusterCacheMockRecorder struct {
	mock *MockClusterCache
}

// NewMockClusterCache creates a new mock instance.
func NewMockClusterCache(ctrl *gomock.Controller) *MockClusterCache {
	mock := &MockClusterCache{ctrl: ctrl}
	mock.recorder = &MockClusterCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterCache) EXPECT() *MockClusterCacheMockRecorder {
	return m.recorder
}

// AddIndexer mocks base method.
func (m *MockClusterCache) AddIndexer(arg0 string, arg1 v30.ClusterIndexer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddIndexer", arg0, arg1)
}

// AddIndexer indicates an expected call of AddIndexer.
func (mr *MockClusterCacheMockRecorder) AddIndexer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIndexer", reflect.TypeOf((*MockClusterCache)(nil).AddIndexer), arg0, arg1)
}

// Get mocks base method.
func (m *MockClusterCache) Get(arg0 string) (*v3.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*v3.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClusterCacheMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterCache)(nil).Get), arg0)
}

// GetByIndex mocks base method.
func (m *MockClusterCache) GetByIndex(arg0, arg1 string) ([]*v3.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIndex", arg0, arg1)
	ret0, _ := ret[0].([]*v3.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIndex indicates an expected call of GetByIndex.
func (mr *MockClusterCacheMockRecorder) GetByIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIndex", reflect.TypeOf((*MockClusterCache)(nil).GetByIndex), arg0, arg1)
}

// List mocks base method.
func (m *MockClusterCache) List(arg0 labels.Selector) ([]*v3.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*v3.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockClusterCacheMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterCache)(nil).List), arg0)
}