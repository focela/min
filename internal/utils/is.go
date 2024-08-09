// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"reflect"

	"github.com/focela/aid/internal/empty"
)

// IsNil checks whether `value` is nil, especially for interface{} type value.
func IsNil(value interface{}) bool {
	return empty.IsNil(value)
}

// IsEmpty checks whether `value` is empty.
func IsEmpty(value interface{}) bool {
	return empty.IsEmpty(value)
}

// IsInt checks whether `value` is of type int.
func IsInt(value interface{}) bool {
	switch value.(type) {
	case int, *int, int8, *int8, int16, *int16, int32, *int32, int64, *int64:
		return true
	}
	return false
}

// IsUint checks whether `value` is of type uint.
func IsUint(value interface{}) bool {
	switch value.(type) {
	case uint, *uint, uint8, *uint8, uint16, *uint16, uint32, *uint32, uint64, *uint64:
		return true
	}
	return false
}

// IsFloat checks whether `value` is of type float.
func IsFloat(value interface{}) bool {
	switch value.(type) {
	case float32, *float32, float64, *float64:
		return true
	}
	return false
}

// IsSlice checks whether `value` is of type slice.
func IsSlice(value interface{}) bool {
	reflectValue := reflect.Indirect(reflect.ValueOf(value))
	return reflectValue.Kind() == reflect.Slice || reflectValue.Kind() == reflect.Array
}

// IsMap checks whether `value` is of type map.
func IsMap(value interface{}) bool {
	reflectValue := reflect.Indirect(reflect.ValueOf(value))
	return reflectValue.Kind() == reflect.Map
}

// IsStruct checks whether `value` is of type struct.
func IsStruct(value interface{}) bool {
	reflectValue := reflect.Indirect(reflect.ValueOf(value))
	return reflectValue.Kind() == reflect.Struct
}
