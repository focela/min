// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// PJW implements the classic PJW hash algorithm for 32 bits.
// This algorithm is based on shifting bits and applying bitwise operations
// to compute the hash value.
func PJW(str []byte) uint32 {
	var (
		BitsInUnsignedInt uint32 = 32 // Number of bits in an unsigned integer
		ThreeQuarters            = (BitsInUnsignedInt * 3) / 4
		OneEighth                = BitsInUnsignedInt / 8
		HighBits          uint32 = (0xFFFFFFFF) << (BitsInUnsignedInt - OneEighth)
		hash              uint32
		test              uint32
	)
	for _, b := range str {
		hash = (hash << OneEighth) + uint32(b)
		if test = hash & HighBits; test != 0 {
			hash = (hash ^ (test >> ThreeQuarters)) & (^HighBits + 1)
		}
	}
	return hash
}

// PJW64 implements the classic PJW hash algorithm for 64 bits.
// This algorithm is based on shifting bits and applying bitwise operations
// to compute the hash value.
func PJW64(str []byte) uint64 {
	var (
		BitsInUnsignedInt uint64 = 64 // Number of bits in an unsigned integer for 64-bit version
		ThreeQuarters            = (BitsInUnsignedInt * 3) / 4
		OneEighth                = BitsInUnsignedInt / 8
		HighBits          uint64 = (0xFFFFFFFFFFFFFFFF) << (BitsInUnsignedInt - OneEighth)
		hash              uint64
		test              uint64
	)
	for _, b := range str {
		hash = (hash << OneEighth) + uint64(b)
		if test = hash & HighBits; test != 0 {
			hash = (hash ^ (test >> ThreeQuarters)) & (^HighBits + 1)
		}
	}
	return hash
}
