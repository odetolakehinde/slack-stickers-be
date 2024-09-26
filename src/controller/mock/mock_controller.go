// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go
//
// Generated by this command:
//
//	mockgen -source controller.go -destination ./mock/mock_controller.go -package mock Operations
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/odetolakehinde/slack-stickers-be/src/model"
	middleware "github.com/odetolakehinde/slack-stickers-be/src/pkg/middleware"
	gomock "go.uber.org/mock/gomock"
)

// MockOperations is a mock of Operations interface.
type MockOperations struct {
	ctrl     *gomock.Controller
	recorder *MockOperationsMockRecorder
}

// MockOperationsMockRecorder is the mock recorder for MockOperations.
type MockOperationsMockRecorder struct {
	mock *MockOperations
}

// NewMockOperations creates a new mock instance.
func NewMockOperations(ctrl *gomock.Controller) *MockOperations {
	mock := &MockOperations{ctrl: ctrl}
	mock.recorder = &MockOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperations) EXPECT() *MockOperationsMockRecorder {
	return m.recorder
}

// Middleware mocks base method.
func (m *MockOperations) Middleware() *middleware.Middleware {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Middleware")
	ret0, _ := ret[0].(*middleware.Middleware)
	return ret0
}

// Middleware indicates an expected call of Middleware.
func (mr *MockOperationsMockRecorder) Middleware() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Middleware", reflect.TypeOf((*MockOperations)(nil).Middleware))
}

// SaveAuthDetails mocks base method.
func (m *MockOperations) SaveAuthDetails(ctx context.Context, authDetails model.SlackAuthDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAuthDetails", ctx, authDetails)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAuthDetails indicates an expected call of SaveAuthDetails.
func (mr *MockOperationsMockRecorder) SaveAuthDetails(ctx, authDetails any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAuthDetails", reflect.TypeOf((*MockOperations)(nil).SaveAuthDetails), ctx, authDetails)
}

// SearchByTag mocks base method.
func (m *MockOperations) SearchByTag(ctx context.Context, triggerID, tag, countToReturn, channelID, teamID string, externalViewID *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByTag", ctx, triggerID, tag, countToReturn, channelID, teamID, externalViewID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchByTag indicates an expected call of SearchByTag.
func (mr *MockOperationsMockRecorder) SearchByTag(ctx, triggerID, tag, countToReturn, channelID, teamID, externalViewID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByTag", reflect.TypeOf((*MockOperations)(nil).SearchByTag), ctx, triggerID, tag, countToReturn, channelID, teamID, externalViewID)
}

// SendSticker mocks base method.
func (m *MockOperations) SendSticker(ctx context.Context, channelID, imageURL, teamID, threadTs string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSticker", ctx, channelID, imageURL, teamID, threadTs)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSticker indicates an expected call of SendSticker.
func (mr *MockOperationsMockRecorder) SendSticker(ctx, channelID, imageURL, teamID, threadTs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSticker", reflect.TypeOf((*MockOperations)(nil).SendSticker), ctx, channelID, imageURL, teamID, threadTs)
}

// ShowSearchModal mocks base method.
func (m *MockOperations) ShowSearchModal(ctx context.Context, triggerID, channelID, teamID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowSearchModal", ctx, triggerID, channelID, teamID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShowSearchModal indicates an expected call of ShowSearchModal.
func (mr *MockOperationsMockRecorder) ShowSearchModal(ctx, triggerID, channelID, teamID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowSearchModal", reflect.TypeOf((*MockOperations)(nil).ShowSearchModal), ctx, triggerID, channelID, teamID)
}
