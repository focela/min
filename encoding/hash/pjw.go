// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

const (
	bitsInUnsignedInt32 uint32 = 32 // 4 * 8
	threeQuarters32            = (bitsInUnsignedInt32 * 3) / 4
	oneEighth32                = bitsInUnsignedInt32 / 8
	highBits32                 = (uint32(0xFFFFFFFF)) << (bitsInUnsignedInt32 - oneEighth32)

	bitsInUnsignedInt64 uint64 = 64 // 8 * 8
	threeQuarters64            = (bitsInUnsignedInt64 * 3) / 4
	oneEighth64                = bitsInUnsignedInt64 / 8
	highBits64                 = (uint64(0xFFFFFFFFFFFFFFFF)) << (bitsInUnsignedInt64 - oneEighth64)
)

// PJW implements the classic PJW hash algorithm for 32-bit values.
func PJW(str []byte) uint32 {
	var hash, test uint32
	for i := 0; i < len(str); i++ {
		hash = (hash << oneEighth32) + uint32(str[i])
		if test = hash & highBits32; test != 0 {
			hash = (hash ^ (test >> threeQuarters32)) & (^highBits32)
		}
	}
	return hash
}

// PJW64 implements the classic PJW hash algorithm for 64-bit values.
func PJW64(str []byte) uint64 {
	var hash, test uint64
	for i := 0; i < len(str); i++ {
		hash = (hash << oneEighth64) + uint64(str[i])
		if test = hash & highBits64; test != 0 {
			hash = (hash ^ (test >> threeQuarters64)) & (^highBits64)
		}
	}
	return hash
}
