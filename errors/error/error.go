// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package errplus provides rich functionalities to manipulate errors.
//
// For maintainers, please note that this package is quite a basic package,
// which SHOULD NOT import extra packages except standard packages and internal packages,
// to avoid cycle imports.
package errplus

import (
	"github.com/focela/plume/errors/code"
)

// IsChecker is the interface for the Is feature.
type IsChecker interface {
	Error() string
	Is(target error) bool
}

// EqualComparer is the interface for the Equal feature.
type EqualComparer interface {
	Error() string
	Equal(target error) bool
}

// CodeRetriever is the interface for the Code feature.
type CodeRetriever interface {
	Error() string
	Code() code.Code
}

// StackTracer is the interface for the Stack feature.
type StackTracer interface {
	Error() string
	Stack() string
}

// CauseFinder is the interface for the Cause feature.
type CauseFinder interface {
	Error() string
	Cause() error
}

// CurrentErrorProvider is the interface for the Current feature.
type CurrentErrorProvider interface {
	Error() string
	Current() error
}

// ErrorUnwrapper is the interface for the Unwrap feature.
type ErrorUnwrapper interface {
	Error() string
	Unwrap() error
}
