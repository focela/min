// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// HashELF32 implements the classic ELF hash algorithm for 32 bits.
func HashELF32(str []byte) uint32 {
	var hash uint32
	var x uint32

	for _, b := range str {
		hash = (hash << 4) + uint32(b)
		if x = hash & 0xF0000000; x != 0 {
			hash ^= x >> 24
			hash &= 0xFFFFFFFF
		}
	}
	return hash
}

// HashELF64 implements the classic ELF hash algorithm for 64 bits.
func HashELF64(str []byte) uint64 {
	var hash uint64
	var x uint64

	for _, b := range str {
		hash = (hash << 4) + uint64(b)
		if x = hash & 0xF000000000000000; x != 0 {
			hash ^= x >> 24
			hash &= 0xFFFFFFFFFFFFFFFF
		}
	}
	return hash
}
