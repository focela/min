// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package error

import (
	"fmt"
	"strings"

	"github.com/focela/orca/pkg/errors/code"
)

// NewCode creates and returns an error that has error code and given text.
func NewCode(c code.Code, text ...string) error {
	return &Error{
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  c,
	}
}

// NewCodef returns an error that has error code and formats as the given format and args.
func NewCodef(c code.Code, format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  c,
	}
}

// NewCodeSkip creates and returns an error which has error code and is formatted from given text.
// The parameter `skip` specifies the stack callers skipped amount.
func NewCodeSkip(c code.Code, skip int, text ...string) error {
	return &Error{
		stack: callers(skip),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  c,
	}
}

// NewCodeSkipf returns an error that has error code and formats as the given format and args.
// The parameter `skip` specifies the stack callers skipped amount.
func NewCodeSkipf(c code.Code, skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  c,
	}
}

// WrapCode wraps error with code and text.
// It returns nil if the given err is nil.
func WrapCode(c code.Code, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  c,
	}
}

// WrapCodef wraps error with code and format specifier.
// It returns nil if the given `err` is nil.
func WrapCodef(c code.Code, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  c,
	}
}

// WrapCodeSkip wraps error with code and text.
// It returns nil if the given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
func WrapCodeSkip(c code.Code, skip int, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  c,
	}
}

// WrapCodeSkipf wraps error with code and text that is formatted with given format and args.
// It returns nil if the given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
func WrapCodeSkipf(c code.Code, skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  c,
	}
}

// Code returns the error code of the current error.
// It returns `CodeNil` if it has no error code or it does not implement the Code interface.
func Code(err error) code.Code {
	if err == nil {
		return code.CodeNil
	}
	if e, ok := err.(ErrorCode); ok {
		return e.Code()
	}
	if e, ok := err.(ErrorUnwrap); ok {
		return Code(e.Unwrap())
	}
	return code.CodeNil
}

// HasCode checks and reports whether `err` has `code` in its chaining errors.
func HasCode(err error, c code.Code) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(ErrorCode); ok {
		return c == e.Code()
	}
	if e, ok := err.(ErrorUnwrap); ok {
		return HasCode(e.Unwrap(), c)
	}
	return false
}
