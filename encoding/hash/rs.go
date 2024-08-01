// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// RS implements the classic RS hash algorithm for 32 bits.
// The RS hash algorithm was created by Robert Sedgwicks.
func RS(str []byte) uint32 {
	var (
		b    uint32 = 378551 // A constant used in the RS algorithm
		a    uint32 = 63689  // Another constant used in the RS algorithm
		hash uint32 = 0      // Initial hash value
	)
	for _, char := range str { // Changed 'b' to 'char'
		hash = hash*a + uint32(char) // Update hash value using the RS algorithm
		a *= b                       // Update `a` using constant multiplier `b`
	}
	return hash
}

// RS64 implements the classic RS hash algorithm for 64 bits.
// The RS hash algorithm was created by Robert Sedgwicks.
func RS64(str []byte) uint64 {
	var (
		b    uint64 = 378551 // A constant used in the RS algorithm
		a    uint64 = 63689  // Another constant used in the RS algorithm
		hash uint64 = 0      // Initial hash value
	)
	for _, char := range str { // Changed 'b' to 'char'
		hash = hash*a + uint64(char) // Update hash value using the RS algorithm
		a *= b                       // Update `a` using constant multiplier `b`
	}
	return hash
}
