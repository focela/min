// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package errors provides rich functionalities to manipulate errors.
//
// For maintainers, please very note that,
// this package is quite a basic package, which SHOULD NOT import extra packages
// except standard packages and internal packages, to avoid cycle imports.
package errors

import (
	"errors"
	"runtime"
	"strings"

	"github.com/focela/aid/errors/code"
)

// IsChecker is the interface for Is feature.
type IsChecker interface {
	Error() string
	Is(target error) bool
}

// EqualChecker is the interface for Equal feature.
type EqualChecker interface {
	Error() string
	Equal(target error) bool
}

// CodeChecker is the interface for Code feature.
type CodeChecker interface {
	Error() string
	Code() code.Code
}

// StackTracer is the interface for Stack feature.
type StackTracer interface {
	Error() string
	Stack() string
}

// CauseChecker is the interface for Cause feature.
type CauseChecker interface {
	Error() string
	Cause() error
}

// CurrentChecker is the interface for Current feature.
type CurrentChecker interface {
	Error() string
	Current() error
}

// Unwrapper is the interface for Unwrap feature.
type Unwrapper interface {
	Error() string
	Unwrap() error
}

// Error is custom error for additional features.
type Error struct {
	error error     // Wrapped error.
	stack stack     // Stack array, which records the stack information when this error is created or wrapped.
	text  string    // Custom Error text when Error is created, might be empty when its code is not nil.
	code  code.Code // Error code if necessary.
}

const (
	// commaSeparatorSpace is the comma separator with space.
	commaSeparatorSpace = ", "
	// Filtering key for current error module paths.
	stackFilterKeyLocal = "/errors/errors"
)

var (
	// goRootForFilter is used for stack filtering in development environment purpose.
	goRootForFilter = runtime.GOROOT()
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}
}

// Error implements the interface of Error, it returns all the error as string.
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
			} else if e, ok := loop.error.(CauseChecker); ok {
				// Other Error that implements CauseChecker interface.
				return e.Cause()
			} else {
				return loop.error
			}
		} else {
			// return loop
			//
			// To be compatible with Cause of https://github.com/pkg/errors.
			return errors.New(loop.text)
		}
	}
	return nil
}

// Current creates and returns the current level error.
// It returns nil if current level error is nil.
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

// Unwrap is alias of function `Next`.
// It is just for implements for stdlib errors.Unwrap from Go version 1.17.
func (err *Error) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.error
}

// Equal reports whether current error `err` equals to error `target`.
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
	if err.text != target.Error() {
		return false
	}
	return true
}

// Is reports whether current error `err` has error `target` in its chaining errors.
// It is just for implements for stdlib errors.Is from Go version 1.17.
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
	if e, ok := nextErr.(IsChecker); ok {
		return e.Is(target)
	}
	return false
}
