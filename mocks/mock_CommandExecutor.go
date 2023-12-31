// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockCommandExecutor is an autogenerated mock type for the CommandExecutor type
type MockCommandExecutor struct {
	mock.Mock
}

type MockCommandExecutor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCommandExecutor) EXPECT() *MockCommandExecutor_Expecter {
	return &MockCommandExecutor_Expecter{mock: &_m.Mock}
}

// CombinedOutput provides a mock function with given fields:
func (_m *MockCommandExecutor) CombinedOutput() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCommandExecutor_CombinedOutput_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CombinedOutput'
type MockCommandExecutor_CombinedOutput_Call struct {
	*mock.Call
}

// CombinedOutput is a helper method to define mock.On call
func (_e *MockCommandExecutor_Expecter) CombinedOutput() *MockCommandExecutor_CombinedOutput_Call {
	return &MockCommandExecutor_CombinedOutput_Call{Call: _e.mock.On("CombinedOutput")}
}

func (_c *MockCommandExecutor_CombinedOutput_Call) Run(run func()) *MockCommandExecutor_CombinedOutput_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandExecutor_CombinedOutput_Call) Return(_a0 []byte, _a1 error) *MockCommandExecutor_CombinedOutput_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCommandExecutor_CombinedOutput_Call) RunAndReturn(run func() ([]byte, error)) *MockCommandExecutor_CombinedOutput_Call {
	_c.Call.Return(run)
	return _c
}

// Environ provides a mock function with given fields:
func (_m *MockCommandExecutor) Environ() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockCommandExecutor_Environ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Environ'
type MockCommandExecutor_Environ_Call struct {
	*mock.Call
}

// Environ is a helper method to define mock.On call
func (_e *MockCommandExecutor_Expecter) Environ() *MockCommandExecutor_Environ_Call {
	return &MockCommandExecutor_Environ_Call{Call: _e.mock.On("Environ")}
}

func (_c *MockCommandExecutor_Environ_Call) Run(run func()) *MockCommandExecutor_Environ_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandExecutor_Environ_Call) Return(_a0 []string) *MockCommandExecutor_Environ_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommandExecutor_Environ_Call) RunAndReturn(run func() []string) *MockCommandExecutor_Environ_Call {
	_c.Call.Return(run)
	return _c
}

// Output provides a mock function with given fields:
func (_m *MockCommandExecutor) Output() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCommandExecutor_Output_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Output'
type MockCommandExecutor_Output_Call struct {
	*mock.Call
}

// Output is a helper method to define mock.On call
func (_e *MockCommandExecutor_Expecter) Output() *MockCommandExecutor_Output_Call {
	return &MockCommandExecutor_Output_Call{Call: _e.mock.On("Output")}
}

func (_c *MockCommandExecutor_Output_Call) Run(run func()) *MockCommandExecutor_Output_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandExecutor_Output_Call) Return(_a0 []byte, _a1 error) *MockCommandExecutor_Output_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCommandExecutor_Output_Call) RunAndReturn(run func() ([]byte, error)) *MockCommandExecutor_Output_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields:
func (_m *MockCommandExecutor) Run() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCommandExecutor_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockCommandExecutor_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
func (_e *MockCommandExecutor_Expecter) Run() *MockCommandExecutor_Run_Call {
	return &MockCommandExecutor_Run_Call{Call: _e.mock.On("Run")}
}

func (_c *MockCommandExecutor_Run_Call) Run(run func()) *MockCommandExecutor_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandExecutor_Run_Call) Return(_a0 error) *MockCommandExecutor_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommandExecutor_Run_Call) RunAndReturn(run func() error) *MockCommandExecutor_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockCommandExecutor) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCommandExecutor_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockCommandExecutor_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockCommandExecutor_Expecter) Start() *MockCommandExecutor_Start_Call {
	return &MockCommandExecutor_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockCommandExecutor_Start_Call) Run(run func()) *MockCommandExecutor_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandExecutor_Start_Call) Return(_a0 error) *MockCommandExecutor_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommandExecutor_Start_Call) RunAndReturn(run func() error) *MockCommandExecutor_Start_Call {
	_c.Call.Return(run)
	return _c
}

// StderrPipe provides a mock function with given fields:
func (_m *MockCommandExecutor) StderrPipe() (io.ReadCloser, error) {
	ret := _m.Called()

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func() (io.ReadCloser, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() io.ReadCloser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCommandExecutor_StderrPipe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StderrPipe'
type MockCommandExecutor_StderrPipe_Call struct {
	*mock.Call
}

// StderrPipe is a helper method to define mock.On call
func (_e *MockCommandExecutor_Expecter) StderrPipe() *MockCommandExecutor_StderrPipe_Call {
	return &MockCommandExecutor_StderrPipe_Call{Call: _e.mock.On("StderrPipe")}
}

func (_c *MockCommandExecutor_StderrPipe_Call) Run(run func()) *MockCommandExecutor_StderrPipe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandExecutor_StderrPipe_Call) Return(_a0 io.ReadCloser, _a1 error) *MockCommandExecutor_StderrPipe_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCommandExecutor_StderrPipe_Call) RunAndReturn(run func() (io.ReadCloser, error)) *MockCommandExecutor_StderrPipe_Call {
	_c.Call.Return(run)
	return _c
}

// StdinPipe provides a mock function with given fields:
func (_m *MockCommandExecutor) StdinPipe() (io.WriteCloser, error) {
	ret := _m.Called()

	var r0 io.WriteCloser
	var r1 error
	if rf, ok := ret.Get(0).(func() (io.WriteCloser, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() io.WriteCloser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.WriteCloser)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCommandExecutor_StdinPipe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StdinPipe'
type MockCommandExecutor_StdinPipe_Call struct {
	*mock.Call
}

// StdinPipe is a helper method to define mock.On call
func (_e *MockCommandExecutor_Expecter) StdinPipe() *MockCommandExecutor_StdinPipe_Call {
	return &MockCommandExecutor_StdinPipe_Call{Call: _e.mock.On("StdinPipe")}
}

func (_c *MockCommandExecutor_StdinPipe_Call) Run(run func()) *MockCommandExecutor_StdinPipe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandExecutor_StdinPipe_Call) Return(_a0 io.WriteCloser, _a1 error) *MockCommandExecutor_StdinPipe_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCommandExecutor_StdinPipe_Call) RunAndReturn(run func() (io.WriteCloser, error)) *MockCommandExecutor_StdinPipe_Call {
	_c.Call.Return(run)
	return _c
}

// StdoutPipe provides a mock function with given fields:
func (_m *MockCommandExecutor) StdoutPipe() (io.ReadCloser, error) {
	ret := _m.Called()

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func() (io.ReadCloser, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() io.ReadCloser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCommandExecutor_StdoutPipe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StdoutPipe'
type MockCommandExecutor_StdoutPipe_Call struct {
	*mock.Call
}

// StdoutPipe is a helper method to define mock.On call
func (_e *MockCommandExecutor_Expecter) StdoutPipe() *MockCommandExecutor_StdoutPipe_Call {
	return &MockCommandExecutor_StdoutPipe_Call{Call: _e.mock.On("StdoutPipe")}
}

func (_c *MockCommandExecutor_StdoutPipe_Call) Run(run func()) *MockCommandExecutor_StdoutPipe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandExecutor_StdoutPipe_Call) Return(_a0 io.ReadCloser, _a1 error) *MockCommandExecutor_StdoutPipe_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCommandExecutor_StdoutPipe_Call) RunAndReturn(run func() (io.ReadCloser, error)) *MockCommandExecutor_StdoutPipe_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockCommandExecutor) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockCommandExecutor_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockCommandExecutor_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockCommandExecutor_Expecter) String() *MockCommandExecutor_String_Call {
	return &MockCommandExecutor_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockCommandExecutor_String_Call) Run(run func()) *MockCommandExecutor_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandExecutor_String_Call) Return(_a0 string) *MockCommandExecutor_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommandExecutor_String_Call) RunAndReturn(run func() string) *MockCommandExecutor_String_Call {
	_c.Call.Return(run)
	return _c
}

// Wait provides a mock function with given fields:
func (_m *MockCommandExecutor) Wait() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCommandExecutor_Wait_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Wait'
type MockCommandExecutor_Wait_Call struct {
	*mock.Call
}

// Wait is a helper method to define mock.On call
func (_e *MockCommandExecutor_Expecter) Wait() *MockCommandExecutor_Wait_Call {
	return &MockCommandExecutor_Wait_Call{Call: _e.mock.On("Wait")}
}

func (_c *MockCommandExecutor_Wait_Call) Run(run func()) *MockCommandExecutor_Wait_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommandExecutor_Wait_Call) Return(_a0 error) *MockCommandExecutor_Wait_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommandExecutor_Wait_Call) RunAndReturn(run func() error) *MockCommandExecutor_Wait_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCommandExecutor creates a new instance of MockCommandExecutor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCommandExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCommandExecutor {
	mock := &MockCommandExecutor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
