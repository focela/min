// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package empty provides functions for checking empty/nil variables.
package empty

import (
	"reflect"
	"time"

	"github.com/focela/plume/internal/reflection"
)

// Stringer is used for type assertion API for String().
type Stringer interface {
	String() string
}

// InterfaceProvider is used for type assertion API for Interfaces.
type InterfaceProvider interface {
	Interfaces() []interface{}
}

// MapStringInterfaceProvider is the interface supporting conversion of a struct parameter to a map.
type MapStringInterfaceProvider interface {
	MapStrAny() map[string]interface{}
}

// TimeProvider provides time-related functionalities for objects.
type TimeProvider interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// IsEmpty checks whether a given `value` is empty.
// It returns true if `value` is in: 0, nil, false, "", len(slice/map/chan) == 0,
// or else it returns false.
//
// The parameter `traceSource` is used for tracing to the source variable if the given `value` is a pointer
// that points to another pointer. It returns true if the source is empty when `traceSource` is true.
// Note that it might use reflection, which affects performance slightly.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}
	// It first checks the variable as common types using assertion to enhance performance,
	// and then uses reflection.
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
	case []byte, []rune, []int, []string, []float32, []float64:
		return len(result) == 0
	case map[string]interface{}:
		return len(result) == 0

	default:
		// Finally, use reflection.
		var rv reflect.Value
		if v, ok := value.(reflect.Value); ok {
			rv = v
		} else {
			// =========================
			// Common interface checks.
			// =========================
			if f, ok := value.(TimeProvider); ok {
				if f == (*time.Time)(nil) {
					return true
				}
				return f.IsZero()
			}
			if f, ok := value.(Stringer); ok {
				if f == nil {
					return true
				}
				return f.String() == ""
			}
			if f, ok := value.(InterfaceProvider); ok {
				if f == nil {
					return true
				}
				return len(f.Interfaces()) == 0
			}
			if f, ok := value.(MapStringInterfaceProvider); ok {
				if f == nil {
					return true
				}
				return len(f.MapStrAny()) == 0
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
			var fieldValueInterface interface{}
			for i := 0; i < rv.NumField(); i++ {
				fieldValueInterface, _ = reflection.ValueToInterface(rv.Field(i))
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
	}
	return false
}

// IsNil checks whether the given `value` is nil, especially for interface{} type values.
// The parameter `traceSource` is used for tracing to the source variable if the given `value` is a pointer
// that points to another pointer. It returns nil if the source is nil when `traceSource` is true.
// Note that it might use reflection, which affects performance slightly.
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
