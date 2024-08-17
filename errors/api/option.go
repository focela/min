// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package api

import (
	"github.com/focela/min/errors"
	"github.com/focela/min/errors/code"
)

// Option defines the parameters for creating an error.
type Option struct {
	Err   error     // Wrapped error if any.
	Stack bool      // Whether to record stack information into the error.
	Text  string    // Error text, typically created by New* functions.
	Code  code.Code // Error code if necessary.
}

// NewWithOption creates and returns a custom error with Option.
// It is an advanced method for creating an error, often used internally in the framework.
func NewWithOption(option Option) error {
	err := &errors.Error{
		Err:  option.Err,
		Text: option.Text,
		Code: option.Code,
	}
	if option.Stack {
		err.Stack = Callers()
	}
	return err
}

// NewOption creates and returns a custom error with Option.
// Deprecated: Use NewWithOption instead.
func NewOption(option Option) error {
	return NewWithOption(option)
}
