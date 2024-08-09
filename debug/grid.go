// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package debug

import (
	"regexp"
	"runtime"
	"strconv"
)

var (
	// gridRegex is the regular expression object for parsing goroutine id from stack information.
	gridRegex = regexp.MustCompile(`^\w+\s+(\d+)\s+`)
)

// GoroutineId retrieves and returns the current goroutine id from stack information.
// Be very aware that, it is with low performance as it uses runtime.Stack function.
// It is commonly used for debugging purpose.
func GoroutineId() int {
	buf := make([]byte, 64) // Increase buffer size to ensure capturing full goroutine id line
	n := runtime.Stack(buf, false)
	if n == 0 {
		return -1
	}
	buf = buf[:n]
	match := gridRegex.FindSubmatch(buf)
	if len(match) < 2 {
		return -1 // Return an invalid ID if matching fails
	}
	id, err := strconv.Atoi(string(match[1]))
	if err != nil {
		return -1 // Return an invalid ID if conversion fails
	}
	return id
}
