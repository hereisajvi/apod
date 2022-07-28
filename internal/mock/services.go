// Code generated by MockGen. DO NOT EDIT.
// Source: iface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/chiefcake/apod/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockPictureService is a mock of PictureService interface.
type MockPictureService struct {
	ctrl     *gomock.Controller
	recorder *MockPictureServiceMockRecorder
}

// MockPictureServiceMockRecorder is the mock recorder for MockPictureService.
type MockPictureServiceMockRecorder struct {
	mock *MockPictureService
}

// NewMockPictureService creates a new mock instance.
func NewMockPictureService(ctrl *gomock.Controller) *MockPictureService {
	mock := &MockPictureService{ctrl: ctrl}
	mock.recorder = &MockPictureServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPictureService) EXPECT() *MockPictureServiceMockRecorder {
	return m.recorder
}

// GetByDate mocks base method.
func (m *MockPictureService) GetByDate(ctx context.Context, date string) (model.Picture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDate", ctx, date)
	ret0, _ := ret[0].(model.Picture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDate indicates an expected call of GetByDate.
func (mr *MockPictureServiceMockRecorder) GetByDate(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDate", reflect.TypeOf((*MockPictureService)(nil).GetByDate), ctx, date)
}

// List mocks base method.
func (m *MockPictureService) List(ctx context.Context) ([]model.Picture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]model.Picture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPictureServiceMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPictureService)(nil).List), ctx)
}