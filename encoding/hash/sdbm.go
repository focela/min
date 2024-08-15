// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

// SDBM implements the classic SDBM hash algorithm for 32 bits.
// SDBM is a simple hash function that uses bitwise shifts
// and subtraction to generate a hash value.
func SDBM(str []byte) uint32 {
	var hash uint32
	for i := 0; i < len(str); i++ {
		// equivalent to: hash = 65599*hash + uint32(str[i]);
		hash = uint32(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

// SDBM64 implements the classic SDBM hash algorithm for 64 bits.
// This is a 64-bit version of the SDBM hash algorithm, suitable for larger data sets.
func SDBM64(str []byte) uint64 {
	var hash uint64
	for i := 0; i < len(str); i++ {
		// equivalent to: hash = 65599*hash + uint32(str[i])
		hash = uint64(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}
