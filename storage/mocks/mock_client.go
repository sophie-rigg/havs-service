// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/sophie-rigg/havs-service/models"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetEquipmentItem mocks base method.
func (m *MockClient) GetEquipmentItem(id string) (*models.EquipmentItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEquipmentItem", id)
	ret0, _ := ret[0].(*models.EquipmentItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEquipmentItem indicates an expected call of GetEquipmentItem.
func (mr *MockClientMockRecorder) GetEquipmentItem(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEquipmentItem", reflect.TypeOf((*MockClient)(nil).GetEquipmentItem), id)
}

// GetExposure mocks base method.
func (m *MockClient) GetExposure(id string) (*models.Exposure, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExposure", id)
	ret0, _ := ret[0].(*models.Exposure)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExposure indicates an expected call of GetExposure.
func (mr *MockClientMockRecorder) GetExposure(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExposure", reflect.TypeOf((*MockClient)(nil).GetExposure), id)
}

// GetExposures mocks base method.
func (m *MockClient) GetExposures() ([]*models.Exposure, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExposures")
	ret0, _ := ret[0].([]*models.Exposure)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExposures indicates an expected call of GetExposures.
func (mr *MockClientMockRecorder) GetExposures() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExposures", reflect.TypeOf((*MockClient)(nil).GetExposures))
}

// GetExposuresByUserID mocks base method.
func (m *MockClient) GetExposuresByUserID(userID string, startTime, endTime time.Time) ([]*models.Exposure, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExposuresByUserID", userID, startTime, endTime)
	ret0, _ := ret[0].([]*models.Exposure)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExposuresByUserID indicates an expected call of GetExposuresByUserID.
func (mr *MockClientMockRecorder) GetExposuresByUserID(userID, startTime, endTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExposuresByUserID", reflect.TypeOf((*MockClient)(nil).GetExposuresByUserID), userID, startTime, endTime)
}

// GetUser mocks base method.
func (m *MockClient) GetUser(id string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockClientMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockClient)(nil).GetUser), id)
}

// InsertExposure mocks base method.
func (m *MockClient) InsertExposure(exposure *models.Exposure) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertExposure", exposure)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertExposure indicates an expected call of InsertExposure.
func (mr *MockClientMockRecorder) InsertExposure(exposure interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertExposure", reflect.TypeOf((*MockClient)(nil).InsertExposure), exposure)
}
