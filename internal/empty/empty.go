// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package empty provides functions for checking empty or nil variables.
package empty

import (
	"reflect"
	"time"

	"github.com/focela/min/internal/reflection"
)

// Stringer is used for type assertion to a String() method.
type Stringer interface {
	String() string
}

// Interfacer is used for type assertion to an Interfaces() method.
type Interfacer interface {
	Interfaces() []interface{}
}

// MapStrAny is the interface for converting struct parameters to a map.
type MapStrAny interface {
	MapStrAny() map[string]interface{}
}

// Timer is the interface for working with time-related operations.
type Timer interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// IsEmpty checks whether the given `value` is empty.
// It returns true if `value` is 0, nil, false, "", or has a len(slice/map/chan) == 0.
// It uses reflection for more complex types, which may impact performance slightly.
//
// The parameter `traceSource` is used to trace back to the source variable if the given `value`
// is a pointer that also points to another pointer. It returns true if the source is empty when `traceSource` is true.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	// First check common types using assertion for performance, then use reflection.
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		return v == 0
	case uint, uint8, uint16, uint32, uint64:
		return v == 0
	case float32, float64:
		return v == 0
	case bool:
		return !v
	case string:
		return v == ""
	case []byte, []rune, []int, []string, []float32, []float64:
		return reflect.ValueOf(value).Len() == 0
	case map[string]interface{}:
		return len(v) == 0
	}

	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		// Handle common interfaces
		if f, ok := value.(Timer); ok {
			if f == (*time.Time)(nil) {
				return true
			}
			return f.IsZero()
		}
		if f, ok := value.(Stringer); ok {
			return f == nil || f.String() == ""
		}
		if f, ok := value.(Interfacer); ok {
			return f == nil || len(f.Interfaces()) == 0
		}
		if f, ok := value.(MapStrAny); ok {
			return f == nil || len(f.MapStrAny()) == 0
		}

		rv = reflect.ValueOf(value)
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
	case reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return rv.IsNil()
	case reflect.Invalid:
		return true
	}
	return false
}

// IsNil checks whether the given `value` is nil, especially for interface{} type values.
// The `traceSource` parameter is used to trace back to the source variable if `value` is a pointer
// that also points to another pointer. It returns true if the source is nil when `traceSource` is true.
func IsNil(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	rv := reflect.ValueOf(value)
	if v, ok := value.(reflect.Value); ok {
		rv = v
	}

	switch rv.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return !rv.IsValid() || rv.IsNil()
	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] {
			for rv.Kind() == reflect.Ptr {
				rv = rv.Elem()
			}
			if !rv.IsValid() {
				return true
			}
			if rv.Kind() == reflect.Ptr {
				return rv.IsNil()
			}
		} else {
			return !rv.IsValid() || rv.IsNil()
		}
	}
	return false
}
