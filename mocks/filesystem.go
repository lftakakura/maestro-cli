// Automatically generated by MockGen. DO NOT EDIT!
// Source: interfaces/filesystem.go

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	afero "github.com/spf13/afero"
	os "os"
)

// Mock of FileSystem interface
type MockFileSystem struct {
	ctrl     *gomock.Controller
	recorder *_MockFileSystemRecorder
}

// Recorder for MockFileSystem (not exported)
type _MockFileSystemRecorder struct {
	mock *MockFileSystem
}

func NewMockFileSystem(ctrl *gomock.Controller) *MockFileSystem {
	mock := &MockFileSystem{ctrl: ctrl}
	mock.recorder = &_MockFileSystemRecorder{mock}
	return mock
}

func (_m *MockFileSystem) EXPECT() *_MockFileSystemRecorder {
	return _m.recorder
}

func (_m *MockFileSystem) MkdirAll(path string, perm os.FileMode) error {
	ret := _m.ctrl.Call(_m, "MkdirAll", path, perm)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFileSystemRecorder) MkdirAll(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MkdirAll", arg0, arg1)
}

func (_m *MockFileSystem) Create(name string) (afero.File, error) {
	ret := _m.ctrl.Call(_m, "Create", name)
	ret0, _ := ret[0].(afero.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFileSystemRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0)
}

func (_m *MockFileSystem) IsNotExist(err error) bool {
	ret := _m.ctrl.Call(_m, "IsNotExist", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockFileSystemRecorder) IsNotExist(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsNotExist", arg0)
}

func (_m *MockFileSystem) Stat(name string) (os.FileInfo, error) {
	ret := _m.ctrl.Call(_m, "Stat", name)
	ret0, _ := ret[0].(os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFileSystemRecorder) Stat(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stat", arg0)
}
