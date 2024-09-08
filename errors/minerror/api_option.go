// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package minerror

import (
	"github.com/focela/min/errors/mincode"
)

// Option represents the options for creating a custom error.
type Option struct {
	Error error        // Wrapped error, if any.
	Stack bool         // Whether to record stack information into the error.
	Text  string       // Error message, used in New* functions.
	Code  mincode.Code // Error code, if applicable.
}

// NewErrorWithOption creates and returns a custom error using the provided options.
// This function is primarily used for internal error handling within the framework.
func NewErrorWithOption(option Option) error {
	err := &Error{
		error: option.Error,
		text:  option.Text,
		code:  option.Code,
	}
	// Record stack trace if Stack is set to true.
	if option.Stack {
		err.stack = callers()
	}
	return err
}

// NewOption creates and returns a custom error using Option.
// Deprecated: use NewErrorWithOption instead.
func NewOption(option Option) error {
	return NewErrorWithOption(option)
}
