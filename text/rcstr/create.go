// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcstr

import "strings"

// Repeat returns a new string consisting of multiplier copies of the string input.
//
// Example:
// Repeat("a", 3) -> "aaa"
func Repeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}
