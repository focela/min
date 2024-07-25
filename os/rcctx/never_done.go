// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcctx

import (
	"context"
	"time"
)

// neverDoneCtx never done.
type neverDoneCtx struct {
	context.Context
}

// Done forbids the context done from parent context.
func (*neverDoneCtx) Done() <-chan struct{} {
	return nil
}

// Deadline forbids the context deadline from parent context.
func (*neverDoneCtx) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

// Err forbids the context done from parent context.
func (c *neverDoneCtx) Err() error {
	return nil
}

// NeverDone wraps and returns a new context object that will be never done,
// which forbids the context manually done, to make the context can be propagated
// to asynchronous goroutines.
//
// Note that, it does not affect the closing (canceling) of the parent context,
// as it is a wrapper for its parent, which only affects the next context handling.
func NeverDone(ctx context.Context) context.Context {
	return &neverDoneCtx{ctx}
}
