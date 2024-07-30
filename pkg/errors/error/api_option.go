// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package error

import "github.com/focela/orca/pkg/errors/code"

// Option is an option for creating an error.
type Option struct {
	Error error     // Wrapped error if any.
	Stack bool      // Whether to record stack information into the error.
	Text  string    // Error text, which is created by New* functions.
	Code  code.Code // Error code if necessary.
}

// NewWithOption creates and returns a custom error with Option.
// It is the advanced usage for creating errors, which is often used internally in the framework.
func NewWithOption(option Option) error {
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

// NewOption creates and returns a custom error with Option.
// Deprecated: use NewWithOption instead.
func NewOption(option Option) error {
	return NewWithOption(option)
}
