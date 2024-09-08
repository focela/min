// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package minerror

import (
	"runtime"
)

// stack represents a stack of program counters.
type stack []uintptr

const (
	// maxStackDepth marks the maximum stack depth for error back traces.
	maxStackDepth = 64
)

// Cause returns the root cause error of `err`.
func Cause(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(CauseRetriever); ok {
		return e.Cause()
	}
	if e, ok := err.(Unwrapper); ok {
		return Cause(e.Unwrap())
	}
	return err
}

// Stack returns the stack callers as a string.
// It returns the error string directly if the `err` does not support stack tracing.
func Stack(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(StackTracer); ok {
		return e.Stack()
	}
	return err.Error()
}

// Current creates and returns the current level error.
// It returns nil if the current level error is nil.
func Current(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(CurrentRetriever); ok {
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
	if e, ok := err.(Unwrapper); ok {
		return e.Unwrap()
	}
	return nil
}

// HasStack checks and reports whether `err` implements the `minerror.StackTracer` interface.
func HasStack(err error) bool {
	_, ok := err.(StackTracer)
	return ok
}

// Equal reports whether the current error `err` equals to the error `target`.
// Note: By default, two `Error` instances are considered equal if both the `code` and `text` match.
func Equal(err, target error) bool {
	if err == target {
		return true
	}
	if e, ok := err.(EqualChecker); ok {
		return e.Equal(target)
	}
	if e, ok := target.(EqualChecker); ok {
		return e.Equal(err)
	}
	return false
}

// Is reports whether the current error `err` has error `target` in its chaining errors.
// This is an implementation of the standard library's errors.Is from Go version 1.17.
func Is(err, target error) bool {
	if e, ok := err.(IsChecker); ok {
		return e.Is(target)
	}
	return false
}

// callers returns the program counters (addresses) for the current stack.
// It does not include detailed caller information such as file names or line numbers.
func callers(skip ...int) stack {
	var (
		pcs [maxStackDepth]uintptr
		n   = 3
	)
	if len(skip) > 0 {
		n += skip[0]
	}
	return pcs[:runtime.Callers(n, pcs[:])]
}
