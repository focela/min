// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package minerror

import (
	"fmt"

	"github.com/focela/min/errors/mincode"
)

// New creates and returns an error which is formatted from given text.
func New(text string) error {
	return &Error{
		stack: callers(),
		text:  text,
		code:  mincode.CodeNil,
	}
}

// Newf returns an error that formats as the given format and args.
func Newf(format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  mincode.CodeNil,
	}
}

// NewWithSkip creates and returns an error which is formatted from given text.
// The parameter `skip` specifies how many stack frames to skip when capturing the stack trace.
func NewWithSkip(skip int, text string) error {
	return &Error{
		stack: callers(skip),
		text:  text,
		code:  mincode.CodeNil,
	}
}

// NewWithSkipf returns an error that formats as the given format and args.
// The parameter `skip` specifies how many stack frames to skip when capturing the stack trace.
func NewWithSkipf(skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  mincode.CodeNil,
	}
}

// Wrap wraps error with text and inherits the error code from the wrapped error.
// It returns nil if the provided error is nil.
func Wrap(err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  text,
		code:  Code(err),
	}
}

// Wrapf wraps error with text formatted with the provided format and args,
// and inherits the error code from the wrapped error. It returns nil if the provided error is nil.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  Code(err),
	}
}

// WrapWithSkip wraps error with text and inherits the error code from the wrapped error.
// The parameter `skip` specifies how many stack frames to skip when capturing the stack trace.
func WrapWithSkip(skip int, err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  text,
		code:  Code(err),
	}
}

// WrapWithSkipf wraps error with text formatted with the provided format and args,
// and inherits the error code from the wrapped error. The parameter `skip` specifies
// how many stack frames to skip when capturing the stack trace.
func WrapWithSkipf(skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  Code(err),
	}
}
