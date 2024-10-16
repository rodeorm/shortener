// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rodeorm/shortener/internal/repo (interfaces: Storager)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/rodeorm/shortener/internal/core"
)

// MockStorager is a mock of Storager interface.
type MockStorager struct {
	ctrl     *gomock.Controller
	recorder *MockStoragerMockRecorder
}

// MockStoragerMockRecorder is the mock recorder for MockStorager.
type MockStoragerMockRecorder struct {
	mock *MockStorager
}

// NewMockStorager creates a new mock instance.
func NewMockStorager(ctrl *gomock.Controller) *MockStorager {
	mock := &MockStorager{ctrl: ctrl}
	mock.recorder = &MockStoragerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorager) EXPECT() *MockStoragerMockRecorder {
	return m.recorder
}

// CloseConnection mocks base method.
func (m *MockStorager) CloseConnection() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseConnection")
}

// CloseConnection indicates an expected call of CloseConnection.
func (mr *MockStoragerMockRecorder) CloseConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseConnection", reflect.TypeOf((*MockStorager)(nil).CloseConnection))
}

// DeleteURLs mocks base method.
func (m *MockStorager) DeleteURLs(arg0 []core.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteURLs", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteURLs indicates an expected call of DeleteURLs.
func (mr *MockStoragerMockRecorder) DeleteURLs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteURLs", reflect.TypeOf((*MockStorager)(nil).DeleteURLs), arg0)
}

// InsertURL mocks base method.
func (m *MockStorager) InsertURL(arg0, arg1 string, arg2 *core.User) (*core.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertURL", arg0, arg1, arg2)
	ret0, _ := ret[0].(*core.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertURL indicates an expected call of InsertURL.
func (mr *MockStoragerMockRecorder) InsertURL(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertURL", reflect.TypeOf((*MockStorager)(nil).InsertURL), arg0, arg1, arg2)
}

// InsertUser mocks base method.
func (m *MockStorager) InsertUser(arg0 int) (*core.User, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", arg0)
	ret0, _ := ret[0].(*core.User)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockStoragerMockRecorder) InsertUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockStorager)(nil).InsertUser), arg0)
}

// PingDB mocks base method.
func (m *MockStorager) PingDB() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingDB")
	ret0, _ := ret[0].(error)
	return ret0
}

// PingDB indicates an expected call of PingDB.
func (mr *MockStoragerMockRecorder) PingDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingDB", reflect.TypeOf((*MockStorager)(nil).PingDB))
}

// SelectOriginalURL mocks base method.
func (m *MockStorager) SelectOriginalURL(arg0 string) (*core.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectOriginalURL", arg0)
	ret0, _ := ret[0].(*core.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectOriginalURL indicates an expected call of SelectOriginalURL.
func (mr *MockStoragerMockRecorder) SelectOriginalURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectOriginalURL", reflect.TypeOf((*MockStorager)(nil).SelectOriginalURL), arg0)
}

// SelectUserURLHistory mocks base method.
func (m *MockStorager) SelectUserURLHistory(arg0 *core.User) ([]core.UserURLPair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserURLHistory", arg0)
	ret0, _ := ret[0].([]core.UserURLPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserURLHistory indicates an expected call of SelectUserURLHistory.
func (mr *MockStoragerMockRecorder) SelectUserURLHistory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserURLHistory", reflect.TypeOf((*MockStorager)(nil).SelectUserURLHistory), arg0)
}