// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import "reflect"

// CanCallIsNil Can reflect.Value call reflect.Value.IsNil.
// It can avoid reflect.Value.IsNil panics.
func CanCallIsNil(v interface{}) bool {
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
