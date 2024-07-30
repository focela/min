// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package error provides rich functionalities to manipulate errors.
//
// For maintainers, please note that
// this package is a basic package, which SHOULD NOT import extra packages
// except standard packages and internal packages, to avoid cycle imports.
package error

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/focela/orca/pkg/errors/code"
)

// ErrorIs is the interface for Is feature.
type ErrorIs interface {
	Error() string
	Is(target error) bool
}

// ErrorEqual is the interface for Equal feature.
type ErrorEqual interface {
	Error() string
	Equal(target error) bool
}

// ErrorCode is the interface for Code feature.
type ErrorCode interface {
	Error() string
	Code() code.Code
}

// ErrorStack is the interface for Stack feature.
type ErrorStack interface {
	Error() string
	Stack() string
}

// ErrorCause is the interface for Cause feature.
type ErrorCause interface {
	Error() string
	Cause() error
}

// ErrorCurrent is the interface for Current feature.
type ErrorCurrent interface {
	Error() string
	Current() error
}

// ErrorUnwrap is the interface for Unwrap feature.
type ErrorUnwrap interface {
	Error() string
	Unwrap() error
}

// Error is a custom error for additional features.
type Error struct {
	error error     // Wrapped error.
	stack stack     // Stack array, which records the stack information when this error is created or wrapped.
	text  string    // Custom Error text when Error is created, might be empty when its code is not nil.
	code  code.Code // Error code if necessary.
}

const (
	// commaSeparatorSpace is the comma separator with space.
	commaSeparatorSpace = ", "
	// stackFilterKeyLocal is the filtering key for current error module paths.
	stackFilterKeyLocal = "/errors/error/error"
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

// Error implements the interface of Error, it returns all the error as a string.
func (err *Error) Error() string {
	if err == nil {
		return ""
	}
	errStr := err.text
	if errStr == "" && err.code != nil {
		errStr = err.code.Message()
	}
	if err.error != nil {
		if errStr != "" {
			errStr += ": "
		}
		errStr += err.error.Error()
	}
	return errStr
}

// Cause returns the root cause error.
func (err *Error) Cause() error {
	if err == nil {
		return nil
	}
	loop := err
	for loop != nil {
		if loop.error != nil {
			if e, ok := loop.error.(*Error); ok {
				// Internal Error struct.
				loop = e
			} else if e, ok := loop.error.(ErrorCause); ok {
				// Other Error that implements ErrorCause interface.
				return e.Cause()
			} else {
				return loop.error
			}
		} else {
			// To be compatible with cases of https://github.com/pkg/errors.
			return errors.New(loop.text)
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
		error: nil,
		stack: err.stack,
		text:  err.text,
		code:  err.code,
	}
}

// Unwrap is an alias of the function `Next`.
// It is just for implementation of stdlib errors.Unwrap from Go version 1.17.
func (err *Error) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.error
}

// Equal reports whether the current error `err` equals to the error `target`.
// Please note that, in default comparison for `Error`,
// the errors are considered the same if both the `code` and `text` of them are the same.
func (err *Error) Equal(target error) bool {
	if err == target {
		return true
	}
	// Code should be the same.
	// Note that if both errors have `nil` code, they are also considered equal.
	if err.code != Code(target) {
		return false
	}
	// Text should be the same.
	if err.text != fmt.Sprintf(`%-s`, target) {
		return false
	}
	return true
}

// Is reports whether the current error `err` has error `target` in its chaining errors.
// It is just for implementation of stdlib errors.Is from Go version 1.17.
func (err *Error) Is(target error) bool {
	if Equal(err, target) {
		return true
	}
	nextErr := err.Unwrap()
	if nextErr == nil {
		return false
	}
	if Equal(nextErr, target) {
		return true
	}
	if e, ok := nextErr.(ErrorIs); ok {
		return e.Is(target)
	}
	return false
}
