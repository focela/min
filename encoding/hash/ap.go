// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// AP implements the classic AP hash algorithm for 32 bits.
// It returns a 32-bit hash value for the given byte slice.
func AP(str []byte) uint32 {
	var hash uint32
	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 { // If the index is even
			hash ^= (hash << 7) ^ uint32(str[i]) ^ (hash >> 3)
		} else { // If the index is odd
			hash ^= ^((hash << 11) ^ uint32(str[i]) ^ (hash >> 5)) + 1
		}
	}
	return hash
}

// AP64 implements the classic AP hash algorithm for 64 bits.
// It returns a 64-bit hash value for the given byte slice.
func AP64(str []byte) uint64 {
	var hash uint64
	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 { // If the index is even
			hash ^= (hash << 7) ^ uint64(str[i]) ^ (hash >> 3)
		} else { // If the index is odd
			hash ^= ^((hash << 11) ^ uint64(str[i]) ^ (hash >> 5)) + 1
		}
	}
	return hash
}
