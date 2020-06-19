// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	bytes "bytes"

	exec "k8s.io/utils/exec"

	mock "github.com/stretchr/testify/mock"
)

// ExecRunner is an autogenerated mock type for the ExecRunner type
type ExecRunner struct {
	mock.Mock
}

// RunCmd provides a mock function with given fields: cmd, cmdPath, envVars, args
func (_m *ExecRunner) RunCmd(cmd exec.Cmd, cmdPath string, envVars []string, args ...string) (*bytes.Buffer, *bytes.Buffer, error) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, cmd, cmdPath, envVars)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *bytes.Buffer
	if rf, ok := ret.Get(0).(func(exec.Cmd, string, []string, ...string) *bytes.Buffer); ok {
		r0 = rf(cmd, cmdPath, envVars, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bytes.Buffer)
		}
	}

	var r1 *bytes.Buffer
	if rf, ok := ret.Get(1).(func(exec.Cmd, string, []string, ...string) *bytes.Buffer); ok {
		r1 = rf(cmd, cmdPath, envVars, args...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*bytes.Buffer)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(exec.Cmd, string, []string, ...string) error); ok {
		r2 = rf(cmd, cmdPath, envVars, args...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}