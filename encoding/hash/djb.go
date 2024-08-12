// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// DJB implements the classic DJB hash algorithm for 32-bit integers.
// This algorithm is simple and efficient for string hashing.
func DJB(str []byte) uint32 {
	var hash uint32 = 5381
	for _, b := range str {
		hash += (hash << 5) + uint32(b)
	}
	return hash
}

// DJB64 implements the classic DJB hash algorithm for 64-bit integers.
// This version extends the DJB algorithm to operate on 64-bit values.
func DJB64(str []byte) uint64 {
	var hash uint64 = 5381
	for _, b := range str {
		hash += (hash << 5) + uint64(b)
	}
	return hash
}
