// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package errors provides functionalities for manipulating errors.
//
// Note for maintainers:
// This is a basic package. Do NOT import any external packages
// other than standard and internal packages to avoid circular dependencies.
package errors

import (
	"github.com/focela/min/errors/code"
)

// IsError defines the Is feature interface.
type IsError interface {
	Error() string
	Is(target error) bool
}

// EqualError defines the Equal feature interface.
type EqualError interface {
	Error() string
	Equal(target error) bool
}

// CodeError defines the Code feature interface.
type CodeError interface {
	Error() string
	Code() code.Code
}

// StackError defines the Stack feature interface.
type StackError interface {
	Error() string
	Stack() string
}

// CauseError defines the Cause feature interface.
type CauseError interface {
	Error() string
	Cause() error
}

// CurrentError defines the Current feature interface.
type CurrentError interface {
	Error() string
	Current() error
}

// UnwrapError defines the Unwrap feature interface.
type UnwrapError interface {
	Error() string
	Unwrap() error
}

const (
	// CommaSeparatorSpace is the comma separator with space.
	CommaSeparatorSpace = ", "
)
