// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/endpoint/rest/types.go

// Package rest is a generated GoMock package.
package rest

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthResourceEndpoint is a mock of AuthResourceEndpoint interface.
type MockAuthResourceEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockAuthResourceEndpointMockRecorder
}

// MockAuthResourceEndpointMockRecorder is the mock recorder for MockAuthResourceEndpoint.
type MockAuthResourceEndpointMockRecorder struct {
	mock *MockAuthResourceEndpoint
}

// NewMockAuthResourceEndpoint creates a new mock instance.
func NewMockAuthResourceEndpoint(ctrl *gomock.Controller) *MockAuthResourceEndpoint {
	mock := &MockAuthResourceEndpoint{ctrl: ctrl}
	mock.recorder = &MockAuthResourceEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthResourceEndpoint) EXPECT() *MockAuthResourceEndpointMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuthResourceEndpoint) Create(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", ctx)
}

// Create indicates an expected call of Create.
func (mr *MockAuthResourceEndpointMockRecorder) Create(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthResourceEndpoint)(nil).Create), ctx)
}

// Delete mocks base method.
func (m *MockAuthResourceEndpoint) Delete(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", ctx)
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthResourceEndpointMockRecorder) Delete(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthResourceEndpoint)(nil).Delete), ctx)
}

// FindAll mocks base method.
func (m *MockAuthResourceEndpoint) FindAll(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindAll", ctx)
}

// FindAll indicates an expected call of FindAll.
func (mr *MockAuthResourceEndpointMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAuthResourceEndpoint)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockAuthResourceEndpoint) FindById(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindById", ctx)
}

// FindById indicates an expected call of FindById.
func (mr *MockAuthResourceEndpointMockRecorder) FindById(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockAuthResourceEndpoint)(nil).FindById), ctx)
}

// Update mocks base method.
func (m *MockAuthResourceEndpoint) Update(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", ctx)
}

// Update indicates an expected call of Update.
func (mr *MockAuthResourceEndpointMockRecorder) Update(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuthResourceEndpoint)(nil).Update), ctx)
}

// MockAuthRoleEndpoint is a mock of AuthRoleEndpoint interface.
type MockAuthRoleEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRoleEndpointMockRecorder
}

// MockAuthRoleEndpointMockRecorder is the mock recorder for MockAuthRoleEndpoint.
type MockAuthRoleEndpointMockRecorder struct {
	mock *MockAuthRoleEndpoint
}

// NewMockAuthRoleEndpoint creates a new mock instance.
func NewMockAuthRoleEndpoint(ctrl *gomock.Controller) *MockAuthRoleEndpoint {
	mock := &MockAuthRoleEndpoint{ctrl: ctrl}
	mock.recorder = &MockAuthRoleEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRoleEndpoint) EXPECT() *MockAuthRoleEndpointMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuthRoleEndpoint) Create(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", ctx)
}

// Create indicates an expected call of Create.
func (mr *MockAuthRoleEndpointMockRecorder) Create(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthRoleEndpoint)(nil).Create), ctx)
}

// Delete mocks base method.
func (m *MockAuthRoleEndpoint) Delete(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", ctx)
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthRoleEndpointMockRecorder) Delete(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthRoleEndpoint)(nil).Delete), ctx)
}

// FindAll mocks base method.
func (m *MockAuthRoleEndpoint) FindAll(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindAll", ctx)
}

// FindAll indicates an expected call of FindAll.
func (mr *MockAuthRoleEndpointMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAuthRoleEndpoint)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockAuthRoleEndpoint) FindById(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindById", ctx)
}

// FindById indicates an expected call of FindById.
func (mr *MockAuthRoleEndpointMockRecorder) FindById(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockAuthRoleEndpoint)(nil).FindById), ctx)
}

// Update mocks base method.
func (m *MockAuthRoleEndpoint) Update(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", ctx)
}

// Update indicates an expected call of Update.
func (mr *MockAuthRoleEndpointMockRecorder) Update(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuthRoleEndpoint)(nil).Update), ctx)
}

// MockAuthAccessControlListEndpoint is a mock of AuthAccessControlListEndpoint interface.
type MockAuthAccessControlListEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockAuthAccessControlListEndpointMockRecorder
}

// MockAuthAccessControlListEndpointMockRecorder is the mock recorder for MockAuthAccessControlListEndpoint.
type MockAuthAccessControlListEndpointMockRecorder struct {
	mock *MockAuthAccessControlListEndpoint
}

// NewMockAuthAccessControlListEndpoint creates a new mock instance.
func NewMockAuthAccessControlListEndpoint(ctrl *gomock.Controller) *MockAuthAccessControlListEndpoint {
	mock := &MockAuthAccessControlListEndpoint{ctrl: ctrl}
	mock.recorder = &MockAuthAccessControlListEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthAccessControlListEndpoint) EXPECT() *MockAuthAccessControlListEndpointMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuthAccessControlListEndpoint) Create(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", ctx)
}

// Create indicates an expected call of Create.
func (mr *MockAuthAccessControlListEndpointMockRecorder) Create(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthAccessControlListEndpoint)(nil).Create), ctx)
}

// Delete mocks base method.
func (m *MockAuthAccessControlListEndpoint) Delete(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", ctx)
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthAccessControlListEndpointMockRecorder) Delete(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthAccessControlListEndpoint)(nil).Delete), ctx)
}

// FindAll mocks base method.
func (m *MockAuthAccessControlListEndpoint) FindAll(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindAll", ctx)
}

// FindAll indicates an expected call of FindAll.
func (mr *MockAuthAccessControlListEndpointMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAuthAccessControlListEndpoint)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockAuthAccessControlListEndpoint) FindById(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindById", ctx)
}

// FindById indicates an expected call of FindById.
func (mr *MockAuthAccessControlListEndpointMockRecorder) FindById(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockAuthAccessControlListEndpoint)(nil).FindById), ctx)
}

// Update mocks base method.
func (m *MockAuthAccessControlListEndpoint) Update(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", ctx)
}

// Update indicates an expected call of Update.
func (mr *MockAuthAccessControlListEndpointMockRecorder) Update(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuthAccessControlListEndpoint)(nil).Update), ctx)
}

// MockAuthUserEndpoint is a mock of AuthUserEndpoint interface.
type MockAuthUserEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUserEndpointMockRecorder
}

// MockAuthUserEndpointMockRecorder is the mock recorder for MockAuthUserEndpoint.
type MockAuthUserEndpointMockRecorder struct {
	mock *MockAuthUserEndpoint
}

// NewMockAuthUserEndpoint creates a new mock instance.
func NewMockAuthUserEndpoint(ctrl *gomock.Controller) *MockAuthUserEndpoint {
	mock := &MockAuthUserEndpoint{ctrl: ctrl}
	mock.recorder = &MockAuthUserEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthUserEndpoint) EXPECT() *MockAuthUserEndpointMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuthUserEndpoint) Create(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", ctx)
}

// Create indicates an expected call of Create.
func (mr *MockAuthUserEndpointMockRecorder) Create(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthUserEndpoint)(nil).Create), ctx)
}

// Delete mocks base method.
func (m *MockAuthUserEndpoint) Delete(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", ctx)
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthUserEndpointMockRecorder) Delete(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthUserEndpoint)(nil).Delete), ctx)
}

// FindAll mocks base method.
func (m *MockAuthUserEndpoint) FindAll(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindAll", ctx)
}

// FindAll indicates an expected call of FindAll.
func (mr *MockAuthUserEndpointMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAuthUserEndpoint)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockAuthUserEndpoint) FindById(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindById", ctx)
}

// FindById indicates an expected call of FindById.
func (mr *MockAuthUserEndpointMockRecorder) FindById(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockAuthUserEndpoint)(nil).FindById), ctx)
}

// Update mocks base method.
func (m *MockAuthUserEndpoint) Update(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", ctx)
}

// Update indicates an expected call of Update.
func (mr *MockAuthUserEndpointMockRecorder) Update(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuthUserEndpoint)(nil).Update), ctx)
}

// MockAuthPrincipalEndpoint is a mock of AuthPrincipalEndpoint interface.
type MockAuthPrincipalEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockAuthPrincipalEndpointMockRecorder
}

// MockAuthPrincipalEndpointMockRecorder is the mock recorder for MockAuthPrincipalEndpoint.
type MockAuthPrincipalEndpointMockRecorder struct {
	mock *MockAuthPrincipalEndpoint
}

// NewMockAuthPrincipalEndpoint creates a new mock instance.
func NewMockAuthPrincipalEndpoint(ctrl *gomock.Controller) *MockAuthPrincipalEndpoint {
	mock := &MockAuthPrincipalEndpoint{ctrl: ctrl}
	mock.recorder = &MockAuthPrincipalEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthPrincipalEndpoint) EXPECT() *MockAuthPrincipalEndpointMockRecorder {
	return m.recorder
}

// GetCurrentPrincipal mocks base method.
func (m *MockAuthPrincipalEndpoint) GetCurrentPrincipal(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetCurrentPrincipal", ctx)
}

// GetCurrentPrincipal indicates an expected call of GetCurrentPrincipal.
func (mr *MockAuthPrincipalEndpointMockRecorder) GetCurrentPrincipal(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentPrincipal", reflect.TypeOf((*MockAuthPrincipalEndpoint)(nil).GetCurrentPrincipal), ctx)
}
