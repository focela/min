// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

// RS implements the classic RS hash algorithm for 32 bits.
// The RS hash algorithm is a simple but effective hash function.
// It uses two multipliers `a` and `b` to generate the hash value.
func RS(str []byte) uint32 {
	var (
		b    uint32 = 378551
		a    uint32 = 63689
		hash uint32 = 0
	)
	for i := 0; i < len(str); i++ {
		hash = hash*a + uint32(str[i])
		a *= b
	}
	return hash
}

// RS64 implements the classic RS hash algorithm for 64 bits.
// This is a 64-bit version of the RS hash algorithm, useful for larger data sets.
func RS64(str []byte) uint64 {
	var (
		b    uint64 = 378551
		a    uint64 = 63689
		hash uint64 = 0
	)
	for i := 0; i < len(str); i++ {
		hash = hash*a + uint64(str[i])
		a *= b
	}
	return hash
}
