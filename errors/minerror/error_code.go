// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package minerror

import (
	"github.com/focela/min/errors/mincode"
)

// Code returns the error code.
// It returns CodeNil if the current error has no code, and checks for the wrapped error's code.
func (err *Error) Code() mincode.Code {
	if err == nil {
		return mincode.CodeNil
	}
	if err.code == mincode.CodeNil {
		// Recursively check the wrapped error for its code.
		return Code(err.Unwrap())
	}
	return err.code
}

// SetCode updates the internal code with the given code.
// If the provided code is CodeNil, the error code will not be updated.
func (err *Error) SetCode(code mincode.Code) {
	if err == nil || code == mincode.CodeNil {
		return
	}
	err.code = code
}
