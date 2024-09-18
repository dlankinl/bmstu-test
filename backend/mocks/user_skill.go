// Code generated by MockGen. DO NOT EDIT.
// Source: user_skill.go
//
// Generated by this command:
//
//	mockgen -source=user_skill.go -destination=../mocks/user_skill.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "ppo/domain"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockIUserSkillRepository is a mock of IUserSkillRepository interface.
type MockIUserSkillRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserSkillRepositoryMockRecorder
}

// MockIUserSkillRepositoryMockRecorder is the mock recorder for MockIUserSkillRepository.
type MockIUserSkillRepositoryMockRecorder struct {
	mock *MockIUserSkillRepository
}

// NewMockIUserSkillRepository creates a new mock instance.
func NewMockIUserSkillRepository(ctrl *gomock.Controller) *MockIUserSkillRepository {
	mock := &MockIUserSkillRepository{ctrl: ctrl}
	mock.recorder = &MockIUserSkillRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserSkillRepository) EXPECT() *MockIUserSkillRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserSkillRepository) Create(arg0 context.Context, arg1 *domain.UserSkill) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIUserSkillRepositoryMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserSkillRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIUserSkillRepository) Delete(arg0 context.Context, arg1 *domain.UserSkill) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIUserSkillRepositoryMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUserSkillRepository)(nil).Delete), arg0, arg1)
}

// GetUserSkillsBySkillId mocks base method.
func (m *MockIUserSkillRepository) GetUserSkillsBySkillId(arg0 context.Context, arg1 uuid.UUID, arg2 int) ([]*domain.UserSkill, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSkillsBySkillId", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*domain.UserSkill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSkillsBySkillId indicates an expected call of GetUserSkillsBySkillId.
func (mr *MockIUserSkillRepositoryMockRecorder) GetUserSkillsBySkillId(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSkillsBySkillId", reflect.TypeOf((*MockIUserSkillRepository)(nil).GetUserSkillsBySkillId), arg0, arg1, arg2)
}

// GetUserSkillsByUserId mocks base method.
func (m *MockIUserSkillRepository) GetUserSkillsByUserId(arg0 context.Context, arg1 uuid.UUID, arg2 int, arg3 bool) ([]*domain.UserSkill, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSkillsByUserId", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*domain.UserSkill)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserSkillsByUserId indicates an expected call of GetUserSkillsByUserId.
func (mr *MockIUserSkillRepositoryMockRecorder) GetUserSkillsByUserId(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSkillsByUserId", reflect.TypeOf((*MockIUserSkillRepository)(nil).GetUserSkillsByUserId), arg0, arg1, arg2, arg3)
}

// MockIUserSkillService is a mock of IUserSkillService interface.
type MockIUserSkillService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserSkillServiceMockRecorder
}

// MockIUserSkillServiceMockRecorder is the mock recorder for MockIUserSkillService.
type MockIUserSkillServiceMockRecorder struct {
	mock *MockIUserSkillService
}

// NewMockIUserSkillService creates a new mock instance.
func NewMockIUserSkillService(ctrl *gomock.Controller) *MockIUserSkillService {
	mock := &MockIUserSkillService{ctrl: ctrl}
	mock.recorder = &MockIUserSkillServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserSkillService) EXPECT() *MockIUserSkillServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserSkillService) Create(arg0 context.Context, arg1 *domain.UserSkill) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIUserSkillServiceMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserSkillService)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIUserSkillService) Delete(arg0 context.Context, arg1 *domain.UserSkill) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIUserSkillServiceMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUserSkillService)(nil).Delete), arg0, arg1)
}

// DeleteSkillsForUser mocks base method.
func (m *MockIUserSkillService) DeleteSkillsForUser(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSkillsForUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSkillsForUser indicates an expected call of DeleteSkillsForUser.
func (mr *MockIUserSkillServiceMockRecorder) DeleteSkillsForUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSkillsForUser", reflect.TypeOf((*MockIUserSkillService)(nil).DeleteSkillsForUser), arg0, arg1)
}

// GetSkillsForUser mocks base method.
func (m *MockIUserSkillService) GetSkillsForUser(arg0 context.Context, arg1 uuid.UUID, arg2 int, arg3 bool) ([]*domain.Skill, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSkillsForUser", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*domain.Skill)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSkillsForUser indicates an expected call of GetSkillsForUser.
func (mr *MockIUserSkillServiceMockRecorder) GetSkillsForUser(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSkillsForUser", reflect.TypeOf((*MockIUserSkillService)(nil).GetSkillsForUser), arg0, arg1, arg2, arg3)
}

// GetUsersForSkill mocks base method.
func (m *MockIUserSkillService) GetUsersForSkill(arg0 context.Context, arg1 uuid.UUID, arg2 int) ([]*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersForSkill", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersForSkill indicates an expected call of GetUsersForSkill.
func (mr *MockIUserSkillServiceMockRecorder) GetUsersForSkill(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersForSkill", reflect.TypeOf((*MockIUserSkillService)(nil).GetUsersForSkill), arg0, arg1, arg2)
}
