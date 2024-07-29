// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rtime

import "time"

// wrapper is a wrapper for stdlib struct time.Time.
// It's used for overwriting some functions of time.Time, for example: String.
type wrapper struct {
	time.Time
}

// String overwrites the String function of time.Time.
func (t wrapper) String() string {
	if t.IsZero() {
		return ""
	}
	if t.Year() == 0 {
		// Only time.
		return t.Format("15:04:05")
	}
	return t.Format("2006-01-02 15:04:05")
}
