// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	config "github.com/creativeprojects/resticprofile/config"
	mock "github.com/stretchr/testify/mock"

	restic "github.com/creativeprojects/resticprofile/restic"
)

// PropertyInfo is an autogenerated mock type for the PropertyInfo type
type PropertyInfo struct {
	mock.Mock
}

type PropertyInfo_Expecter struct {
	mock *mock.Mock
}

func (_m *PropertyInfo) EXPECT() *PropertyInfo_Expecter {
	return &PropertyInfo_Expecter{mock: &_m.Mock}
}

// CanBeBool provides a mock function with given fields:
func (_m *PropertyInfo) CanBeBool() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CanBeBool")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_CanBeBool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CanBeBool'
type PropertyInfo_CanBeBool_Call struct {
	*mock.Call
}

// CanBeBool is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) CanBeBool() *PropertyInfo_CanBeBool_Call {
	return &PropertyInfo_CanBeBool_Call{Call: _e.mock.On("CanBeBool")}
}

func (_c *PropertyInfo_CanBeBool_Call) Run(run func()) *PropertyInfo_CanBeBool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_CanBeBool_Call) Return(_a0 bool) *PropertyInfo_CanBeBool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_CanBeBool_Call) RunAndReturn(run func() bool) *PropertyInfo_CanBeBool_Call {
	_c.Call.Return(run)
	return _c
}

// CanBeNil provides a mock function with given fields:
func (_m *PropertyInfo) CanBeNil() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CanBeNil")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_CanBeNil_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CanBeNil'
type PropertyInfo_CanBeNil_Call struct {
	*mock.Call
}

// CanBeNil is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) CanBeNil() *PropertyInfo_CanBeNil_Call {
	return &PropertyInfo_CanBeNil_Call{Call: _e.mock.On("CanBeNil")}
}

func (_c *PropertyInfo_CanBeNil_Call) Run(run func()) *PropertyInfo_CanBeNil_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_CanBeNil_Call) Return(_a0 bool) *PropertyInfo_CanBeNil_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_CanBeNil_Call) RunAndReturn(run func() bool) *PropertyInfo_CanBeNil_Call {
	_c.Call.Return(run)
	return _c
}

// CanBeNumeric provides a mock function with given fields:
func (_m *PropertyInfo) CanBeNumeric() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CanBeNumeric")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_CanBeNumeric_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CanBeNumeric'
type PropertyInfo_CanBeNumeric_Call struct {
	*mock.Call
}

// CanBeNumeric is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) CanBeNumeric() *PropertyInfo_CanBeNumeric_Call {
	return &PropertyInfo_CanBeNumeric_Call{Call: _e.mock.On("CanBeNumeric")}
}

func (_c *PropertyInfo_CanBeNumeric_Call) Run(run func()) *PropertyInfo_CanBeNumeric_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_CanBeNumeric_Call) Return(_a0 bool) *PropertyInfo_CanBeNumeric_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_CanBeNumeric_Call) RunAndReturn(run func() bool) *PropertyInfo_CanBeNumeric_Call {
	_c.Call.Return(run)
	return _c
}

// CanBePropertySet provides a mock function with given fields:
func (_m *PropertyInfo) CanBePropertySet() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CanBePropertySet")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_CanBePropertySet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CanBePropertySet'
type PropertyInfo_CanBePropertySet_Call struct {
	*mock.Call
}

// CanBePropertySet is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) CanBePropertySet() *PropertyInfo_CanBePropertySet_Call {
	return &PropertyInfo_CanBePropertySet_Call{Call: _e.mock.On("CanBePropertySet")}
}

func (_c *PropertyInfo_CanBePropertySet_Call) Run(run func()) *PropertyInfo_CanBePropertySet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_CanBePropertySet_Call) Return(_a0 bool) *PropertyInfo_CanBePropertySet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_CanBePropertySet_Call) RunAndReturn(run func() bool) *PropertyInfo_CanBePropertySet_Call {
	_c.Call.Return(run)
	return _c
}

// CanBeString provides a mock function with given fields:
func (_m *PropertyInfo) CanBeString() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CanBeString")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_CanBeString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CanBeString'
type PropertyInfo_CanBeString_Call struct {
	*mock.Call
}

// CanBeString is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) CanBeString() *PropertyInfo_CanBeString_Call {
	return &PropertyInfo_CanBeString_Call{Call: _e.mock.On("CanBeString")}
}

func (_c *PropertyInfo_CanBeString_Call) Run(run func()) *PropertyInfo_CanBeString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_CanBeString_Call) Return(_a0 bool) *PropertyInfo_CanBeString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_CanBeString_Call) RunAndReturn(run func() bool) *PropertyInfo_CanBeString_Call {
	_c.Call.Return(run)
	return _c
}

// DefaultValue provides a mock function with given fields:
func (_m *PropertyInfo) DefaultValue() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DefaultValue")
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

// PropertyInfo_DefaultValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DefaultValue'
type PropertyInfo_DefaultValue_Call struct {
	*mock.Call
}

// DefaultValue is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) DefaultValue() *PropertyInfo_DefaultValue_Call {
	return &PropertyInfo_DefaultValue_Call{Call: _e.mock.On("DefaultValue")}
}

func (_c *PropertyInfo_DefaultValue_Call) Run(run func()) *PropertyInfo_DefaultValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_DefaultValue_Call) Return(_a0 []string) *PropertyInfo_DefaultValue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_DefaultValue_Call) RunAndReturn(run func() []string) *PropertyInfo_DefaultValue_Call {
	_c.Call.Return(run)
	return _c
}

// Description provides a mock function with given fields:
func (_m *PropertyInfo) Description() string {
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

// PropertyInfo_Description_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Description'
type PropertyInfo_Description_Call struct {
	*mock.Call
}

// Description is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) Description() *PropertyInfo_Description_Call {
	return &PropertyInfo_Description_Call{Call: _e.mock.On("Description")}
}

func (_c *PropertyInfo_Description_Call) Run(run func()) *PropertyInfo_Description_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_Description_Call) Return(_a0 string) *PropertyInfo_Description_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_Description_Call) RunAndReturn(run func() string) *PropertyInfo_Description_Call {
	_c.Call.Return(run)
	return _c
}

// EnumValues provides a mock function with given fields:
func (_m *PropertyInfo) EnumValues() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EnumValues")
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

// PropertyInfo_EnumValues_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnumValues'
type PropertyInfo_EnumValues_Call struct {
	*mock.Call
}

// EnumValues is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) EnumValues() *PropertyInfo_EnumValues_Call {
	return &PropertyInfo_EnumValues_Call{Call: _e.mock.On("EnumValues")}
}

func (_c *PropertyInfo_EnumValues_Call) Run(run func()) *PropertyInfo_EnumValues_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_EnumValues_Call) Return(_a0 []string) *PropertyInfo_EnumValues_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_EnumValues_Call) RunAndReturn(run func() []string) *PropertyInfo_EnumValues_Call {
	_c.Call.Return(run)
	return _c
}

// ExampleValues provides a mock function with given fields:
func (_m *PropertyInfo) ExampleValues() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ExampleValues")
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

// PropertyInfo_ExampleValues_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExampleValues'
type PropertyInfo_ExampleValues_Call struct {
	*mock.Call
}

// ExampleValues is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) ExampleValues() *PropertyInfo_ExampleValues_Call {
	return &PropertyInfo_ExampleValues_Call{Call: _e.mock.On("ExampleValues")}
}

func (_c *PropertyInfo_ExampleValues_Call) Run(run func()) *PropertyInfo_ExampleValues_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_ExampleValues_Call) Return(_a0 []string) *PropertyInfo_ExampleValues_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_ExampleValues_Call) RunAndReturn(run func() []string) *PropertyInfo_ExampleValues_Call {
	_c.Call.Return(run)
	return _c
}

// Format provides a mock function with given fields:
func (_m *PropertyInfo) Format() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Format")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PropertyInfo_Format_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Format'
type PropertyInfo_Format_Call struct {
	*mock.Call
}

// Format is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) Format() *PropertyInfo_Format_Call {
	return &PropertyInfo_Format_Call{Call: _e.mock.On("Format")}
}

func (_c *PropertyInfo_Format_Call) Run(run func()) *PropertyInfo_Format_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_Format_Call) Return(_a0 string) *PropertyInfo_Format_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_Format_Call) RunAndReturn(run func() string) *PropertyInfo_Format_Call {
	_c.Call.Return(run)
	return _c
}

// IsAnyType provides a mock function with given fields:
func (_m *PropertyInfo) IsAnyType() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsAnyType")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_IsAnyType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAnyType'
type PropertyInfo_IsAnyType_Call struct {
	*mock.Call
}

// IsAnyType is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) IsAnyType() *PropertyInfo_IsAnyType_Call {
	return &PropertyInfo_IsAnyType_Call{Call: _e.mock.On("IsAnyType")}
}

func (_c *PropertyInfo_IsAnyType_Call) Run(run func()) *PropertyInfo_IsAnyType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_IsAnyType_Call) Return(_a0 bool) *PropertyInfo_IsAnyType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_IsAnyType_Call) RunAndReturn(run func() bool) *PropertyInfo_IsAnyType_Call {
	_c.Call.Return(run)
	return _c
}

// IsDeprecated provides a mock function with given fields:
func (_m *PropertyInfo) IsDeprecated() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsDeprecated")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_IsDeprecated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsDeprecated'
type PropertyInfo_IsDeprecated_Call struct {
	*mock.Call
}

// IsDeprecated is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) IsDeprecated() *PropertyInfo_IsDeprecated_Call {
	return &PropertyInfo_IsDeprecated_Call{Call: _e.mock.On("IsDeprecated")}
}

func (_c *PropertyInfo_IsDeprecated_Call) Run(run func()) *PropertyInfo_IsDeprecated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_IsDeprecated_Call) Return(_a0 bool) *PropertyInfo_IsDeprecated_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_IsDeprecated_Call) RunAndReturn(run func() bool) *PropertyInfo_IsDeprecated_Call {
	_c.Call.Return(run)
	return _c
}

// IsMultiType provides a mock function with given fields:
func (_m *PropertyInfo) IsMultiType() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsMultiType")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_IsMultiType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsMultiType'
type PropertyInfo_IsMultiType_Call struct {
	*mock.Call
}

// IsMultiType is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) IsMultiType() *PropertyInfo_IsMultiType_Call {
	return &PropertyInfo_IsMultiType_Call{Call: _e.mock.On("IsMultiType")}
}

func (_c *PropertyInfo_IsMultiType_Call) Run(run func()) *PropertyInfo_IsMultiType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_IsMultiType_Call) Return(_a0 bool) *PropertyInfo_IsMultiType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_IsMultiType_Call) RunAndReturn(run func() bool) *PropertyInfo_IsMultiType_Call {
	_c.Call.Return(run)
	return _c
}

// IsOption provides a mock function with given fields:
func (_m *PropertyInfo) IsOption() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsOption")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_IsOption_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsOption'
type PropertyInfo_IsOption_Call struct {
	*mock.Call
}

// IsOption is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) IsOption() *PropertyInfo_IsOption_Call {
	return &PropertyInfo_IsOption_Call{Call: _e.mock.On("IsOption")}
}

func (_c *PropertyInfo_IsOption_Call) Run(run func()) *PropertyInfo_IsOption_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_IsOption_Call) Return(_a0 bool) *PropertyInfo_IsOption_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_IsOption_Call) RunAndReturn(run func() bool) *PropertyInfo_IsOption_Call {
	_c.Call.Return(run)
	return _c
}

// IsRequired provides a mock function with given fields:
func (_m *PropertyInfo) IsRequired() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsRequired")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_IsRequired_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRequired'
type PropertyInfo_IsRequired_Call struct {
	*mock.Call
}

// IsRequired is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) IsRequired() *PropertyInfo_IsRequired_Call {
	return &PropertyInfo_IsRequired_Call{Call: _e.mock.On("IsRequired")}
}

func (_c *PropertyInfo_IsRequired_Call) Run(run func()) *PropertyInfo_IsRequired_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_IsRequired_Call) Return(_a0 bool) *PropertyInfo_IsRequired_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_IsRequired_Call) RunAndReturn(run func() bool) *PropertyInfo_IsRequired_Call {
	_c.Call.Return(run)
	return _c
}

// IsSingle provides a mock function with given fields:
func (_m *PropertyInfo) IsSingle() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSingle")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_IsSingle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSingle'
type PropertyInfo_IsSingle_Call struct {
	*mock.Call
}

// IsSingle is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) IsSingle() *PropertyInfo_IsSingle_Call {
	return &PropertyInfo_IsSingle_Call{Call: _e.mock.On("IsSingle")}
}

func (_c *PropertyInfo_IsSingle_Call) Run(run func()) *PropertyInfo_IsSingle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_IsSingle_Call) Return(_a0 bool) *PropertyInfo_IsSingle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_IsSingle_Call) RunAndReturn(run func() bool) *PropertyInfo_IsSingle_Call {
	_c.Call.Return(run)
	return _c
}

// IsSinglePropertySet provides a mock function with given fields:
func (_m *PropertyInfo) IsSinglePropertySet() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSinglePropertySet")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_IsSinglePropertySet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSinglePropertySet'
type PropertyInfo_IsSinglePropertySet_Call struct {
	*mock.Call
}

// IsSinglePropertySet is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) IsSinglePropertySet() *PropertyInfo_IsSinglePropertySet_Call {
	return &PropertyInfo_IsSinglePropertySet_Call{Call: _e.mock.On("IsSinglePropertySet")}
}

func (_c *PropertyInfo_IsSinglePropertySet_Call) Run(run func()) *PropertyInfo_IsSinglePropertySet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_IsSinglePropertySet_Call) Return(_a0 bool) *PropertyInfo_IsSinglePropertySet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_IsSinglePropertySet_Call) RunAndReturn(run func() bool) *PropertyInfo_IsSinglePropertySet_Call {
	_c.Call.Return(run)
	return _c
}

// MustBeInteger provides a mock function with given fields:
func (_m *PropertyInfo) MustBeInteger() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MustBeInteger")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PropertyInfo_MustBeInteger_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MustBeInteger'
type PropertyInfo_MustBeInteger_Call struct {
	*mock.Call
}

// MustBeInteger is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) MustBeInteger() *PropertyInfo_MustBeInteger_Call {
	return &PropertyInfo_MustBeInteger_Call{Call: _e.mock.On("MustBeInteger")}
}

func (_c *PropertyInfo_MustBeInteger_Call) Run(run func()) *PropertyInfo_MustBeInteger_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_MustBeInteger_Call) Return(_a0 bool) *PropertyInfo_MustBeInteger_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_MustBeInteger_Call) RunAndReturn(run func() bool) *PropertyInfo_MustBeInteger_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *PropertyInfo) Name() string {
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

// PropertyInfo_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type PropertyInfo_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) Name() *PropertyInfo_Name_Call {
	return &PropertyInfo_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *PropertyInfo_Name_Call) Run(run func()) *PropertyInfo_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_Name_Call) Return(_a0 string) *PropertyInfo_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_Name_Call) RunAndReturn(run func() string) *PropertyInfo_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NumericRange provides a mock function with given fields:
func (_m *PropertyInfo) NumericRange() config.NumericRange {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NumericRange")
	}

	var r0 config.NumericRange
	if rf, ok := ret.Get(0).(func() config.NumericRange); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.NumericRange)
	}

	return r0
}

// PropertyInfo_NumericRange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NumericRange'
type PropertyInfo_NumericRange_Call struct {
	*mock.Call
}

// NumericRange is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) NumericRange() *PropertyInfo_NumericRange_Call {
	return &PropertyInfo_NumericRange_Call{Call: _e.mock.On("NumericRange")}
}

func (_c *PropertyInfo_NumericRange_Call) Run(run func()) *PropertyInfo_NumericRange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_NumericRange_Call) Return(_a0 config.NumericRange) *PropertyInfo_NumericRange_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_NumericRange_Call) RunAndReturn(run func() config.NumericRange) *PropertyInfo_NumericRange_Call {
	_c.Call.Return(run)
	return _c
}

// Option provides a mock function with given fields:
func (_m *PropertyInfo) Option() restic.Option {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Option")
	}

	var r0 restic.Option
	if rf, ok := ret.Get(0).(func() restic.Option); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(restic.Option)
	}

	return r0
}

// PropertyInfo_Option_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Option'
type PropertyInfo_Option_Call struct {
	*mock.Call
}

// Option is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) Option() *PropertyInfo_Option_Call {
	return &PropertyInfo_Option_Call{Call: _e.mock.On("Option")}
}

func (_c *PropertyInfo_Option_Call) Run(run func()) *PropertyInfo_Option_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_Option_Call) Return(_a0 restic.Option) *PropertyInfo_Option_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_Option_Call) RunAndReturn(run func() restic.Option) *PropertyInfo_Option_Call {
	_c.Call.Return(run)
	return _c
}

// PropertySet provides a mock function with given fields:
func (_m *PropertyInfo) PropertySet() config.NamedPropertySet {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PropertySet")
	}

	var r0 config.NamedPropertySet
	if rf, ok := ret.Get(0).(func() config.NamedPropertySet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.NamedPropertySet)
		}
	}

	return r0
}

// PropertyInfo_PropertySet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PropertySet'
type PropertyInfo_PropertySet_Call struct {
	*mock.Call
}

// PropertySet is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) PropertySet() *PropertyInfo_PropertySet_Call {
	return &PropertyInfo_PropertySet_Call{Call: _e.mock.On("PropertySet")}
}

func (_c *PropertyInfo_PropertySet_Call) Run(run func()) *PropertyInfo_PropertySet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_PropertySet_Call) Return(_a0 config.NamedPropertySet) *PropertyInfo_PropertySet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_PropertySet_Call) RunAndReturn(run func() config.NamedPropertySet) *PropertyInfo_PropertySet_Call {
	_c.Call.Return(run)
	return _c
}

// ValidationPattern provides a mock function with given fields:
func (_m *PropertyInfo) ValidationPattern() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ValidationPattern")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PropertyInfo_ValidationPattern_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidationPattern'
type PropertyInfo_ValidationPattern_Call struct {
	*mock.Call
}

// ValidationPattern is a helper method to define mock.On call
func (_e *PropertyInfo_Expecter) ValidationPattern() *PropertyInfo_ValidationPattern_Call {
	return &PropertyInfo_ValidationPattern_Call{Call: _e.mock.On("ValidationPattern")}
}

func (_c *PropertyInfo_ValidationPattern_Call) Run(run func()) *PropertyInfo_ValidationPattern_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PropertyInfo_ValidationPattern_Call) Return(_a0 string) *PropertyInfo_ValidationPattern_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PropertyInfo_ValidationPattern_Call) RunAndReturn(run func() string) *PropertyInfo_ValidationPattern_Call {
	_c.Call.Return(run)
	return _c
}

// NewPropertyInfo creates a new instance of PropertyInfo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPropertyInfo(t interface {
	mock.TestingT
	Cleanup(func())
}) *PropertyInfo {
	mock := &PropertyInfo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}