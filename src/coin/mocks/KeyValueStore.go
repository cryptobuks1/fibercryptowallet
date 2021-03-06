// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// KeyValueStore is an autogenerated mock type for the KeyValueStore type
type KeyValueStore struct {
	mock.Mock
}

// GetValue provides a mock function with given fields: key
func (_m *KeyValueStore) GetValue(key string) interface{} {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// SetValue provides a mock function with given fields: key, value
func (_m *KeyValueStore) SetValue(key string, value interface{}) {
	_m.Called(key, value)
}
