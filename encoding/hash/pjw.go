// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// PJW implements the classic PJW hash algorithm for 32 bits.
func PJW(str []byte) uint32 {
	var (
		BitsInUnsignedInt uint32 = 32
		ThreeQuarters            = (BitsInUnsignedInt * 3) / 4
		OneEighth                = BitsInUnsignedInt / 8
		HighBits          uint32 = 0xFFFFFFFF << (BitsInUnsignedInt - OneEighth)
		hash              uint32
		test              uint32
	)
	for _, c := range str {
		hash = (hash << OneEighth) + uint32(c)
		if test = hash & HighBits; test != 0 {
			hash = (hash ^ (test >> ThreeQuarters)) & ^HighBits
		}
	}
	return hash
}

// PJW64 implements the classic PJW hash algorithm for 64 bits.
func PJW64(str []byte) uint64 {
	var (
		BitsInUnsignedInt uint64 = 64
		ThreeQuarters            = (BitsInUnsignedInt * 3) / 4
		OneEighth                = BitsInUnsignedInt / 8
		HighBits          uint64 = 0xFFFFFFFFFFFFFFFF << (BitsInUnsignedInt - OneEighth)
		hash              uint64
		test              uint64
	)
	for _, c := range str {
		hash = (hash << OneEighth) + uint64(c)
		if test = hash & HighBits; test != 0 {
			hash = (hash ^ (test >> ThreeQuarters)) & ^HighBits
		}
	}
	return hash
}
