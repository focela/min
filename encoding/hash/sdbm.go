// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// SDBM implements the classic SDBM hash algorithm for 32-bit integers.
// This algorithm is simple and effective for quick hash calculations.
func SDBM(str []byte) uint32 {
	var hash uint32
	for _, b := range str {
		// equivalent to: hash = 65599*hash + uint32(b);
		hash = uint32(b) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

// SDBM64 implements the classic SDBM hash algorithm for 64-bit integers.
// This version extends the SDBM algorithm to operate on 64-bit values.
func SDBM64(str []byte) uint64 {
	var hash uint64
	for _, b := range str {
		// equivalent to: hash = 65599*hash + uint64(b)
		hash = uint64(b) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}
