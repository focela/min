// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package api

import (
	"fmt"

	"github.com/focela/min/errors"
	"github.com/focela/min/errors/code"
)

// New creates and returns an error formatted from the given text.
func New(text string) error {
	return &errors.Error{
		Stack: Callers(),
		Text:  text,
		Code:  code.CodeNil,
	}
}

// Newf returns an error formatted according to the given format and arguments.
func Newf(format string, args ...interface{}) error {
	return &errors.Error{
		Stack: Callers(),
		Text:  fmt.Sprintf(format, args...),
		Code:  code.CodeNil,
	}
}

// NewSkip creates and returns an error formatted from the given text.
// The `skip` parameter specifies the number of stack callers to skip.
func NewSkip(skip int, text string) error {
	return &errors.Error{
		Stack: Callers(skip),
		Text:  text,
		Code:  code.CodeNil,
	}
}

// NewSkipf returns an error formatted according to the given format and arguments.
// The `skip` parameter specifies the number of stack callers to skip.
func NewSkipf(skip int, format string, args ...interface{}) error {
	return &errors.Error{
		Stack: Callers(skip),
		Text:  fmt.Sprintf(format, args...),
		Code:  code.CodeNil,
	}
}

// Wrap wraps an error with text. It returns nil if the given error is nil.
// The error code of the wrapped error is preserved.
func Wrap(err error, text string) error {
	if err == nil {
		return nil
	}
	return &errors.Error{
		Err:   err,
		Stack: Callers(),
		Text:  text,
		Code:  Code(err),
	}
}

// Wrapf returns an error with a stack trace at the point Wrapf is called and formats the error message.
// It returns nil if the given error is nil.
// The error code of the wrapped error is preserved.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &errors.Error{
		Err:   err,
		Stack: Callers(),
		Text:  fmt.Sprintf(format, args...),
		Code:  Code(err),
	}
}

// WrapSkip wraps an error with text. It returns nil if the given error is nil.
// The `skip` parameter specifies the number of stack callers to skip.
// The error code of the wrapped error is preserved.
func WrapSkip(skip int, err error, text string) error {
	if err == nil {
		return nil
	}
	return &errors.Error{
		Err:   err,
		Stack: Callers(skip),
		Text:  text,
		Code:  Code(err),
	}
}

// WrapSkipf wraps an error with a formatted message. It returns nil if the given error is nil.
// The `skip` parameter specifies the number of stack callers to skip.
// The error code of the wrapped error is preserved.
func WrapSkipf(skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &errors.Error{
		Err:   err,
		Stack: Callers(skip),
		Text:  fmt.Sprintf(format, args...),
		Code:  Code(err),
	}
}
