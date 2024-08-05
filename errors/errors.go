// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package errors provides rich functionalities to manipulate errors.
//
// For maintainers, please note that
// this package is a basic package, which SHOULD NOT import extra packages
// except standard packages and internal packages, to avoid cycle imports.
package errors

import (
	"github.com/focela/plume/code"
)

// IsChecker is the interface for checking if an error matches a target error.
type IsChecker interface {
	Error() string
	Is(target error) bool
}

// EqualChecker is the interface for checking if an error is equal to a target error.
type EqualChecker interface {
	Error() string
	Equal(target error) bool
}

// Coder is the interface for retrieving an error code.
type Coder interface {
	Error() string
	Code() code.ErrorCode
}

// Stacker is the interface for retrieving a stack trace from an error.
type Stacker interface {
	Error() string
	Stack() string
}

// Causer is the interface for retrieving the cause of an error.
type Causer interface {
	Error() string
	Cause() error
}

// CurrentErrorRetriever is the interface for retrieving the current error.
type CurrentErrorRetriever interface {
	Error() string
	Current() error
}

// Unwrapper is the interface for unwrapping an error to its underlying cause.
type Unwrapper interface {
	Error() string
	Unwrap() error
}

const (
	// COMMA_SEPARATOR_SPACE is the comma separator with space.
	COMMA_SEPARATOR_SPACE = ", "
)
