// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package errors provides rich functionalities to manipulate errors.
//
// For maintainers, please note that this package is a basic package,
// which SHOULD NOT import extra packages except standard packages
// and internal packages to avoid cyclic imports.
package errors

import (
	"github.com/focela/min/errors/code"
)

// IsError defines the behavior for checking if an error matches a target error.
type IsError interface {
	Error() string
	Is(target error) bool
}

// EqualError defines the behavior for checking if two errors are equal.
type EqualError interface {
	Error() string
	Equal(target error) bool
}

// Coder defines the behavior for retrieving an error's code.
type Coder interface {
	Error() string
	Code() code.Code
}

// Stacker defines the behavior for retrieving an error's stack trace.
type Stacker interface {
	Error() string
	Stack() string
}

// Causer defines the behavior for retrieving the underlying cause of an error.
type Causer interface {
	Error() string
	Cause() error
}

// CurrentError defines the behavior for retrieving the current error.
type CurrentError interface {
	Error() string
	Current() error
}

// Unwrapper defines the behavior for unwrapping an error to its underlying cause.
type Unwrapper interface {
	Error() string
	Unwrap() error
}

const (
	// commaSeparatorSpace is the comma separator with a space.
	commaSeparatorSpace = ", "
)
