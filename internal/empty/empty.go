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

type Stringer interface {
	String() string
}

type Interfacer interface {
	Interfaces() []interface{}
}

type MapStrAnyer interface {
	MapStrAny() map[string]interface{}
}

type Timer interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// IsEmpty checks whether the given `value` is empty.
// It returns true if `value` is 0, nil, false, "", len(slice/map/chan) == 0,
// otherwise it returns false.
//
// The parameter `traceSource` is used to trace to the source variable if `value` is a pointer
// that points to another pointer. It returns true if the source is empty when `traceSource` is true.
// Note that this function might use reflection, which could slightly impact performance.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}

	// Check common types using type assertions to enhance performance.
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
	case []byte:
		return len(result) == 0
	case []rune:
		return len(result) == 0
	case []int:
		return len(result) == 0
	case []string:
		return len(result) == 0
	case []float32:
		return len(result) == 0
	case []float64:
		return len(result) == 0
	case map[string]interface{}:
		return len(result) == 0
	}

	// Check using reflection as the last resort.
	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		// Check common interfaces.
		switch f := value.(type) {
		case Timer:
			return f == (*time.Time)(nil) || f.IsZero()
		case Stringer:
			return f == nil || f.String() == ""
		case Interfacer:
			return f == nil || len(f.Interfaces()) == 0
		case MapStrAnyer:
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
// Parameter `traceSource` is used to trace to the source variable if `value` is a pointer
// that points to another pointer. It returns nil if the source is nil when `traceSource` is true.
// Note that this function might use reflection, which could slightly impact performance.
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
			return rv.Kind() == reflect.Ptr && rv.IsNil()
		}
		return !rv.IsValid() || rv.IsNil()
	}
	return false
}
