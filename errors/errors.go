// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package errors provides rich functionalities to manipulate errors.
//
// For maintainers, please note that this package is quite a basic package,
// which SHOULD NOT import extra packages except standard packages and internal packages,
// to avoid cycle imports.
package errors

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/focela/min/errors/api"
	"github.com/focela/min/errors/code"
)

// IsComparer is the interface for Is feature.
type IsComparer interface {
	Error() string
	Is(target error) bool
}

// Equaler is the interface for Equal feature.
type Equaler interface {
	Error() string
	Equal(target error) bool
}

// Coder is the interface for Code feature.
type Coder interface {
	Error() string
	Code() code.Code
}

// Stacker is the interface for Stack feature.
type Stacker interface {
	Error() string
	Stack() string
}

// Causer is the interface for Cause feature.
type Causer interface {
	Error() string
	Cause() error
}

// Currenter is the interface for Current feature.
type Currenter interface {
	Error() string
	Current() error
}

// Unwrapper is the interface for Unwrap feature.
type Unwrapper interface {
	Error() string
	Unwrap() error
}

// Error is a custom error structure for additional features.
type Error struct {
	Err   error         // Wrapped error.
	Stack api.CallStack // Stack array, which records the stack information when this error is created or wrapped.
	Text  string        // Custom error text when Error is created, might be empty when its code is not nil.
	Code  code.Code     // Error code if necessary.
}

const (
	// CommaSeparatorSpace defines a comma followed by a space.
	CommaSeparatorSpace = ", "
	// StackFilterKeyLocal is the filtering key for the current error module paths.
	StackFilterKeyLocal = "/errors/errors"
)

var (
	// goRootForFilter is used for stack filtering in development environment purposes.
	goRootForFilter = runtime.GOROOT()
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}
}

// Error implements the Error interface and returns all error information as a string.
func (err *Error) Error() string {
	if err == nil {
		return ""
	}
	var sb strings.Builder
	if err.Text != "" {
		sb.WriteString(err.Text)
	} else if err.Code != nil {
		sb.WriteString(err.Code.Message())
	}
	if err.Err != nil {
		if sb.Len() > 0 {
			sb.WriteString(": ")
		}
		sb.WriteString(err.Err.Error())
	}
	return sb.String()
}

// Cause returns the root cause error.
func (err *Error) Cause() error {
	if err == nil {
		return nil
	}
	if err.Err != nil {
		if e, ok := err.Err.(*Error); ok {
			return e.Cause()
		} else if e, ok := err.Err.(Causer); ok {
			return e.Cause()
		} else {
			return err.Err
		}
	}
	return nil
}

// Current creates and returns the current level error.
// It returns nil if the current level error is nil.
func (err *Error) Current() error {
	if err == nil {
		return nil
	}
	return &Error{
		Err:   nil,
		Stack: err.Stack,
		Text:  err.Text,
		Code:  err.Code,
	}
}

// GetNextError is an alias of function `Next`.
// It is implemented for stdlib errors.Unwrap from Go version 1.17.
func (err *Error) GetNextError() error {
	if err == nil {
		return nil
	}
	return err.Err
}

// Equal reports whether the current error `err` equals the error `target`.
// Note that in the default comparison for `Error`,
// errors are considered the same if both the `code` and `text` are identical.
func (err *Error) Equal(target error) bool {
	if err == target {
		return true
	}
	if err.Code != api.Code(target) {
		return false
	}
	if err.Text != fmt.Sprintf(`%-s`, target) {
		return false
	}
	return true
}

// Is reports whether the current error `err` has the error `target` in its chaining errors.
// It is implemented for stdlib errors.Is from Go version 1.17.
func (err *Error) Is(target error) bool {
	if api.Equal(err, target) {
		return true
	}
	nextErr := err.GetNextError()
	if nextErr == nil {
		return false
	}
	if api.Equal(nextErr, target) {
		return true
	}
	if e, ok := nextErr.(IsComparer); ok {
		return e.Is(target)
	}
	return false
}
