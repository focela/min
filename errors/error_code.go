// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"github.com/focela/aid/errors/code"
)

// Code returns the error code.
// It returns CodeNil if it has no error code.
func (err *Error) Code() code.Code {
	if err == nil {
		return code.CodeNil
	}
	if err.code == code.CodeNil {
		unwrappedErr := err.Unwrap()
		if unwrappedErr == nil {
			return code.CodeNil
		}
		return Code(unwrappedErr)
	}
	return err.code
}

// SetCode updates the internal code with given code.
func (err *Error) SetCode(code code.Code) {
	if err == nil {
		return
	}
	err.code = code
}
