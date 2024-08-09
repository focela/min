// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package times

import (
	"time"
)

// wrapper is a wrapper for the standard time.Time struct.
// It overwrites some methods of time.Time, for example: String.
type wrapper struct {
	time.Time
}

// String overwrites the String function of time.Time.
// It returns a formatted time string based on whether the year is present or not.
func (t wrapper) String() string {
	if t.IsZero() {
		return ""
	}
	if t.Year() == 0 {
		return t.Format("15:04:05") // Time only
	}
	return t.Format("2006-01-02 15:04:05") // Full date and time
}
