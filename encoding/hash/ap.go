// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// AP implements the classic AP hash algorithm for 32-bit integers.
// This algorithm is known for its simplicity and effectiveness in specific contexts.
func AP(str []byte) uint32 {
	var hash uint32
	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint32(str[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint32(str[i]) ^ (hash >> 5)) + 1
		}
	}
	return hash
}

// AP64 implements the classic AP hash algorithm for 64-bit integers.
// This version extends the AP algorithm to operate on 64-bit values.
func AP64(str []byte) uint64 {
	var hash uint64
	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint64(str[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint64(str[i]) ^ (hash >> 5)) + 1
		}
	}
	return hash
}
