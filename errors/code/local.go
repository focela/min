// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package code

import (
	"fmt"
)

// localCode implements the Code interface and is intended for internal use only.
type localCode struct {
	code    int         // Error code, usually an integer.
	message string      // Brief description of the error.
	detail  interface{} // Optional field for additional error information.
}

// Code returns the integer value of the error code.
func (c localCode) Code() int {
	return c.code
}

// Message returns a brief description of the error.
func (c localCode) Message() string {
	return c.message
}

// Detail provides additional information for the error code.
func (c localCode) Detail() interface{} {
	return c.detail
}

// String converts the error code to a string representation.
func (c localCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}
