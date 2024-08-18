// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package code

import (
	"fmt"
)

// localCode is an implementation of the Code interface for internal usage only.
type localCode struct {
	code    int         // Error code, typically an integer.
	message string      // Brief message associated with this error code.
	detail  interface{} // An extension field for additional error information.
}

// Code returns the integer value of the current error code.
func (c localCode) Code() int {
	return c.code
}

// Message returns the brief message associated with the current error code.
func (c localCode) Message() string {
	return c.message
}

// Detail returns the additional information related to the current error code,
// which serves as an extension field for further details.
func (c localCode) Detail() interface{} {
	return c.detail
}

// String returns the current error code as a formatted string.
func (c localCode) String() string {
	var result string
	if c.detail != nil {
		result = fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.detail)
	} else if c.message != "" {
		result = fmt.Sprintf(`%d:%s`, c.code, c.message)
	} else {
		result = fmt.Sprintf(`%d`, c.code)
	}
	return result
}
