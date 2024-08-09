// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// HashSDBM32 implements the classic SDBM hash algorithm for 32 bits.
func HashSDBM32(str []byte) uint32 {
	var hash uint32

	for _, b := range str {
		// equivalent to: hash = 65599*hash + uint32(b)
		hash = uint32(b) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

// HashSDBM64 implements the classic SDBM hash algorithm for 64 bits.
func HashSDBM64(str []byte) uint64 {
	var hash uint64

	for _, b := range str {
		// equivalent to: hash = 65599*hash + uint32(b)
		hash = uint64(b) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}
