// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// PJW implements the classic PJW hash algorithm for 32-bit integers.
// This algorithm is designed by Peter J. Weinberger and is known for its effectiveness in hashing strings.
func PJW(str []byte) uint32 {
	var (
		bitsInUint    uint32 = 32
		threeQuarters        = (bitsInUint * 3) / 4
		oneEighth            = bitsInUint / 8
		highBits      uint32 = 0xF0000000
		hash          uint32
		test          uint32
	)
	for _, b := range str {
		hash = (hash << oneEighth) + uint32(b)
		if test = hash & highBits; test != 0 {
			hash = (hash ^ (test >> threeQuarters)) & (^highBits + 1)
		}
	}
	return hash
}

// PJW64 implements the classic PJW hash algorithm for 64-bit integers.
// This version extends the PJW algorithm to operate on 64-bit values.
func PJW64(str []byte) uint64 {
	var (
		bitsInUint    uint64 = 64
		threeQuarters        = (bitsInUint * 3) / 4
		oneEighth            = bitsInUint / 8
		highBits      uint64 = 0xF000000000000000
		hash          uint64
		test          uint64
	)
	for _, b := range str {
		hash = (hash << oneEighth) + uint64(b)
		if test = hash & highBits; test != 0 {
			hash = (hash ^ (test >> threeQuarters)) & (^highBits + 1)
		}
	}
	return hash
}
