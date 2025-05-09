// Code generated by MockGen. DO NOT EDIT.
// Source: language.go

// Package model_mock is a generated GoMock package.
package model_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/ppdx999/kyopro/internal/domain/model"
)

// MockLanguage is a mock of Language interface.
type MockLanguage struct {
	ctrl     *gomock.Controller
	recorder *MockLanguageMockRecorder
}

// MockLanguageMockRecorder is the mock recorder for MockLanguage.
type MockLanguageMockRecorder struct {
	mock *MockLanguage
}

// NewMockLanguage creates a new mock instance.
func NewMockLanguage(ctrl *gomock.Controller) *MockLanguage {
	mock := &MockLanguage{ctrl: ctrl}
	mock.recorder = &MockLanguageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLanguage) EXPECT() *MockLanguageMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockLanguage) Build(p *model.Pipeline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockLanguageMockRecorder) Build(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockLanguage)(nil).Build), p)
}

// Clean mocks base method.
func (m *MockLanguage) Clean(p *model.Pipeline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clean", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clean indicates an expected call of Clean.
func (mr *MockLanguageMockRecorder) Clean(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clean", reflect.TypeOf((*MockLanguage)(nil).Clean), p)
}

// Name mocks base method.
func (m *MockLanguage) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockLanguageMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockLanguage)(nil).Name))
}

// Run mocks base method.
func (m *MockLanguage) Run(p *model.Pipeline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockLanguageMockRecorder) Run(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockLanguage)(nil).Run), p)
}

// SourceCode mocks base method.
func (m *MockLanguage) SourceCode() *model.SourceCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SourceCode")
	ret0, _ := ret[0].(*model.SourceCode)
	return ret0
}

// SourceCode indicates an expected call of SourceCode.
func (mr *MockLanguageMockRecorder) SourceCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SourceCode", reflect.TypeOf((*MockLanguage)(nil).SourceCode))
}
