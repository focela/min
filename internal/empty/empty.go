// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package empty provides functions for checking empty/nil variables.
package empty

import (
	"reflect"
	"time"

	"github.com/focela/min/internal/reflection"
)

// Stringer is used for type assert API for String().
type Stringer interface {
	String() string
}

// InterfaceProvider is used for type assert API for providing interfaces.
type InterfaceProvider interface {
	Interfaces() []interface{}
}

// Mapper is the interface for converting struct parameters to a map.
type Mapper interface {
	MapStrAny() map[string]interface{}
}

type Timer interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// IsEmpty checks if the given value is empty. It returns true for values like
// 0, nil, false, "", and empty slices/maps. When traceSource is true, it traces
// pointers to their original values and checks them.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	switch result := value.(type) {
	case int, int8, int16, int32, int64:
		return result == 0
	case uint, uint8, uint16, uint32, uint64:
		return result == 0
	case float32, float64:
		return result == 0
	case bool:
		return !result
	case string:
		return result == ""
	default:
		rv := reflect.ValueOf(result)
		switch rv.Kind() {
		case reflect.Slice, reflect.Map, reflect.Array, reflect.Chan, reflect.String:
			return rv.Len() == 0
		default:
			return isEmptyUsingReflection(value, traceSource...)
		}
	}
}

// isEmptyUsingReflection checks if a given value is empty using reflection.
func isEmptyUsingReflection(value interface{}, traceSource ...bool) bool {
	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		if isCommonInterfaceEmpty(value) {
			return true
		}
		rv = reflect.ValueOf(value)
	}

	switch rv.Kind() {
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.String:
		return rv.Len() == 0
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			fieldValue, _ := reflection.ValueToInterface(rv.Field(i))
			if !IsEmpty(fieldValue) {
				return false
			}
		}
		return true
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return rv.Len() == 0
	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] {
			return IsEmpty(rv.Elem())
		}
		return rv.IsNil()
	default:
		return !rv.IsValid()
	}
}

// isCommonInterfaceEmpty checks if the common interfaces are empty.
func isCommonInterfaceEmpty(value interface{}) bool {
	switch f := value.(type) {
	case Timer:
		return f == (*time.Time)(nil) || f.IsZero()
	case Stringer:
		return f == nil || f.String() == ""
	case InterfaceProvider:
		return f == nil || len(f.Interfaces()) == 0
	case Mapper:
		return f == nil || len(f.MapStrAny()) == 0
	}
	return false
}

// IsNil checks if the given value is nil, supporting interface{} type.
// When traceSource is true, it traces through pointers and checks the original value.
func IsNil(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}
	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(value)
	}

	switch rv.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return !rv.IsValid() || rv.IsNil()
	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] {
			for rv.Kind() == reflect.Ptr {
				rv = rv.Elem()
			}
			return !rv.IsValid() || (rv.Kind() == reflect.Ptr && rv.IsNil())
		}
		return !rv.IsValid() || rv.IsNil()
	}
	return false
}
