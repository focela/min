// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package reflection provides some reflection functions for internal usage.
package reflection

import "reflect"

// OriginValueAndKindOutput holds the input and original reflect value and kind.
type OriginValueAndKindOutput struct {
	InputValue  reflect.Value // The input reflect value.
	InputKind   reflect.Kind  // The kind of the input reflect value.
	OriginValue reflect.Value // The original reflect value after dereferencing pointers.
	OriginKind  reflect.Kind  // The kind of the original reflect value.
}

// OriginValueAndKind retrieves and returns the original reflect value and kind.
// It dereferences pointers until a non-pointer kind is reached.
func OriginValueAndKind(value interface{}) (out OriginValueAndKindOutput) {
	if v, ok := value.(reflect.Value); ok { // Check if the input is already a reflect.Value
		out.InputValue = v
	} else {
		out.InputValue = reflect.ValueOf(value) // Convert the input to reflect.Value
	}
	out.InputKind = out.InputValue.Kind()
	out.OriginValue = out.InputValue
	out.OriginKind = out.InputKind
	for out.OriginKind == reflect.Ptr { // Dereference pointers to find the origin kind and value
		out.OriginValue = out.OriginValue.Elem()
		out.OriginKind = out.OriginValue.Kind()
	}
	return
}

// OriginTypeAndKindOutput holds the input and original reflect type and kind.
type OriginTypeAndKindOutput struct {
	InputType  reflect.Type // The input reflect type.
	InputKind  reflect.Kind // The kind of the input reflect type.
	OriginType reflect.Type // The original reflect type after dereferencing pointers.
	OriginKind reflect.Kind // The kind of the original reflect type.
}

// OriginTypeAndKind retrieves and returns the original reflect type and kind.
// It dereferences pointers until a non-pointer kind is reached.
func OriginTypeAndKind(value interface{}) (out OriginTypeAndKindOutput) {
	if value == nil {
		return
	}
	if reflectType, ok := value.(reflect.Type); ok { // Check if the input is already a reflect.Type
		out.InputType = reflectType
	} else {
		if reflectValue, ok := value.(reflect.Value); ok { // Check if the input is a reflect.Value
			out.InputType = reflectValue.Type()
		} else {
			out.InputType = reflect.TypeOf(value) // Convert the input to reflect.Type
		}
	}
	out.InputKind = out.InputType.Kind()
	out.OriginType = out.InputType
	out.OriginKind = out.InputKind
	for out.OriginKind == reflect.Ptr { // Dereference pointers to find the origin kind and type
		out.OriginType = out.OriginType.Elem()
		out.OriginKind = out.OriginType.Kind()
	}
	return
}

// ValueToInterface converts a reflect.Value to its interface type.
// It returns the interface value and a boolean indicating success.
func ValueToInterface(v reflect.Value) (value interface{}, ok bool) {
	if v.IsValid() && v.CanInterface() { // Check if the value can be converted to an interface
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
	case reflect.Ptr:
		return ValueToInterface(v.Elem()) // Recursively convert pointer values
	case reflect.Interface:
		return ValueToInterface(v.Elem()) // Recursively convert interface values
	default:
		return nil, false // Return false if conversion is not possible
	}
}
