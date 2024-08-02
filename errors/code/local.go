// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package code

import "fmt"

// localCode implements the Code interface for internal use only.
type localCode struct {
	code    int         // Error code, typically an integer.
	message string      // A brief message associated with the error code.
	detail  interface{} // An extension field for additional details about the error.
}

// Code returns the integer representation of the current error code.
func (c localCode) Code() int {
	return c.code
}

// Message returns the brief message associated with the current error code.
func (c localCode) Message() string {
	return c.message
}

// Detail provides additional information about the current error code.
// This field is mainly designed for extending the error code with more context.
func (c localCode) Detail() interface{} {
	return c.detail
}

// String returns a string representation of the current error code.
// If both message and detail are present, they are included in the format "code:message detail".
func (c localCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d: %s %v`, c.code, c.message, c.detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d: %s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}
