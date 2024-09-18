// Code generated by MockGen. DO NOT EDIT.
// Source: auth.go
//
// Generated by this command:
//
//	mockgen -source=auth.go -destination=../mocks/auth.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "ppo/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIAuthRepository is a mock of IAuthRepository interface.
type MockIAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthRepositoryMockRecorder
}

// MockIAuthRepositoryMockRecorder is the mock recorder for MockIAuthRepository.
type MockIAuthRepositoryMockRecorder struct {
	mock *MockIAuthRepository
}

// NewMockIAuthRepository creates a new mock instance.
func NewMockIAuthRepository(ctrl *gomock.Controller) *MockIAuthRepository {
	mock := &MockIAuthRepository{ctrl: ctrl}
	mock.recorder = &MockIAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthRepository) EXPECT() *MockIAuthRepositoryMockRecorder {
	return m.recorder
}

// GetByUsername mocks base method.
func (m *MockIAuthRepository) GetByUsername(arg0 context.Context, arg1 string) (*domain.UserAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", arg0, arg1)
	ret0, _ := ret[0].(*domain.UserAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockIAuthRepositoryMockRecorder) GetByUsername(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockIAuthRepository)(nil).GetByUsername), arg0, arg1)
}

// Register mocks base method.
func (m *MockIAuthRepository) Register(arg0 context.Context, arg1 *domain.UserAuth) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockIAuthRepositoryMockRecorder) Register(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIAuthRepository)(nil).Register), arg0, arg1)
}

// MockIAuthService is a mock of IAuthService interface.
type MockIAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthServiceMockRecorder
}

// MockIAuthServiceMockRecorder is the mock recorder for MockIAuthService.
type MockIAuthServiceMockRecorder struct {
	mock *MockIAuthService
}

// NewMockIAuthService creates a new mock instance.
func NewMockIAuthService(ctrl *gomock.Controller) *MockIAuthService {
	mock := &MockIAuthService{ctrl: ctrl}
	mock.recorder = &MockIAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthService) EXPECT() *MockIAuthServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockIAuthService) Login(arg0 context.Context, arg1 *domain.UserAuth) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockIAuthServiceMockRecorder) Login(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockIAuthService)(nil).Login), arg0, arg1)
}

// Register mocks base method.
func (m *MockIAuthService) Register(arg0 context.Context, arg1 *domain.UserAuth) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockIAuthServiceMockRecorder) Register(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIAuthService)(nil).Register), arg0, arg1)
}
