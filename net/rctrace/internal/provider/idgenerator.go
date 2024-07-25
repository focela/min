// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package provider

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/focela/ratcatcher/internal/tracing"
)

// IDGenerator is a trace ID generator.
type IDGenerator struct{}

// NewIDGenerator returns a new IDGenerator.
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

// NewIDs creates and returns a new trace and span ID.
func (id *IDGenerator) NewIDs(ctx context.Context) (traceID trace.TraceID, spanID trace.SpanID) {
	return tracing.NewIDs()
}

// NewSpanID returns an ID for a new span in the trace with traceID.
func (id *IDGenerator) NewSpanID(ctx context.Context, traceID trace.TraceID) (spanID trace.SpanID) {
	return tracing.NewSpanID()
}
