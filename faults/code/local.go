// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package code

import (
	"fmt"
)

// localCode implements the Code interface for internal use only.
type localCode struct {
	code    int         // Error code, usually an integer.
	message string      // Brief message for this error code.
	detail  interface{} // As an interface type, this serves as an extension field for the error code.
}

// Code returns the current error code as an integer.
func (c localCode) Code() int {
	return c.code
}

// Message returns a brief message associated with the current error code.
func (c localCode) Message() string {
	return c.message
}

// Detail returns detailed information about the current error code.
// This is mainly designed as an extension field for the error code.
func (c localCode) Detail() interface{} {
	return c.detail
}

// String returns the current error code as a formatted string.
func (c localCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d: %s %v`, c.code, c.message, c.detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d: %s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}
