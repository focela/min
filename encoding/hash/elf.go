// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// ELF implements the classic ELF hash algorithm for 32 bits.
func ELF(str []byte) uint32 {
	var (
		hash uint32
		x    uint32
	)
	for _, c := range str {
		hash = (hash << 4) + uint32(c)
		if x = hash & 0xF0000000; x != 0 {
			hash ^= x >> 24
			hash &= ^x
		}
	}
	return hash
}

// ELF64 implements the classic ELF hash algorithm for 64 bits.
func ELF64(str []byte) uint64 {
	var (
		hash uint64
		x    uint64
	)
	for _, c := range str {
		hash = (hash << 4) + uint64(c)
		if x = hash & 0xF000000000000000; x != 0 {
			hash ^= x >> 56
			hash &= ^x
		}
	}
	return hash
}
