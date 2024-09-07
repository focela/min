// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package mincode

import (
	"fmt"
)

// errorCode is an implementer for interface Code for internal usage only.
type errorCode struct {
	code    int         // Error code, usually an integer.
	message string      // Brief message for this error code.
	detail  interface{} // Extension field for detailed information of the error code.
}

// Code returns the integer representation of the error code.
func (c errorCode) Code() int {
	return c.code
}

// Message returns the brief message for the error code.
func (c errorCode) Message() string {
	return c.message
}

// Detail returns detailed information of the error code.
func (c errorCode) Detail() interface{} {
	return c.detail
}

// String returns the error code as a string.
func (c errorCode) String() string {
	switch {
	case c.detail != nil:
		return fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.detail)
	case c.message != "":
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	default:
		return fmt.Sprintf(`%d`, c.code)
	}
}
