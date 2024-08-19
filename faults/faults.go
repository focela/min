// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package faults provides rich functionalities to manipulate errors.
//
// Note: This is a basic package and should only import standard
// or internal packages to avoid cyclic dependencies.
package faults

import (
	"github.com/focela/min/faults/code"
)

// IsChecker provides the Is feature for errors.
type IsChecker interface {
	Error() string
	Is(target error) bool
}

// EqualityChecker provides the Equal feature for errors.
type EqualityChecker interface {
	Error() string
	Equal(target error) bool
}

// Coder provides the Code feature for errors.
type Coder interface {
	Error() string
	Code() code.Code
}

// StackTracer provides the Stack feature for errors.
type StackTracer interface {
	Error() string
	Stack() string
}

// CauseFinder provides the Cause feature for errors.
type CauseFinder interface {
	Error() string
	Cause() error
}

// CurrentError provides the Current feature for errors.
type CurrentError interface {
	Error() string
	Current() error
}

// Unwrapper provides the Unwrap feature for errors.
type Unwrapper interface {
	Error() string
	Unwrap() error
}

const (
	// CommaSeparatorSpace is the comma separator with space.
	CommaSeparatorSpace = ", "
)
