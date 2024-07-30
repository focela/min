// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package error

import "github.com/focela/orca/pkg/errors/code"

// Code returns the error code.
// It returns CodeNil if it has no error code.
func (err *Error) Code() code.Code {
	if err == nil {
		return code.CodeNil
	}
	if err.code == code.CodeNil {
		return Code(err.Unwrap())
	}
	return err.code
}

// SetCode updates the internal code with the given code.
func (err *Error) SetCode(c code.Code) {
	if err == nil {
		return
	}
	err.code = c
}
