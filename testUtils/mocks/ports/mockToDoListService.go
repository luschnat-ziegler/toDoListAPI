// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luschnat-ziegler/toDoListAPI/core/ports (interfaces: ToDoListService)

// Package ports is a generated GoMock package.
package ports

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/luschnat-ziegler/toDoListAPI/core/domain"
	errs "github.com/luschnat-ziegler/toDoListAPI/errs"
	reflect "reflect"
)

// MockToDoListService is a mock of ToDoListService interface
type MockToDoListService struct {
	ctrl     *gomock.Controller
	recorder *MockToDoListServiceMockRecorder
}

// MockToDoListServiceMockRecorder is the mock recorder for MockToDoListService
type MockToDoListServiceMockRecorder struct {
	mock *MockToDoListService
}

// NewMockToDoListService creates a new mock instance
func NewMockToDoListService(ctrl *gomock.Controller) *MockToDoListService {
	mock := &MockToDoListService{ctrl: ctrl}
	mock.recorder = &MockToDoListServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockToDoListService) EXPECT() *MockToDoListServiceMockRecorder {
	return m.recorder
}

// DeleteList mocks base method
func (m *MockToDoListService) DeleteListById(arg0 string) *errs.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteListById", arg0)
	ret0, _ := ret[0].(*errs.AppError)
	return ret0
}

// DeleteList indicates an expected call of DeleteList
func (mr *MockToDoListServiceMockRecorder) DeleteList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteListById", reflect.TypeOf((*MockToDoListService)(nil).DeleteListById), arg0)
}

// GetAllLists mocks base method
func (m *MockToDoListService) GetAllLists() (*[]domain.ToDoList, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLists")
	ret0, _ := ret[0].(*[]domain.ToDoList)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetAllLists indicates an expected call of GetAllLists
func (mr *MockToDoListServiceMockRecorder) GetAllLists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLists", reflect.TypeOf((*MockToDoListService)(nil).GetAllLists))
}

// GetOneListById mocks base method
func (m *MockToDoListService) GetOneListById(arg0 string) (*domain.ToDoList, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneListById", arg0)
	ret0, _ := ret[0].(*domain.ToDoList)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetOneListById indicates an expected call of GetOneListById
func (mr *MockToDoListServiceMockRecorder) GetOneListById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneListById", reflect.TypeOf((*MockToDoListService)(nil).GetOneListById), arg0)
}

// SaveList mocks base method
func (m *MockToDoListService) SaveList(arg0 domain.ToDoList) (*domain.ToDoList, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveList", arg0)
	ret0, _ := ret[0].(*domain.ToDoList)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// SaveList indicates an expected call of SaveList
func (mr *MockToDoListServiceMockRecorder) SaveList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveList", reflect.TypeOf((*MockToDoListService)(nil).SaveList), arg0)
}

// UpdateOneListById mocks base method
func (m *MockToDoListService) UpdateOneListById(arg0 string, arg1 domain.ToDoList) (*domain.ToDoList, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOneListById", arg0, arg1)
	ret0, _ := ret[0].(*domain.ToDoList)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// UpdateOneListById indicates an expected call of UpdateOneListById
func (mr *MockToDoListServiceMockRecorder) UpdateOneListById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOneListById", reflect.TypeOf((*MockToDoListService)(nil).UpdateOneListById), arg0, arg1)
}
