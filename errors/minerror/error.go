// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package minerror provides rich functionalities to manipulate errors.
//
// For maintainers, please note that this package is a basic package,
// and SHOULD NOT import extra packages except standard and internal packages
// to avoid cycle imports.
package minerror

import (
	"errors"
	"runtime"
	"strings"

	"github.com/focela/min/errors/mincode"
)

// IsChecker defines an interface for checking if an error matches a target error.
type IsChecker interface {
	Error() string
	Is(target error) bool
}

// EqualChecker defines an interface for checking if an error equals a target error.
type EqualChecker interface {
	Error() string
	Equal(target error) bool
}

// CodeRetriever defines an interface for retrieving an error code.
type CodeRetriever interface {
	Error() string
	Code() mincode.Code
}

// StackTracer defines an interface for retrieving a stack trace of the error.
type StackTracer interface {
	Error() string
	Stack() string
}

// CauseRetriever defines an interface for retrieving the cause of the error.
type CauseRetriever interface {
	Error() string
	Cause() error
}

// CurrentRetriever defines an interface for retrieving the current error.
type CurrentRetriever interface {
	Error() string
	Current() error
}

// Unwrapper defines an interface for unwrapping an error to its underlying cause.
type Unwrapper interface {
	Error() string
	Unwrap() error
}

type Error struct {
	error error        // Wrapped error.
	stack stack        // Stack array, which records the stack information when this error is created or wrapped.
	text  string       // Custom Error text when Error is created, might be empty when its code is not nil.
	code  mincode.Code // Error code if necessary.
}

const (
	// commaSeparatorSpace is the comma separator with space.
	commaSeparatorSpace = ", "
	// Filtering key for current error module paths.
	stackFilterModulePath = "/errors/minerror/minerror"
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

// Error returns a string representation of the error, including custom text and wrapped error messages, if available.
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

// Cause returns the root cause error, iterating through the chain of wrapped errors.
func (err *Error) Cause() error {
	if err == nil {
		return nil
	}
	currentErr := err
	for currentErr != nil {
		if currentErr.error != nil {
			if e, ok := currentErr.error.(*Error); ok {
				currentErr = e
			} else if e, ok := currentErr.error.(CauseRetriever); ok {
				return e.Cause()
			} else {
				return currentErr.error
			}
		} else {
			return errors.New(currentErr.text)
		}
	}
	return nil
}

// BaseError returns the current level error without any wrapped errors.
func (err *Error) BaseError() error {
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

// Unwrap returns the wrapped error.
func (err *Error) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.error
}

// Equal checks if the current error is equal to the target error based on error code and text.
func (err *Error) Equal(target error) bool {
	if err == target {
		return true
	}
	if err.code != Code(target) {
		return false
	}
	if err.text != target.Error() {
		return false
	}
	return true
}

// Is checks if the target error exists in the chain of errors.
func (err *Error) Is(target error) bool {
	if err.Equal(target) {
		return true
	}
	nextErr := err.Unwrap()
	if nextErr == nil {
		return false
	}
	if errors.Is(nextErr, target) {
		return true
	}
	if e, ok := nextErr.(IsChecker); ok {
		return e.Is(target)
	}
	return false
}
