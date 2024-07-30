// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package hash

// PJW implements the classic PJW hash algorithm for 32 bits.
func PJW(str []byte) uint32 {
	var (
		BitsInUnsignedInt uint32 = 32 // 4 * 8
		ThreeQuarters            = (BitsInUnsignedInt * 3) / 4
		OneEighth                = BitsInUnsignedInt / 8
		HighBits          uint32 = 0xF0000000
		hash              uint32
		test              uint32
	)
	for i := 0; i < len(str); i++ {
		hash = (hash << OneEighth) + uint32(str[i])
		if test = hash & HighBits; test != 0 {
			hash = (hash ^ (test >> ThreeQuarters)) &^ HighBits
		}
	}
	return hash
}

// PJW64 implements the classic PJW hash algorithm for 64 bits.
func PJW64(str []byte) uint64 {
	var (
		BitsInUnsignedInt uint64 = 64 // 8 * 8
		ThreeQuarters            = (BitsInUnsignedInt * 3) / 4
		OneEighth                = BitsInUnsignedInt / 8
		HighBits          uint64 = 0xF000000000000000
		hash              uint64
		test              uint64
	)
	for i := 0; i < len(str); i++ {
		hash = (hash << OneEighth) + uint64(str[i])
		if test = hash & HighBits; test != 0 {
			hash = (hash ^ (test >> ThreeQuarters)) &^ HighBits
		}
	}
	return hash
}
