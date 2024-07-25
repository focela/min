// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcutil

import "github.com/focela/ratcatcher/internal/reflection"

type (
	OriginValueAndKindOutput = reflection.OriginValueAndKindOutput
	OriginTypeAndKindOutput  = reflection.OriginTypeAndKindOutput
)

// OriginValueAndKind retrieves and returns the original reflect value and kind.
func OriginValueAndKind(value interface{}) (out OriginValueAndKindOutput) {
	return reflection.OriginValueAndKind(value)
}

// OriginTypeAndKind retrieves and returns the original reflect type and kind.
func OriginTypeAndKind(value interface{}) (out OriginTypeAndKindOutput) {
	return reflection.OriginTypeAndKind(value)
}
