// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package code

import (
	"fmt"
)

// localCode is an implementation of the Code interface for internal use only.
type localCode struct {
	errorCode    int         // Integer representing the error code.
	errorMessage string      // Brief message associated with this error code.
	errorDetail  interface{} // Extension field for additional error details.
}

// Code returns the integer value of the current error code.
func (c localCode) Code() int {
	return c.errorCode
}

// Message returns a brief message for the current error code.
func (c localCode) Message() string {
	return c.errorMessage
}

// Detail returns detailed information about the current error code.
// This is primarily designed as an extension field for additional error details.
func (c localCode) Detail() interface{} {
	return c.errorDetail
}

// String returns the current error code as a formatted string.
func (c localCode) String() string {
	if c.errorDetail != nil {
		return fmt.Sprintf(`%d: %s %v`, c.errorCode, c.errorMessage, c.errorDetail)
	}
	if c.errorMessage != "" {
		return fmt.Sprintf(`%d: %s`, c.errorCode, c.errorMessage)
	}
	return fmt.Sprintf(`%d`, c.errorCode)
}
