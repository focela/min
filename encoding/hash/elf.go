// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

// ELF calculates a 32-bit hash value using the classic ELF hash algorithm.
func ELF(str []byte) uint32 {
	var (
		hash uint32
		x    uint32
	)
	for i := 0; i < len(str); i++ {
		hash = (hash << 4) + uint32(str[i])
		if x = hash & 0xF0000000; x != 0 {
			hash ^= x >> 24
			hash &= ^x // Chỉnh sửa logic tại đây
		}
	}
	return hash
}

// ELF64 calculates a 64-bit hash value using the classic ELF hash algorithm.
func ELF64(str []byte) uint64 {
	var (
		hash uint64
		x    uint64
	)
	for i := 0; i < len(str); i++ {
		hash = (hash << 4) + uint64(str[i])
		if x = hash & 0xF000000000000000; x != 0 {
			hash ^= x >> 24
			hash &= ^x // Chỉnh sửa logic tại đây
		}
	}
	return hash
}
