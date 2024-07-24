// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcstr

import "strings"

// Contains reports whether `substr` is within `str`, case-sensitively.
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// ContainsI reports whether substr is within str, case-insensitively.
func ContainsI(str, substr string) bool {
	return PosI(str, substr) != -1
}

// ContainsAny reports whether any Unicode code points in `chars` are within `s`.
func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}
