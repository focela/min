// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"reflect"
)

// IsArray checks whether the given value is an array or slice.
// Note that it uses reflect internally to implement this feature.
func IsArray(value interface{}) bool {
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}
