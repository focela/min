// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package tracing provides some utility functions for tracing functionality.
package tracing

import (
	"math"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/focela/ratcatcher/container/rctype"
	"github.com/focela/ratcatcher/encoding/rcbinary"
	"github.com/focela/ratcatcher/utils/rcrand"
)

var (
	randomInitSequence = int32(rcrand.Intn(math.MaxInt32))
	sequence           = rctype.NewInt32(randomInitSequence)
)

// NewIDs creates and returns a new trace and span ID.
func NewIDs() (traceID trace.TraceID, spanID trace.SpanID) {
	return NewTraceID(), NewSpanID()
}

// NewTraceID creates and returns a trace ID.
func NewTraceID() (traceID trace.TraceID) {
	var (
		timestampNanoBytes = rcbinary.EncodeInt64(time.Now().UnixNano())
		sequenceBytes      = rcbinary.EncodeInt32(sequence.Add(1))
		randomBytes        = rcrand.B(4)
	)
	copy(traceID[:], timestampNanoBytes)
	copy(traceID[8:], sequenceBytes)
	copy(traceID[12:], randomBytes)
	return
}

// NewSpanID creates and returns a span ID.
func NewSpanID() (spanID trace.SpanID) {
	copy(spanID[:], rcbinary.EncodeInt64(time.Now().UnixNano()/1e3))
	copy(spanID[4:], rcrand.B(4))
	return
}
