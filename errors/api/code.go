// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package api

import (
	"fmt"
	"strings"

	"github.com/focela/min/errors"
	"github.com/focela/min/errors/code"
)

// NewCode creates and returns an error with the specified error code and text.
func NewCode(code code.Code, text ...string) error {
	return &errors.Error{
		Stack: Callers(),
		Text:  strings.Join(text, errors.CommaSeparatorSpace),
		Code:  code,
	}
}

// NewCodef returns an error with the specified error code and a formatted message.
func NewCodef(code code.Code, format string, args ...interface{}) error {
	return &errors.Error{
		Stack: Callers(),
		Text:  fmt.Sprintf(format, args...),
		Code:  code,
	}
}

// NewCodeSkip creates and returns an error with the specified error code and text.
// The `skip` parameter specifies the number of stack callers to skip.
func NewCodeSkip(code code.Code, skip int, text ...string) error {
	return &errors.Error{
		Stack: Callers(skip),
		Text:  strings.Join(text, errors.CommaSeparatorSpace),
		Code:  code,
	}
}

// NewCodeSkipf returns an error with the specified error code and a formatted message.
// The `skip` parameter specifies the number of stack callers to skip.
func NewCodeSkipf(code code.Code, skip int, format string, args ...interface{}) error {
	return &errors.Error{
		Stack: Callers(skip),
		Text:  fmt.Sprintf(format, args...),
		Code:  code,
	}
}

// WrapCode wraps an error with the specified code and text.
// It returns nil if the given error is nil.
func WrapCode(code code.Code, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &errors.Error{
		Err:   err,
		Stack: Callers(),
		Text:  strings.Join(text, errors.CommaSeparatorSpace),
		Code:  code,
	}
}

// WrapCodef wraps an error with the specified code and a formatted message.
// It returns nil if the given error is nil.
func WrapCodef(code code.Code, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &errors.Error{
		Err:   err,
		Stack: Callers(),
		Text:  fmt.Sprintf(format, args...),
		Code:  code,
	}
}

// WrapCodeSkip wraps an error with the specified code and text.
// It returns nil if the given error is nil.
// The `skip` parameter specifies the number of stack callers to skip.
func WrapCodeSkip(code code.Code, skip int, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &errors.Error{
		Err:   err,
		Stack: Callers(skip),
		Text:  strings.Join(text, errors.CommaSeparatorSpace),
		Code:  code,
	}
}

// WrapCodeSkipf wraps an error with the specified code and a formatted message.
// It returns nil if the given error is nil.
// The `skip` parameter specifies the number of stack callers to skip.
func WrapCodeSkipf(code code.Code, skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &errors.Error{
		Err:   err,
		Stack: Callers(skip),
		Text:  fmt.Sprintf(format, args...),
		Code:  code,
	}
}

// Code returns the error code of the current error.
// It returns `CodeNil` if there is no error code or if the error does not implement the Coder interface.
func Code(err error) code.Code {
	if err == nil {
		return code.CodeNil
	}
	if e, ok := err.(errors.Coder); ok {
		return e.Code()
	}
	if e, ok := err.(errors.Unwrapper); ok {
		return Code(e.Unwrap())
	}
	return code.CodeNil
}

// HasCode checks whether the given error or any error in its chain has the specified code.
func HasCode(err error, code code.Code) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(errors.Coder); ok {
		return code == e.Code()
	}
	if e, ok := err.(errors.Unwrapper); ok {
		return HasCode(e.Unwrap(), code)
	}
	return false
}
