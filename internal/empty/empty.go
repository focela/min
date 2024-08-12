// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package empty provides functions for checking empty/nil variables.
package empty

import (
	"reflect"
	"time"

	"github.com/focela/aid/internal/reflection"
)

// Stringer is used for type assertion on the String() method.
type Stringer interface {
	String() string
}

// InterfaceProvider is used for type assertion on the Interfaces() method.
type InterfaceProvider interface {
	Interfaces() []interface{}
}

// MapStringAny is the interface support for converting struct parameter to a map.
type MapStringAny interface {
	MapStrAny() map[string]interface{}
}

type TimeProvider interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// IsEmpty checks whether the given `value` is empty.
// It returns true if `value` is 0, nil, false, "", len(slice/map/chan) == 0,
// or else it returns false.
//
// The `traceSource` parameter is used to trace to the source variable if the given `value`
// is a pointer that may point to another pointer. It returns true if the source is empty
// when `traceSource` is true. Note that it might use reflection, which can affect performance.
// IsEmpty checks whether the given `value` is empty.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	// First check common types to enhance performance, then use reflection.
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
	case map[string]interface{}:
		return len(result) == 0

	default:
		// Handle slices and other types using reflection.
		rv := reflect.ValueOf(result)
		switch rv.Kind() {
		case reflect.Slice, reflect.Array:
			return rv.Len() == 0
		}
	}

	// Check common interfaces
	switch f := value.(type) {
	case TimeProvider:
		if f == (*time.Time)(nil) {
			return true
		}
		return f.IsZero()
	case Stringer:
		if f == nil {
			return true
		}
		return f.String() == ""
	case InterfaceProvider:
		if f == nil {
			return true
		}
		return len(f.Interfaces()) == 0
	case MapStringAny:
		if f == nil {
			return true
		}
		return len(f.MapStrAny()) == 0
	}

	// Finally, use reflection.
	return isEmptyReflect(reflect.ValueOf(value), traceSource...)
}

func isEmptyReflect(rv reflect.Value, traceSource ...bool) bool {
	if !rv.IsValid() {
		return true
	}

	switch rv.Kind() {
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.String:
		return rv.Len() == 0
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			fieldValueInterface, _ := reflection.ValueToInterface(rv.Field(i))
			if !IsEmpty(fieldValueInterface) {
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
	case reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return rv.IsNil()
	case reflect.Invalid:
		return true
	}
	return false
}

// IsNil checks whether the given `value` is nil, especially for interface{} type values.
// The `traceSource` parameter is used to trace to the source variable if the given `value`
// is a pointer that may point to another pointer. It returns true if the source is nil
// when `traceSource` is true. Note that it might use reflection, which can affect performance.
func IsNil(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	return isNilReflect(rv, traceSource...)
}

func isNilReflect(rv reflect.Value, traceSource ...bool) bool {
	if !rv.IsValid() {
		return true
	}

	switch rv.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return rv.IsNil()
	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] {
			for rv.Kind() == reflect.Ptr {
				rv = rv.Elem()
			}
			if !rv.IsValid() || rv.Kind() == reflect.Ptr && rv.IsNil() {
				return true
			}
		} else {
			return rv.IsNil()
		}
	}
	return false
}
