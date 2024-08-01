// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// BKDR implements the classic BKDR hash algorithm for 32 bits.
// It uses a seed of 131, a commonly used prime number, for better distribution.
func BKDR(str []byte) uint32 {
	const seed uint32 = 131 // Seed for BKDR hash algorithm

	var hash uint32
	for _, b := range str { // Iterate over each byte in the slice
		hash = hash*seed + uint32(b) // Update the hash value
	}
	return hash
}

// BKDR64 implements the classic BKDR hash algorithm for 64 bits.
// It uses a seed of 131, a commonly used prime number, for better distribution.
func BKDR64(str []byte) uint64 {
	const seed uint64 = 131 // Seed for BKDR64 hash algorithm

	var hash uint64
	for _, b := range str { // Iterate over each byte in the slice
		hash = hash*seed + uint64(b) // Update the hash value
	}
	return hash
}
