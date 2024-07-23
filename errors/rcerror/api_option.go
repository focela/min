// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcerror

import "github.com/focela/ratcatcher/errors/rccode"

// Option is option for creating error.
type Option struct {
	Error error       // Wrapped error if any.
	Stack bool        // Whether recording stack information into error.
	Text  string      // Error text, which is created by New* functions.
	Code  rccode.Code // Error code if necessary.
}

// NewWithOption creates and returns a custom error with Option.
// It is the senior usage for creating error, which is often used internally in framework.
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
