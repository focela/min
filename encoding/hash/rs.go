// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

const (
	initialMultiplier32 uint32 = 63689
	primeFactor32       uint32 = 378551

	initialMultiplier64 uint64 = 63689
	primeFactor64       uint64 = 378551
)

// ComputeRSHash implements the classic RS hash algorithm, which produces a 32-bit hash value
// by iteratively combining character codes with a prime-based multiplier.
func ComputeRSHash(str []byte) uint32 {
	var (
		a           = initialMultiplier32
		hash uint32 = 0
	)
	length := len(str)
	for i := 0; i < length; i++ {
		hash = hash*a + uint32(str[i])
		a *= primeFactor32
	}
	return hash
}

// ComputeRSHash64 implements the classic RS hash algorithm, which produces a 64-bit hash value
// by iteratively combining character codes with a prime-based multiplier.
func ComputeRSHash64(str []byte) uint64 {
	var (
		a           = initialMultiplier64
		hash uint64 = 0
	)
	length := len(str)
	for i := 0; i < length; i++ {
		hash = hash*a + uint64(str[i])
		a *= primeFactor64
	}
	return hash
}
