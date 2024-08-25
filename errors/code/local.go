// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package code

import (
	"fmt"
)

// localCode is an internal implementation of the Code interface.
type localCode struct {
	code    int         // Error code, usually an integer.
	message string      // Brief message for this error code.
	detail  interface{} // Extension field for additional error details.
}

// Code returns the integer number of the current error code.
func (c localCode) Code() int {
	return c.code
}

// Message returns the brief message for the current error code.
func (c localCode) Message() string {
	return c.message
}

// Detail returns the detailed information of the current error code,
// which is mainly designed as an extension field for the error code.
func (c localCode) Detail() interface{} {
	return c.detail
}

// String returns the current error code as a string.
func (c localCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}
