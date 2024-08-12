// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// ELF implements the classic ELF hash algorithm for 32-bit integers.
// This algorithm is commonly used in Unix systems for string hashing.
func ELF(str []byte) uint32 {
	var (
		hash uint32
		mask uint32
	)
	for _, b := range str {
		hash = (hash << 4) + uint32(b)
		if mask = hash & 0xF0000000; mask != 0 {
			hash ^= mask >> 24
			hash &= ^mask + 1
		}
	}
	return hash
}

// ELF64 implements the classic ELF hash algorithm for 64-bit integers.
// This version extends the ELF algorithm to operate on 64-bit values.
func ELF64(str []byte) uint64 {
	var (
		hash uint64
		mask uint64
	)
	for _, b := range str {
		hash = (hash << 4) + uint64(b)
		if mask = hash & 0xF000000000000000; mask != 0 {
			hash ^= mask >> 24
			hash &= ^mask + 1
		}
	}
	return hash
}
