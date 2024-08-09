// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"reflect"
)

// IsNilCallable determines if the reflect.Value can safely call the IsNil method.
// This function prevents panics that occur when calling reflect.Value.IsNil on unsupported types.
func IsNilCallable(v interface{}) bool {
	rv, ok := v.(reflect.Value)
	if !ok {
		return false
	}
	switch rv.Kind() {
	case reflect.Interface, reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return true
	default:
		return false
	}
}
