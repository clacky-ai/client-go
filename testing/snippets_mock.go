// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: SnippetsServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=snippets_mock.go -package=testing gitlab.com/gitlab-org/api/client-go SnippetsServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockSnippetsServiceInterface is a mock of SnippetsServiceInterface interface.
type MockSnippetsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSnippetsServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockSnippetsServiceInterfaceMockRecorder is the mock recorder for MockSnippetsServiceInterface.
type MockSnippetsServiceInterfaceMockRecorder struct {
	mock *MockSnippetsServiceInterface
}

// NewMockSnippetsServiceInterface creates a new mock instance.
func NewMockSnippetsServiceInterface(ctrl *gomock.Controller) *MockSnippetsServiceInterface {
	mock := &MockSnippetsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockSnippetsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnippetsServiceInterface) EXPECT() *MockSnippetsServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateSnippet mocks base method.
func (m *MockSnippetsServiceInterface) CreateSnippet(opt *gitlab.CreateSnippetOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSnippet indicates an expected call of CreateSnippet.
func (mr *MockSnippetsServiceInterfaceMockRecorder) CreateSnippet(opt any, options ...any) *MockSnippetsServiceInterfaceCreateSnippetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnippet", reflect.TypeOf((*MockSnippetsServiceInterface)(nil).CreateSnippet), varargs...)
	return &MockSnippetsServiceInterfaceCreateSnippetCall{Call: call}
}

// MockSnippetsServiceInterfaceCreateSnippetCall wrap *gomock.Call
type MockSnippetsServiceInterfaceCreateSnippetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSnippetsServiceInterfaceCreateSnippetCall) Return(arg0 *gitlab.Snippet, arg1 *gitlab.Response, arg2 error) *MockSnippetsServiceInterfaceCreateSnippetCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSnippetsServiceInterfaceCreateSnippetCall) Do(f func(*gitlab.CreateSnippetOptions, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *MockSnippetsServiceInterfaceCreateSnippetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSnippetsServiceInterfaceCreateSnippetCall) DoAndReturn(f func(*gitlab.CreateSnippetOptions, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *MockSnippetsServiceInterfaceCreateSnippetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteSnippet mocks base method.
func (m *MockSnippetsServiceInterface) DeleteSnippet(snippet int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSnippet indicates an expected call of DeleteSnippet.
func (mr *MockSnippetsServiceInterfaceMockRecorder) DeleteSnippet(snippet any, options ...any) *MockSnippetsServiceInterfaceDeleteSnippetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnippet", reflect.TypeOf((*MockSnippetsServiceInterface)(nil).DeleteSnippet), varargs...)
	return &MockSnippetsServiceInterfaceDeleteSnippetCall{Call: call}
}

// MockSnippetsServiceInterfaceDeleteSnippetCall wrap *gomock.Call
type MockSnippetsServiceInterfaceDeleteSnippetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSnippetsServiceInterfaceDeleteSnippetCall) Return(arg0 *gitlab.Response, arg1 error) *MockSnippetsServiceInterfaceDeleteSnippetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSnippetsServiceInterfaceDeleteSnippetCall) Do(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockSnippetsServiceInterfaceDeleteSnippetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSnippetsServiceInterfaceDeleteSnippetCall) DoAndReturn(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockSnippetsServiceInterfaceDeleteSnippetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ExploreSnippets mocks base method.
func (m *MockSnippetsServiceInterface) ExploreSnippets(opt *gitlab.ExploreSnippetsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExploreSnippets", varargs...)
	ret0, _ := ret[0].([]*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExploreSnippets indicates an expected call of ExploreSnippets.
func (mr *MockSnippetsServiceInterfaceMockRecorder) ExploreSnippets(opt any, options ...any) *MockSnippetsServiceInterfaceExploreSnippetsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExploreSnippets", reflect.TypeOf((*MockSnippetsServiceInterface)(nil).ExploreSnippets), varargs...)
	return &MockSnippetsServiceInterfaceExploreSnippetsCall{Call: call}
}

// MockSnippetsServiceInterfaceExploreSnippetsCall wrap *gomock.Call
type MockSnippetsServiceInterfaceExploreSnippetsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSnippetsServiceInterfaceExploreSnippetsCall) Return(arg0 []*gitlab.Snippet, arg1 *gitlab.Response, arg2 error) *MockSnippetsServiceInterfaceExploreSnippetsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSnippetsServiceInterfaceExploreSnippetsCall) Do(f func(*gitlab.ExploreSnippetsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error)) *MockSnippetsServiceInterfaceExploreSnippetsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSnippetsServiceInterfaceExploreSnippetsCall) DoAndReturn(f func(*gitlab.ExploreSnippetsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error)) *MockSnippetsServiceInterfaceExploreSnippetsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSnippet mocks base method.
func (m *MockSnippetsServiceInterface) GetSnippet(snippet int, options ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSnippet indicates an expected call of GetSnippet.
func (mr *MockSnippetsServiceInterfaceMockRecorder) GetSnippet(snippet any, options ...any) *MockSnippetsServiceInterfaceGetSnippetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnippet", reflect.TypeOf((*MockSnippetsServiceInterface)(nil).GetSnippet), varargs...)
	return &MockSnippetsServiceInterfaceGetSnippetCall{Call: call}
}

// MockSnippetsServiceInterfaceGetSnippetCall wrap *gomock.Call
type MockSnippetsServiceInterfaceGetSnippetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSnippetsServiceInterfaceGetSnippetCall) Return(arg0 *gitlab.Snippet, arg1 *gitlab.Response, arg2 error) *MockSnippetsServiceInterfaceGetSnippetCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSnippetsServiceInterfaceGetSnippetCall) Do(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *MockSnippetsServiceInterfaceGetSnippetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSnippetsServiceInterfaceGetSnippetCall) DoAndReturn(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *MockSnippetsServiceInterfaceGetSnippetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListAllSnippets mocks base method.
func (m *MockSnippetsServiceInterface) ListAllSnippets(opt *gitlab.ListAllSnippetsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAllSnippets", varargs...)
	ret0, _ := ret[0].([]*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAllSnippets indicates an expected call of ListAllSnippets.
func (mr *MockSnippetsServiceInterfaceMockRecorder) ListAllSnippets(opt any, options ...any) *MockSnippetsServiceInterfaceListAllSnippetsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllSnippets", reflect.TypeOf((*MockSnippetsServiceInterface)(nil).ListAllSnippets), varargs...)
	return &MockSnippetsServiceInterfaceListAllSnippetsCall{Call: call}
}

// MockSnippetsServiceInterfaceListAllSnippetsCall wrap *gomock.Call
type MockSnippetsServiceInterfaceListAllSnippetsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSnippetsServiceInterfaceListAllSnippetsCall) Return(arg0 []*gitlab.Snippet, arg1 *gitlab.Response, arg2 error) *MockSnippetsServiceInterfaceListAllSnippetsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSnippetsServiceInterfaceListAllSnippetsCall) Do(f func(*gitlab.ListAllSnippetsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error)) *MockSnippetsServiceInterfaceListAllSnippetsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSnippetsServiceInterfaceListAllSnippetsCall) DoAndReturn(f func(*gitlab.ListAllSnippetsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error)) *MockSnippetsServiceInterfaceListAllSnippetsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListSnippets mocks base method.
func (m *MockSnippetsServiceInterface) ListSnippets(opt *gitlab.ListSnippetsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSnippets", varargs...)
	ret0, _ := ret[0].([]*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSnippets indicates an expected call of ListSnippets.
func (mr *MockSnippetsServiceInterfaceMockRecorder) ListSnippets(opt any, options ...any) *MockSnippetsServiceInterfaceListSnippetsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnippets", reflect.TypeOf((*MockSnippetsServiceInterface)(nil).ListSnippets), varargs...)
	return &MockSnippetsServiceInterfaceListSnippetsCall{Call: call}
}

// MockSnippetsServiceInterfaceListSnippetsCall wrap *gomock.Call
type MockSnippetsServiceInterfaceListSnippetsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSnippetsServiceInterfaceListSnippetsCall) Return(arg0 []*gitlab.Snippet, arg1 *gitlab.Response, arg2 error) *MockSnippetsServiceInterfaceListSnippetsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSnippetsServiceInterfaceListSnippetsCall) Do(f func(*gitlab.ListSnippetsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error)) *MockSnippetsServiceInterfaceListSnippetsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSnippetsServiceInterfaceListSnippetsCall) DoAndReturn(f func(*gitlab.ListSnippetsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error)) *MockSnippetsServiceInterfaceListSnippetsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SnippetContent mocks base method.
func (m *MockSnippetsServiceInterface) SnippetContent(snippet int, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SnippetContent", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SnippetContent indicates an expected call of SnippetContent.
func (mr *MockSnippetsServiceInterfaceMockRecorder) SnippetContent(snippet any, options ...any) *MockSnippetsServiceInterfaceSnippetContentCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnippetContent", reflect.TypeOf((*MockSnippetsServiceInterface)(nil).SnippetContent), varargs...)
	return &MockSnippetsServiceInterfaceSnippetContentCall{Call: call}
}

// MockSnippetsServiceInterfaceSnippetContentCall wrap *gomock.Call
type MockSnippetsServiceInterfaceSnippetContentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSnippetsServiceInterfaceSnippetContentCall) Return(arg0 []byte, arg1 *gitlab.Response, arg2 error) *MockSnippetsServiceInterfaceSnippetContentCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSnippetsServiceInterfaceSnippetContentCall) Do(f func(int, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)) *MockSnippetsServiceInterfaceSnippetContentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSnippetsServiceInterfaceSnippetContentCall) DoAndReturn(f func(int, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)) *MockSnippetsServiceInterfaceSnippetContentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SnippetFileContent mocks base method.
func (m *MockSnippetsServiceInterface) SnippetFileContent(snippet int, ref, filename string, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet, ref, filename}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SnippetFileContent", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SnippetFileContent indicates an expected call of SnippetFileContent.
func (mr *MockSnippetsServiceInterfaceMockRecorder) SnippetFileContent(snippet, ref, filename any, options ...any) *MockSnippetsServiceInterfaceSnippetFileContentCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet, ref, filename}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnippetFileContent", reflect.TypeOf((*MockSnippetsServiceInterface)(nil).SnippetFileContent), varargs...)
	return &MockSnippetsServiceInterfaceSnippetFileContentCall{Call: call}
}

// MockSnippetsServiceInterfaceSnippetFileContentCall wrap *gomock.Call
type MockSnippetsServiceInterfaceSnippetFileContentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSnippetsServiceInterfaceSnippetFileContentCall) Return(arg0 []byte, arg1 *gitlab.Response, arg2 error) *MockSnippetsServiceInterfaceSnippetFileContentCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSnippetsServiceInterfaceSnippetFileContentCall) Do(f func(int, string, string, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)) *MockSnippetsServiceInterfaceSnippetFileContentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSnippetsServiceInterfaceSnippetFileContentCall) DoAndReturn(f func(int, string, string, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)) *MockSnippetsServiceInterfaceSnippetFileContentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateSnippet mocks base method.
func (m *MockSnippetsServiceInterface) UpdateSnippet(snippet int, opt *gitlab.UpdateSnippetOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateSnippet indicates an expected call of UpdateSnippet.
func (mr *MockSnippetsServiceInterfaceMockRecorder) UpdateSnippet(snippet, opt any, options ...any) *MockSnippetsServiceInterfaceUpdateSnippetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSnippet", reflect.TypeOf((*MockSnippetsServiceInterface)(nil).UpdateSnippet), varargs...)
	return &MockSnippetsServiceInterfaceUpdateSnippetCall{Call: call}
}

// MockSnippetsServiceInterfaceUpdateSnippetCall wrap *gomock.Call
type MockSnippetsServiceInterfaceUpdateSnippetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSnippetsServiceInterfaceUpdateSnippetCall) Return(arg0 *gitlab.Snippet, arg1 *gitlab.Response, arg2 error) *MockSnippetsServiceInterfaceUpdateSnippetCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSnippetsServiceInterfaceUpdateSnippetCall) Do(f func(int, *gitlab.UpdateSnippetOptions, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *MockSnippetsServiceInterfaceUpdateSnippetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSnippetsServiceInterfaceUpdateSnippetCall) DoAndReturn(f func(int, *gitlab.UpdateSnippetOptions, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *MockSnippetsServiceInterfaceUpdateSnippetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
