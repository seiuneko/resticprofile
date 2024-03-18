// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	config "github.com/creativeprojects/resticprofile/config"
	mock "github.com/stretchr/testify/mock"
)

// NamedPropertySet is an autogenerated mock type for the NamedPropertySet type
type NamedPropertySet struct {
	mock.Mock
}

type NamedPropertySet_Expecter struct {
	mock *mock.Mock
}

func (_m *NamedPropertySet) EXPECT() *NamedPropertySet_Expecter {
	return &NamedPropertySet_Expecter{mock: &_m.Mock}
}

// Description provides a mock function with given fields:
func (_m *NamedPropertySet) Description() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Description")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NamedPropertySet_Description_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Description'
type NamedPropertySet_Description_Call struct {
	*mock.Call
}

// Description is a helper method to define mock.On call
func (_e *NamedPropertySet_Expecter) Description() *NamedPropertySet_Description_Call {
	return &NamedPropertySet_Description_Call{Call: _e.mock.On("Description")}
}

func (_c *NamedPropertySet_Description_Call) Run(run func()) *NamedPropertySet_Description_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NamedPropertySet_Description_Call) Return(_a0 string) *NamedPropertySet_Description_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NamedPropertySet_Description_Call) RunAndReturn(run func() string) *NamedPropertySet_Description_Call {
	_c.Call.Return(run)
	return _c
}

// IsAllOptions provides a mock function with given fields:
func (_m *NamedPropertySet) IsAllOptions() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsAllOptions")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NamedPropertySet_IsAllOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAllOptions'
type NamedPropertySet_IsAllOptions_Call struct {
	*mock.Call
}

// IsAllOptions is a helper method to define mock.On call
func (_e *NamedPropertySet_Expecter) IsAllOptions() *NamedPropertySet_IsAllOptions_Call {
	return &NamedPropertySet_IsAllOptions_Call{Call: _e.mock.On("IsAllOptions")}
}

func (_c *NamedPropertySet_IsAllOptions_Call) Run(run func()) *NamedPropertySet_IsAllOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NamedPropertySet_IsAllOptions_Call) Return(_a0 bool) *NamedPropertySet_IsAllOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NamedPropertySet_IsAllOptions_Call) RunAndReturn(run func() bool) *NamedPropertySet_IsAllOptions_Call {
	_c.Call.Return(run)
	return _c
}

// IsClosed provides a mock function with given fields:
func (_m *NamedPropertySet) IsClosed() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsClosed")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NamedPropertySet_IsClosed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsClosed'
type NamedPropertySet_IsClosed_Call struct {
	*mock.Call
}

// IsClosed is a helper method to define mock.On call
func (_e *NamedPropertySet_Expecter) IsClosed() *NamedPropertySet_IsClosed_Call {
	return &NamedPropertySet_IsClosed_Call{Call: _e.mock.On("IsClosed")}
}

func (_c *NamedPropertySet_IsClosed_Call) Run(run func()) *NamedPropertySet_IsClosed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NamedPropertySet_IsClosed_Call) Return(_a0 bool) *NamedPropertySet_IsClosed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NamedPropertySet_IsClosed_Call) RunAndReturn(run func() bool) *NamedPropertySet_IsClosed_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *NamedPropertySet) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NamedPropertySet_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type NamedPropertySet_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *NamedPropertySet_Expecter) Name() *NamedPropertySet_Name_Call {
	return &NamedPropertySet_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *NamedPropertySet_Name_Call) Run(run func()) *NamedPropertySet_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NamedPropertySet_Name_Call) Return(_a0 string) *NamedPropertySet_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NamedPropertySet_Name_Call) RunAndReturn(run func() string) *NamedPropertySet_Name_Call {
	_c.Call.Return(run)
	return _c
}

// OtherPropertyInfo provides a mock function with given fields:
func (_m *NamedPropertySet) OtherPropertyInfo() config.PropertyInfo {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OtherPropertyInfo")
	}

	var r0 config.PropertyInfo
	if rf, ok := ret.Get(0).(func() config.PropertyInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.PropertyInfo)
		}
	}

	return r0
}

// NamedPropertySet_OtherPropertyInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OtherPropertyInfo'
type NamedPropertySet_OtherPropertyInfo_Call struct {
	*mock.Call
}

// OtherPropertyInfo is a helper method to define mock.On call
func (_e *NamedPropertySet_Expecter) OtherPropertyInfo() *NamedPropertySet_OtherPropertyInfo_Call {
	return &NamedPropertySet_OtherPropertyInfo_Call{Call: _e.mock.On("OtherPropertyInfo")}
}

func (_c *NamedPropertySet_OtherPropertyInfo_Call) Run(run func()) *NamedPropertySet_OtherPropertyInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NamedPropertySet_OtherPropertyInfo_Call) Return(_a0 config.PropertyInfo) *NamedPropertySet_OtherPropertyInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NamedPropertySet_OtherPropertyInfo_Call) RunAndReturn(run func() config.PropertyInfo) *NamedPropertySet_OtherPropertyInfo_Call {
	_c.Call.Return(run)
	return _c
}

// Properties provides a mock function with given fields:
func (_m *NamedPropertySet) Properties() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Properties")
	}

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

// NamedPropertySet_Properties_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Properties'
type NamedPropertySet_Properties_Call struct {
	*mock.Call
}

// Properties is a helper method to define mock.On call
func (_e *NamedPropertySet_Expecter) Properties() *NamedPropertySet_Properties_Call {
	return &NamedPropertySet_Properties_Call{Call: _e.mock.On("Properties")}
}

func (_c *NamedPropertySet_Properties_Call) Run(run func()) *NamedPropertySet_Properties_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NamedPropertySet_Properties_Call) Return(_a0 []string) *NamedPropertySet_Properties_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NamedPropertySet_Properties_Call) RunAndReturn(run func() []string) *NamedPropertySet_Properties_Call {
	_c.Call.Return(run)
	return _c
}

// PropertyInfo provides a mock function with given fields: name
func (_m *NamedPropertySet) PropertyInfo(name string) config.PropertyInfo {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for PropertyInfo")
	}

	var r0 config.PropertyInfo
	if rf, ok := ret.Get(0).(func(string) config.PropertyInfo); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.PropertyInfo)
		}
	}

	return r0
}

// NamedPropertySet_PropertyInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PropertyInfo'
type NamedPropertySet_PropertyInfo_Call struct {
	*mock.Call
}

// PropertyInfo is a helper method to define mock.On call
//   - name string
func (_e *NamedPropertySet_Expecter) PropertyInfo(name interface{}) *NamedPropertySet_PropertyInfo_Call {
	return &NamedPropertySet_PropertyInfo_Call{Call: _e.mock.On("PropertyInfo", name)}
}

func (_c *NamedPropertySet_PropertyInfo_Call) Run(run func(name string)) *NamedPropertySet_PropertyInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *NamedPropertySet_PropertyInfo_Call) Return(_a0 config.PropertyInfo) *NamedPropertySet_PropertyInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NamedPropertySet_PropertyInfo_Call) RunAndReturn(run func(string) config.PropertyInfo) *NamedPropertySet_PropertyInfo_Call {
	_c.Call.Return(run)
	return _c
}

// TypeName provides a mock function with given fields:
func (_m *NamedPropertySet) TypeName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TypeName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NamedPropertySet_TypeName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TypeName'
type NamedPropertySet_TypeName_Call struct {
	*mock.Call
}

// TypeName is a helper method to define mock.On call
func (_e *NamedPropertySet_Expecter) TypeName() *NamedPropertySet_TypeName_Call {
	return &NamedPropertySet_TypeName_Call{Call: _e.mock.On("TypeName")}
}

func (_c *NamedPropertySet_TypeName_Call) Run(run func()) *NamedPropertySet_TypeName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NamedPropertySet_TypeName_Call) Return(_a0 string) *NamedPropertySet_TypeName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NamedPropertySet_TypeName_Call) RunAndReturn(run func() string) *NamedPropertySet_TypeName_Call {
	_c.Call.Return(run)
	return _c
}

// NewNamedPropertySet creates a new instance of NamedPropertySet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNamedPropertySet(t interface {
	mock.TestingT
	Cleanup(func())
}) *NamedPropertySet {
	mock := &NamedPropertySet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}