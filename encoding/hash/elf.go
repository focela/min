// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

// ComputeELFHash implements the classic ELF hash algorithm, which produces a 32-bit hash value
// using a shift-and-add method with bitwise manipulation for efficiency.
func ComputeELFHash(str []byte) uint32 {
	var (
		hash uint32
		x    uint32
	)
	length := len(str)
	for i := 0; i < length; i++ {
		hash = (hash << 4) + uint32(str[i])
		if x = hash & 0xF0000000; x != 0 {
			hash ^= x >> 24
			hash &= ^(x)
		}
	}
	return hash
}

// ComputeELFHash64 implements the classic ELF hash algorithm, which produces a 64-bit hash value
// using a shift-and-add method with bitwise manipulation for efficiency.
func ComputeELFHash64(str []byte) uint64 {
	var (
		hash uint64
		x    uint64
	)
	length := len(str)
	for i := 0; i < length; i++ {
		hash = (hash << 4) + uint64(str[i])
		if x = hash & 0xF000000000000000; x != 0 {
			hash ^= x >> 24
			hash &= ^(x)
		}
	}
	return hash
}
