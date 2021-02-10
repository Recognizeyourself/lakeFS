// Code generated by MockGen. DO NOT EDIT.
// Source: hook.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	actions "github.com/treeverse/lakefs/actions"
	io "io"
	reflect "reflect"
)

// MockOutputWriter is a mock of OutputWriter interface
type MockOutputWriter struct {
	ctrl     *gomock.Controller
	recorder *MockOutputWriterMockRecorder
}

// MockOutputWriterMockRecorder is the mock recorder for MockOutputWriter
type MockOutputWriterMockRecorder struct {
	mock *MockOutputWriter
}

// NewMockOutputWriter creates a new mock instance
func NewMockOutputWriter(ctrl *gomock.Controller) *MockOutputWriter {
	mock := &MockOutputWriter{ctrl: ctrl}
	mock.recorder = &MockOutputWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOutputWriter) EXPECT() *MockOutputWriterMockRecorder {
	return m.recorder
}

// OutputWrite mocks base method
func (m *MockOutputWriter) OutputWrite(ctx context.Context, name string, reader io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OutputWrite", ctx, name, reader)
	ret0, _ := ret[0].(error)
	return ret0
}

// OutputWrite indicates an expected call of OutputWrite
func (mr *MockOutputWriterMockRecorder) OutputWrite(ctx, name, reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutputWrite", reflect.TypeOf((*MockOutputWriter)(nil).OutputWrite), ctx, name, reader)
}

// MockHook is a mock of Hook interface
type MockHook struct {
	ctrl     *gomock.Controller
	recorder *MockHookMockRecorder
}

// MockHookMockRecorder is the mock recorder for MockHook
type MockHookMockRecorder struct {
	mock *MockHook
}

// NewMockHook creates a new mock instance
func NewMockHook(ctrl *gomock.Controller) *MockHook {
	mock := &MockHook{ctrl: ctrl}
	mock.recorder = &MockHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHook) EXPECT() *MockHookMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockHook) Run(ctx context.Context, runID string, event actions.Event, writer actions.OutputWriter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx, runID, event, writer)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockHookMockRecorder) Run(ctx, runID, event, writer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockHook)(nil).Run), ctx, runID, event, writer)
}

// MockSource is a mock of Source interface
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *MockSourceMockRecorder
}

// MockSourceMockRecorder is the mock recorder for MockSource
type MockSourceMockRecorder struct {
	mock *MockSource
}

// NewMockSource creates a new mock instance
func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &MockSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSource) EXPECT() *MockSourceMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockSource) List() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockSourceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSource)(nil).List))
}

// Load mocks base method
func (m *MockSource) Load(name string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", name)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load
func (mr *MockSourceMockRecorder) Load(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockSource)(nil).Load), name)
}