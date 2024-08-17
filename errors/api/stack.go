// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package api

import (
	"runtime"

	"github.com/focela/min/errors"
)

// CallStack represents a stack of program counters.
type CallStack []uintptr

const (
	// maxStackDepth defines the maximum stack depth for error backtraces.
	maxStackDepth = 64
)

// Cause returns the root cause of the error `err`.
func Cause(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(errors.Causer); ok {
		return e.Cause()
	}
	if e, ok := err.(errors.Unwrapper); ok {
		return Cause(e.Unwrap())
	}
	return err
}

// Stack returns the stack trace as a string.
// It returns the error string directly if the `err` does not implement the Stacker interface.
func Stack(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(errors.Stacker); ok {
		return e.Stack()
	}
	return err.Error()
}

// Current creates and returns the current error level.
// It returns nil if the current level error is nil.
func Current(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(errors.Currenter); ok {
		return e.Current()
	}
	return err
}

// Unwrap returns the next level error.
// It returns nil if the current level error or the next level error is nil.
func Unwrap(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(errors.Unwrapper); ok {
		return e.Unwrap()
	}
	return nil
}

// HasStack checks and reports whether `err` implements the `Stacker` interface.
func HasStack(err error) bool {
	_, ok := err.(errors.Stacker)
	return ok
}

// Equal reports whether the current error `err` equals the error `target`.
// Note that in the default comparison logic for `Error`,
// the errors are considered identical if both the `code` and `text` are the same.
func Equal(err, target error) bool {
	if err == target {
		return true
	}
	if e, ok := err.(errors.Equaler); ok {
		return e.Equal(target)
	}
	if e, ok := target.(errors.Equaler); ok {
		return e.Equal(err)
	}
	return false
}

// Is reports whether the current error `err` has the error `target` in its error chain.
// This is an implementation of the stdlib `errors.Is` introduced in Go version 1.17.
func Is(err, target error) bool {
	if e, ok := err.(errors.IsComparer); ok {
		return e.Is(target)
	}
	return false
}

// ContainsError is an alias of Is, with more easily understandable semantics.
func ContainsError(err, target error) bool {
	return Is(err, target)
}

// Callers returns the stack callers.
// Note that it only retrieves the caller memory address array, not the caller information.
func Callers(skip ...int) CallStack {
	var (
		pcs [maxStackDepth]uintptr
		n   = 3
	)
	if len(skip) > 0 {
		n += skip[0]
	}
	return pcs[:runtime.Callers(n, pcs[:])]
}
