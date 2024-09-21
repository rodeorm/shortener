// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rodeorm/shortener/internal/repo (interfaces: AbstractStorage)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/rodeorm/shortener/internal/core"
)

// MockAbstractStorage is a mock of AbstractStorage interface.
type MockAbstractStorage struct {
	ctrl     *gomock.Controller
	recorder *MockAbstractStorageMockRecorder
}

// MockAbstractStorageMockRecorder is the mock recorder for MockAbstractStorage.
type MockAbstractStorageMockRecorder struct {
	mock *MockAbstractStorage
}

// NewMockAbstractStorage creates a new mock instance.
func NewMockAbstractStorage(ctrl *gomock.Controller) *MockAbstractStorage {
	mock := &MockAbstractStorage{ctrl: ctrl}
	mock.recorder = &MockAbstractStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAbstractStorage) EXPECT() *MockAbstractStorageMockRecorder {
	return m.recorder
}

// CloseConnection mocks base method.
func (m *MockAbstractStorage) CloseConnection() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseConnection")
}

// CloseConnection indicates an expected call of CloseConnection.
func (mr *MockAbstractStorageMockRecorder) CloseConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseConnection", reflect.TypeOf((*MockAbstractStorage)(nil).CloseConnection))
}

// DeleteURLs mocks base method.
func (m *MockAbstractStorage) DeleteURLs(arg0 string, arg1 *core.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteURLs", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteURLs indicates an expected call of DeleteURLs.
func (mr *MockAbstractStorageMockRecorder) DeleteURLs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteURLs", reflect.TypeOf((*MockAbstractStorage)(nil).DeleteURLs), arg0, arg1)
}

// InsertURL mocks base method.
func (m *MockAbstractStorage) InsertURL(arg0, arg1 string, arg2 *core.User) (string, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertURL", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// InsertURL indicates an expected call of InsertURL.
func (mr *MockAbstractStorageMockRecorder) InsertURL(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertURL", reflect.TypeOf((*MockAbstractStorage)(nil).InsertURL), arg0, arg1, arg2)
}

// InsertUser mocks base method.
func (m *MockAbstractStorage) InsertUser(arg0 int) (*core.User, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", arg0)
	ret0, _ := ret[0].(*core.User)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockAbstractStorageMockRecorder) InsertUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockAbstractStorage)(nil).InsertUser), arg0)
}

// SelectOriginalURL mocks base method.
func (m *MockAbstractStorage) SelectOriginalURL(arg0 string) (string, bool, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectOriginalURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// SelectOriginalURL indicates an expected call of SelectOriginalURL.
func (mr *MockAbstractStorageMockRecorder) SelectOriginalURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectOriginalURL", reflect.TypeOf((*MockAbstractStorage)(nil).SelectOriginalURL), arg0)
}

// SelectUserURLHistory mocks base method.
func (m *MockAbstractStorage) SelectUserURLHistory(arg0 *core.User) (*[]core.UserURLPair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserURLHistory", arg0)
	ret0, _ := ret[0].(*[]core.UserURLPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserURLHistory indicates an expected call of SelectUserURLHistory.
func (mr *MockAbstractStorageMockRecorder) SelectUserURLHistory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserURLHistory", reflect.TypeOf((*MockAbstractStorage)(nil).SelectUserURLHistory), arg0)
}
