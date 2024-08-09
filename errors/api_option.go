// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"github.com/focela/aid/errors/code"
)

// ErrorOption is option for creating error.
type ErrorOption struct {
	Error error     // Wrapped error if any.
	Stack bool      // Whether recording stack information into error.
	Text  string    // Error text, which is created by New* functions.
	Code  code.Code // Error code if necessary.
}

// NewWithOption creates and returns a custom error with ErrorOption.
// It is the senior usage for creating error, which is often used internally in framework.
func NewWithOption(option ErrorOption) error {
	err := &Error{
		error: option.Error,
		text:  option.Text,
		code:  option.Code,
	}
	if option.Stack {
		err.stack = callers()
	}
	return err
}

// NewOption creates and returns a custom error with ErrorOption.
// Deprecated: use NewWithOption instead.
func NewOption(option ErrorOption) error {
	return NewWithOption(option)
}
