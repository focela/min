// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// HashPJW32 implements the classic PJW hash algorithm for 32 bits.
func HashPJW32(str []byte) uint32 {
	const (
		BitsInUnsignedInt uint32 = 32 // 4 * 8
		ThreeQuarters            = (BitsInUnsignedInt * 3) / 4
		OneEighth                = BitsInUnsignedInt / 8
		HighBits          uint32 = 0xF0000000
	)
	var hash uint32
	var test uint32

	for _, b := range str {
		hash = (hash << OneEighth) + uint32(b)
		if test = hash & HighBits; test != 0 {
			hash = (hash ^ (test >> ThreeQuarters)) &^ HighBits
		}
	}
	return hash
}

// HashPJW64 implements the classic PJW hash algorithm for 64 bits.
func HashPJW64(str []byte) uint64 {
	const (
		BitsInUnsignedInt uint64 = 64 // 8 * 8
		ThreeQuarters            = (BitsInUnsignedInt * 3) / 4
		OneEighth                = BitsInUnsignedInt / 8
		HighBits          uint64 = 0xF000000000000000
	)
	var hash uint64
	var test uint64

	for _, b := range str {
		hash = (hash << OneEighth) + uint64(b)
		if test = hash & HighBits; test != 0 {
			hash = (hash ^ (test >> ThreeQuarters)) &^ HighBits
		}
	}
	return hash
}
