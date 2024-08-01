// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// ELF implements the classic ELF hash algorithm for 32 bits.
// This algorithm shifts the hash left by 4 bits and adds the byte value,
// then applies a bitmask to reduce collisions.
func ELF(str []byte) uint32 {
	var hash uint32
	for _, b := range str {
		hash = (hash << 4) + uint32(b)      // Shift left and add byte value
		if x := hash & 0xF0000000; x != 0 { // Apply bitmask to reduce collisions
			hash ^= x >> 24
			hash &= ^x // Reset the masked bits
		}
	}
	return hash
}

// ELF64 implements the classic ELF hash algorithm for 64 bits.
// This algorithm shifts the hash left by 4 bits and adds the byte value,
// then applies a bitmask to reduce collisions.
func ELF64(str []byte) uint64 {
	var hash uint64
	for _, b := range str {
		hash = (hash << 4) + uint64(b)              // Shift left and add byte value
		if x := hash & 0xF000000000000000; x != 0 { // Apply bitmask to reduce collisions
			hash ^= x >> 56
			hash &= ^x // Reset the masked bits
		}
	}
	return hash
}
