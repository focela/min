// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package rctype provides high performance and concurrent-safe basic variable types.
package rctype

// New is alias of NewInterface.
// See NewInterface.
func New(value ...interface{}) *Interface {
	return NewInterface(value...)
}
