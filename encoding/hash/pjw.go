// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

const (
	BitsInUnsignedInt32 uint32 = 32 // 4 * 8
	ThreeQuarters32            = (BitsInUnsignedInt32 * 3) / 4
	OneEighth32                = BitsInUnsignedInt32 / 8
	HighBits32          uint32 = 0xF0000000

	BitsInUnsignedInt64 uint64 = 64 // 8 * 8
	ThreeQuarters64            = (BitsInUnsignedInt64 * 3) / 4
	OneEighth64                = BitsInUnsignedInt64 / 8
	HighBits64          uint64 = 0xF000000000000000
)

// ComputePJWHash implements the classic PJW hash algorithm, which produces a 32-bit hash value
// using shift and mask operations to distribute bits.
func ComputePJWHash(str []byte) uint32 {
	var hash, test uint32
	length := len(str)
	for i := 0; i < length; i++ {
		hash = (hash << OneEighth32) + uint32(str[i])
		if test = hash & HighBits32; test != 0 {
			hash = (hash ^ (test >> ThreeQuarters32)) & (^HighBits32)
		}
	}
	return hash
}

// ComputePJWHash64 implements the classic PJW hash algorithm, which produces a 64-bit hash value
// using shift and mask operations to distribute bits.
func ComputePJWHash64(str []byte) uint64 {
	var hash, test uint64
	length := len(str)
	for i := 0; i < length; i++ {
		hash = (hash << OneEighth64) + uint64(str[i])
		if test = hash & HighBits64; test != 0 {
			hash = (hash ^ (test >> ThreeQuarters64)) & (^HighBits64)
		}
	}
	return hash
}
