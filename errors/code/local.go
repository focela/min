// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package code

import (
	"fmt"
)

// localCode is an implementation of the Code interface for internal usage only.
type localCode struct {
	code    int         // Error code, usually an integer.
	message string      // Brief message for this error code.
	detail  interface{} // Designed as an extension field for additional error details.
}

// Code returns the integer value of the current error code.
func (c localCode) Code() int {
	return c.code
}

// Message returns a brief message for the current error code.
func (c localCode) Message() string {
	return c.message
}

// Detail returns additional information related to the current error code.
func (c localCode) Detail() interface{} {
	return c.detail
}

// String returns a string representation of the current error code.
func (c localCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d: %s %v`, c.code, c.message, c.detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d: %s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}
