// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package code

import "fmt"

// errorCode is an implementation of the Code interface for internal usage only.
type errorCode struct {
	code    int         // Error code, typically an integer.
	message string      // Brief message for this error code.
	detail  interface{} // Extension field for additional error information.
}

// Code returns the integer value of the current error code.
func (e errorCode) Code() int {
	return e.code
}

// Message returns a brief message for the current error code.
func (e errorCode) Message() string {
	return e.message
}

// Detail returns detailed information about the current error code.
// It is primarily designed as an extension field.
func (e errorCode) Detail() interface{} {
	return e.detail
}

// String returns the current error code as a formatted string.
func (e errorCode) String() string {
	if e.message != "" {
		if e.detail != nil {
			return fmt.Sprintf(`%d:%s %v`, e.code, e.message, e.detail)
		}
		return fmt.Sprintf(`%d:%s`, e.code, e.message)
	}
	return fmt.Sprintf(`%d`, e.code)
}
