// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package minerror

import (
	"fmt"
	"strings"

	"github.com/focela/min/errors/mincode"
)

// NewCode creates and returns an error that has error code and given text.
func NewCode(code mincode.Code, text ...string) error {
	return &Error{
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// NewCodef returns an error that has error code and formats as the given format and args.
func NewCodef(code mincode.Code, format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// NewCodeWithSkip creates and returns an error which has error code and is formatted from given text.
// The parameter `skip` specifies how many stack frames to skip when capturing the stack trace.
func NewCodeWithSkip(code mincode.Code, skip int, text ...string) error {
	return &Error{
		stack: callers(skip),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// NewCodeWithSkipf returns an error that has error code and formats as the given format and args.
// The parameter `skip` specifies how many stack frames to skip when capturing the stack trace.
func NewCodeWithSkipf(code mincode.Code, skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCode wraps error with code and text.
// It returns nil if given err is nil.
func WrapCode(code mincode.Code, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// WrapCodef wraps error with code and format specifier.
// It returns nil if given `err` is nil.
func WrapCodef(code mincode.Code, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCodeWithSkip wraps error with code and text.
// It returns nil if given err is nil.
// The parameter `skip` specifies how many stack frames to skip when capturing the stack trace.
func WrapCodeWithSkip(code mincode.Code, skip int, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// WrapCodeWithSkipf wraps error with code and text that is formatted with given format and args.
// It returns nil if given err is nil.
// The parameter `skip` specifies how many stack frames to skip when capturing the stack trace.
func WrapCodeWithSkipf(code mincode.Code, skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// Code returns the error code of current error.
// It returns `CodeNil` if it has no error code, or it does not implement interface Code.
func Code(err error) mincode.Code {
	if err == nil {
		return mincode.CodeNil
	}
	if e, ok := err.(CodeRetriever); ok {
		return e.Code()
	}
	if e, ok := err.(Unwrapper); ok {
		return Code(e.Unwrap())
	}
	return mincode.CodeNil
}

// HasCode checks and reports whether `err` has `code` in its chaining errors.
func HasCode(err error, code mincode.Code) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(CodeRetriever); ok {
		return code == e.Code()
	}
	if e, ok := err.(Unwrapper); ok {
		return HasCode(e.Unwrap(), code)
	}
	return false
}
