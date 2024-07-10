// Code generated by MockGen. DO NOT EDIT.
// Source: get.go

// Package getapi is a generated GoMock package.
package getapi

import (
	models "eratani/TestCase3/c/storage/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockResultData is a mock of ResultData interface.
type MockResultData struct {
	ctrl     *gomock.Controller
	recorder *MockResultDataMockRecorder
}

// MockResultDataMockRecorder is the mock recorder for MockResultData.
type MockResultDataMockRecorder struct {
	mock *MockResultData
}

// NewMockResultData creates a new mock instance.
func NewMockResultData(ctrl *gomock.Controller) *MockResultData {
	mock := &MockResultData{ctrl: ctrl}
	mock.recorder = &MockResultDataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResultData) EXPECT() *MockResultDataMockRecorder {
	return m.recorder
}

// GetDataResponse mocks base method.
func (m *MockResultData) GetDataResponse(req models.RequestData) ([]models.ResponseQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataResponse", req)
	ret0, _ := ret[0].([]models.ResponseQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataResponse indicates an expected call of GetDataResponse.
func (mr *MockResultDataMockRecorder) GetDataResponse(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataResponse", reflect.TypeOf((*MockResultData)(nil).GetDataResponse), req)
}