// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Encoder is an autogenerated mock type for the Encoder type
type Encoder struct {
	mock.Mock
}

// Decode provides a mock function with given fields: _a0, _a1
func (_m *Encoder) Decode(_a0 []byte, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Encode provides a mock function with given fields: value
func (_m *Encoder) Encode(value interface{}) ([]byte, error) {
	ret := _m.Called(value)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(interface{}) []byte); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewEncoderT interface {
	mock.TestingT
	Cleanup(func())
}

// NewEncoder creates a new instance of Encoder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEncoder(t NewEncoderT) *Encoder {
	mock := &Encoder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
