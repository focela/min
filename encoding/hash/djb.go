// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// DJB implements the classic DJB hash algorithm for 32 bits.
// It uses a starting hash value of 5381 and a shift-add approach for hashing.
func DJB(str []byte) uint32 {
	var hash uint32 = 5381
	for _, b := range str { // Iterate over each byte in the slice
		hash = (hash << 5) + hash + uint32(b) // Update the hash value
	}
	return hash
}

// DJB64 implements the classic DJB hash algorithm for 64 bits.
// It uses a starting hash value of 5381 and a shift-add approach for hashing.
func DJB64(str []byte) uint64 {
	var hash uint64 = 5381
	for _, b := range str { // Iterate over each byte in the slice
		hash = (hash << 5) + hash + uint64(b) // Update the hash value
	}
	return hash
}
