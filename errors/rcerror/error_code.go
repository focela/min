// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcerror

import "github.com/focela/ratcatcher/errors/rccode"

// Code returns the error code.
// It returns CodeNil if it has no error code.
func (err *Error) Code() rccode.Code {
	if err == nil {
		return rccode.CodeNil
	}
	if err.code == rccode.CodeNil {
		return Code(err.Unwrap())
	}
	return err.code
}

// SetCode updates the internal code with given code.
func (err *Error) SetCode(code rccode.Code) {
	if err == nil {
		return
	}
	err.code = code
}
