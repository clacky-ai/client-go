// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: GroupLabelsServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=group_labels_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupLabelsServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockGroupLabelsServiceInterface is a mock of GroupLabelsServiceInterface interface.
type MockGroupLabelsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockGroupLabelsServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockGroupLabelsServiceInterfaceMockRecorder is the mock recorder for MockGroupLabelsServiceInterface.
type MockGroupLabelsServiceInterfaceMockRecorder struct {
	mock *MockGroupLabelsServiceInterface
}

// NewMockGroupLabelsServiceInterface creates a new mock instance.
func NewMockGroupLabelsServiceInterface(ctrl *gomock.Controller) *MockGroupLabelsServiceInterface {
	mock := &MockGroupLabelsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockGroupLabelsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupLabelsServiceInterface) EXPECT() *MockGroupLabelsServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateGroupLabel mocks base method.
func (m *MockGroupLabelsServiceInterface) CreateGroupLabel(gid any, opt *gitlab.CreateGroupLabelOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateGroupLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupLabel)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateGroupLabel indicates an expected call of CreateGroupLabel.
func (mr *MockGroupLabelsServiceInterfaceMockRecorder) CreateGroupLabel(gid, opt any, options ...any) *MockGroupLabelsServiceInterfaceCreateGroupLabelCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroupLabel", reflect.TypeOf((*MockGroupLabelsServiceInterface)(nil).CreateGroupLabel), varargs...)
	return &MockGroupLabelsServiceInterfaceCreateGroupLabelCall{Call: call}
}

// MockGroupLabelsServiceInterfaceCreateGroupLabelCall wrap *gomock.Call
type MockGroupLabelsServiceInterfaceCreateGroupLabelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupLabelsServiceInterfaceCreateGroupLabelCall) Return(arg0 *gitlab.GroupLabel, arg1 *gitlab.Response, arg2 error) *MockGroupLabelsServiceInterfaceCreateGroupLabelCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupLabelsServiceInterfaceCreateGroupLabelCall) Do(f func(any, *gitlab.CreateGroupLabelOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error)) *MockGroupLabelsServiceInterfaceCreateGroupLabelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupLabelsServiceInterfaceCreateGroupLabelCall) DoAndReturn(f func(any, *gitlab.CreateGroupLabelOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error)) *MockGroupLabelsServiceInterfaceCreateGroupLabelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteGroupLabel mocks base method.
func (m *MockGroupLabelsServiceInterface) DeleteGroupLabel(gid, lid any, opt *gitlab.DeleteGroupLabelOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, lid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteGroupLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGroupLabel indicates an expected call of DeleteGroupLabel.
func (mr *MockGroupLabelsServiceInterfaceMockRecorder) DeleteGroupLabel(gid, lid, opt any, options ...any) *MockGroupLabelsServiceInterfaceDeleteGroupLabelCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, lid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupLabel", reflect.TypeOf((*MockGroupLabelsServiceInterface)(nil).DeleteGroupLabel), varargs...)
	return &MockGroupLabelsServiceInterfaceDeleteGroupLabelCall{Call: call}
}

// MockGroupLabelsServiceInterfaceDeleteGroupLabelCall wrap *gomock.Call
type MockGroupLabelsServiceInterfaceDeleteGroupLabelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupLabelsServiceInterfaceDeleteGroupLabelCall) Return(arg0 *gitlab.Response, arg1 error) *MockGroupLabelsServiceInterfaceDeleteGroupLabelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupLabelsServiceInterfaceDeleteGroupLabelCall) Do(f func(any, any, *gitlab.DeleteGroupLabelOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockGroupLabelsServiceInterfaceDeleteGroupLabelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupLabelsServiceInterfaceDeleteGroupLabelCall) DoAndReturn(f func(any, any, *gitlab.DeleteGroupLabelOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockGroupLabelsServiceInterfaceDeleteGroupLabelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetGroupLabel mocks base method.
func (m *MockGroupLabelsServiceInterface) GetGroupLabel(gid, lid any, options ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, lid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupLabel)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroupLabel indicates an expected call of GetGroupLabel.
func (mr *MockGroupLabelsServiceInterfaceMockRecorder) GetGroupLabel(gid, lid any, options ...any) *MockGroupLabelsServiceInterfaceGetGroupLabelCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, lid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupLabel", reflect.TypeOf((*MockGroupLabelsServiceInterface)(nil).GetGroupLabel), varargs...)
	return &MockGroupLabelsServiceInterfaceGetGroupLabelCall{Call: call}
}

// MockGroupLabelsServiceInterfaceGetGroupLabelCall wrap *gomock.Call
type MockGroupLabelsServiceInterfaceGetGroupLabelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupLabelsServiceInterfaceGetGroupLabelCall) Return(arg0 *gitlab.GroupLabel, arg1 *gitlab.Response, arg2 error) *MockGroupLabelsServiceInterfaceGetGroupLabelCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupLabelsServiceInterfaceGetGroupLabelCall) Do(f func(any, any, ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error)) *MockGroupLabelsServiceInterfaceGetGroupLabelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupLabelsServiceInterfaceGetGroupLabelCall) DoAndReturn(f func(any, any, ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error)) *MockGroupLabelsServiceInterfaceGetGroupLabelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListGroupLabels mocks base method.
func (m *MockGroupLabelsServiceInterface) ListGroupLabels(gid any, opt *gitlab.ListGroupLabelsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.GroupLabel, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupLabels", varargs...)
	ret0, _ := ret[0].([]*gitlab.GroupLabel)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGroupLabels indicates an expected call of ListGroupLabels.
func (mr *MockGroupLabelsServiceInterfaceMockRecorder) ListGroupLabels(gid, opt any, options ...any) *MockGroupLabelsServiceInterfaceListGroupLabelsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupLabels", reflect.TypeOf((*MockGroupLabelsServiceInterface)(nil).ListGroupLabels), varargs...)
	return &MockGroupLabelsServiceInterfaceListGroupLabelsCall{Call: call}
}

// MockGroupLabelsServiceInterfaceListGroupLabelsCall wrap *gomock.Call
type MockGroupLabelsServiceInterfaceListGroupLabelsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupLabelsServiceInterfaceListGroupLabelsCall) Return(arg0 []*gitlab.GroupLabel, arg1 *gitlab.Response, arg2 error) *MockGroupLabelsServiceInterfaceListGroupLabelsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupLabelsServiceInterfaceListGroupLabelsCall) Do(f func(any, *gitlab.ListGroupLabelsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.GroupLabel, *gitlab.Response, error)) *MockGroupLabelsServiceInterfaceListGroupLabelsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupLabelsServiceInterfaceListGroupLabelsCall) DoAndReturn(f func(any, *gitlab.ListGroupLabelsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.GroupLabel, *gitlab.Response, error)) *MockGroupLabelsServiceInterfaceListGroupLabelsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SubscribeToGroupLabel mocks base method.
func (m *MockGroupLabelsServiceInterface) SubscribeToGroupLabel(gid, lid any, options ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, lid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeToGroupLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupLabel)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SubscribeToGroupLabel indicates an expected call of SubscribeToGroupLabel.
func (mr *MockGroupLabelsServiceInterfaceMockRecorder) SubscribeToGroupLabel(gid, lid any, options ...any) *MockGroupLabelsServiceInterfaceSubscribeToGroupLabelCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, lid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToGroupLabel", reflect.TypeOf((*MockGroupLabelsServiceInterface)(nil).SubscribeToGroupLabel), varargs...)
	return &MockGroupLabelsServiceInterfaceSubscribeToGroupLabelCall{Call: call}
}

// MockGroupLabelsServiceInterfaceSubscribeToGroupLabelCall wrap *gomock.Call
type MockGroupLabelsServiceInterfaceSubscribeToGroupLabelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupLabelsServiceInterfaceSubscribeToGroupLabelCall) Return(arg0 *gitlab.GroupLabel, arg1 *gitlab.Response, arg2 error) *MockGroupLabelsServiceInterfaceSubscribeToGroupLabelCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupLabelsServiceInterfaceSubscribeToGroupLabelCall) Do(f func(any, any, ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error)) *MockGroupLabelsServiceInterfaceSubscribeToGroupLabelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupLabelsServiceInterfaceSubscribeToGroupLabelCall) DoAndReturn(f func(any, any, ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error)) *MockGroupLabelsServiceInterfaceSubscribeToGroupLabelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UnsubscribeFromGroupLabel mocks base method.
func (m *MockGroupLabelsServiceInterface) UnsubscribeFromGroupLabel(gid, lid any, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, lid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsubscribeFromGroupLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnsubscribeFromGroupLabel indicates an expected call of UnsubscribeFromGroupLabel.
func (mr *MockGroupLabelsServiceInterfaceMockRecorder) UnsubscribeFromGroupLabel(gid, lid any, options ...any) *MockGroupLabelsServiceInterfaceUnsubscribeFromGroupLabelCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, lid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeFromGroupLabel", reflect.TypeOf((*MockGroupLabelsServiceInterface)(nil).UnsubscribeFromGroupLabel), varargs...)
	return &MockGroupLabelsServiceInterfaceUnsubscribeFromGroupLabelCall{Call: call}
}

// MockGroupLabelsServiceInterfaceUnsubscribeFromGroupLabelCall wrap *gomock.Call
type MockGroupLabelsServiceInterfaceUnsubscribeFromGroupLabelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupLabelsServiceInterfaceUnsubscribeFromGroupLabelCall) Return(arg0 *gitlab.Response, arg1 error) *MockGroupLabelsServiceInterfaceUnsubscribeFromGroupLabelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupLabelsServiceInterfaceUnsubscribeFromGroupLabelCall) Do(f func(any, any, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockGroupLabelsServiceInterfaceUnsubscribeFromGroupLabelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupLabelsServiceInterfaceUnsubscribeFromGroupLabelCall) DoAndReturn(f func(any, any, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockGroupLabelsServiceInterfaceUnsubscribeFromGroupLabelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateGroupLabel mocks base method.
func (m *MockGroupLabelsServiceInterface) UpdateGroupLabel(gid, lid any, opt *gitlab.UpdateGroupLabelOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, lid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateGroupLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupLabel)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateGroupLabel indicates an expected call of UpdateGroupLabel.
func (mr *MockGroupLabelsServiceInterfaceMockRecorder) UpdateGroupLabel(gid, lid, opt any, options ...any) *MockGroupLabelsServiceInterfaceUpdateGroupLabelCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, lid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroupLabel", reflect.TypeOf((*MockGroupLabelsServiceInterface)(nil).UpdateGroupLabel), varargs...)
	return &MockGroupLabelsServiceInterfaceUpdateGroupLabelCall{Call: call}
}

// MockGroupLabelsServiceInterfaceUpdateGroupLabelCall wrap *gomock.Call
type MockGroupLabelsServiceInterfaceUpdateGroupLabelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupLabelsServiceInterfaceUpdateGroupLabelCall) Return(arg0 *gitlab.GroupLabel, arg1 *gitlab.Response, arg2 error) *MockGroupLabelsServiceInterfaceUpdateGroupLabelCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupLabelsServiceInterfaceUpdateGroupLabelCall) Do(f func(any, any, *gitlab.UpdateGroupLabelOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error)) *MockGroupLabelsServiceInterfaceUpdateGroupLabelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupLabelsServiceInterfaceUpdateGroupLabelCall) DoAndReturn(f func(any, any, *gitlab.UpdateGroupLabelOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error)) *MockGroupLabelsServiceInterfaceUpdateGroupLabelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
