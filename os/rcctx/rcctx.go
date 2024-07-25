// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package rcctx wraps context.Context and provides extra context features.
package rcctx

import (
	"context"
	"os"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/focela/ratcatcher/net/rctrace"
)

type (
	Ctx    = context.Context // Ctx is short name alias for context.Context.
	StrKey string            // StrKey is a type for warps basic type string as context key.
)

var (
	// initCtx is the context initialized from process environment.
	initCtx context.Context
)

func init() {
	// All environment key-value pairs.
	m := make(map[string]string)
	i := 0
	for _, s := range os.Environ() {
		i = strings.IndexByte(s, '=')
		if i == -1 {
			continue
		}
		m[s[0:i]] = s[i+1:]
	}
	// OpenTelemetry from environments.
	initCtx = otel.GetTextMapPropagator().Extract(
		context.Background(),
		propagation.MapCarrier(m),
	)
	initCtx = WithCtx(initCtx)
}

// New creates and returns a context which contains context id.
func New() context.Context {
	return WithCtx(context.Background())
}

// WithCtx creates and returns a context containing context id upon given parent context `ctx`.
func WithCtx(ctx context.Context) context.Context {
	if CtxId(ctx) != "" {
		return ctx
	}
	var span *rctrace.Span
	ctx, span = rctrace.NewSpan(ctx, "rcctx.WithCtx")
	defer span.End()
	return ctx
}

// CtxId retrieves and returns the context id from context.
func CtxId(ctx context.Context) string {
	return rctrace.GetTraceID(ctx)
}

// SetInitCtx sets custom initialization context.
// Note that this function cannot be called in multiple goroutines.
func SetInitCtx(ctx context.Context) {
	initCtx = ctx
}

// GetInitCtx returns the initialization context.
// Initialization context is used in `main` or `init` functions.
func GetInitCtx() context.Context {
	return initCtx
}
