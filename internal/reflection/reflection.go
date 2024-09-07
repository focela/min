// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package reflection provides some reflection functions for internal usage.
package reflection

import (
	"reflect"
)

type ValueAndKindResult struct {
	InputValue  reflect.Value
	InputKind   reflect.Kind
	OriginValue reflect.Value
	OriginKind  reflect.Kind
}

// OriginValueAndKind retrieves the original non-pointer reflect value and kind.
func OriginValueAndKind(value interface{}) (out ValueAndKindResult) {
	out.InputValue = reflect.ValueOf(value)
	if v, ok := value.(reflect.Value); ok {
		out.InputValue = v
	}
	out.InputKind = out.InputValue.Kind()
	out.OriginValue, out.OriginKind = out.InputValue, out.InputKind

	for out.OriginKind == reflect.Ptr {
		out.OriginValue = out.OriginValue.Elem()
		out.OriginKind = out.OriginValue.Kind()
	}
	return
}

type TypeAndKindResult struct {
	InputType  reflect.Type
	InputKind  reflect.Kind
	OriginType reflect.Type
	OriginKind reflect.Kind
}

// OriginTypeAndKind retrieves the original non-pointer reflect type and kind.
func OriginTypeAndKind(value interface{}) (out TypeAndKindResult) {
	if value == nil {
		return
	}
	if reflectType, ok := value.(reflect.Type); ok {
		out.InputType = reflectType
	} else {
		out.InputType = reflect.TypeOf(value)
		if reflectValue, ok := value.(reflect.Value); ok {
			out.InputType = reflectValue.Type()
		}
	}
	out.InputKind = out.InputType.Kind()
	out.OriginType, out.OriginKind = out.InputType, out.InputKind

	for out.OriginKind == reflect.Ptr {
		out.OriginType = out.OriginType.Elem()
		out.OriginKind = out.OriginType.Kind()
	}
	return
}

// ValueToInterface converts reflect value to its interface type.
func ValueToInterface(v reflect.Value) (value interface{}, ok bool) {
	if v.IsValid() && v.CanInterface() {
		return v.Interface(), true
	}

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	case reflect.Complex64, reflect.Complex128:
		return v.Complex(), true
	case reflect.String:
		return v.String(), true
	case reflect.Ptr, reflect.Interface:
		return ValueToInterface(v.Elem())
	default:
		return nil, false
	}
}
