// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package provider

import sdkTrace "go.opentelemetry.io/otel/sdk/trace"

type TracerProvider struct {
	*sdkTrace.TracerProvider
}

// New returns a new and configured TracerProvider, which has no SpanProcessor.
//
// In default the returned TracerProvider is configured with:
// - a ParentBased(AlwaysSample) Sampler;
// - a unix nano timestamp and random umber based IDGenerator;
// - the resource.Default() Resource;
// - the default SpanLimits.
//
// The passed opts are used to override these default values and configure the
// returned TracerProvider appropriately.
func New() *TracerProvider {
	return &TracerProvider{
		TracerProvider: sdkTrace.NewTracerProvider(
			sdkTrace.WithIDGenerator(NewIDGenerator()),
		),
	}
}
