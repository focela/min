// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rctrace

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// Tracer warps trace.Tracer for compatibility and extension.
type Tracer struct {
	trace.Tracer
}

// NewTracer Tracer is a short function for retrieving Tracer.
func NewTracer(name ...string) *Tracer {
	tracerName := ""
	if len(name) > 0 {
		tracerName = name[0]
	}
	return &Tracer{
		Tracer: otel.Tracer(tracerName),
	}
}
