// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// HashDJB32 implements the classic DJB hash algorithm for 32 bits.
func HashDJB32(str []byte) uint32 {
	const seed uint32 = 5381
	var hash = seed

	for _, b := range str {
		hash += (hash << 5) + uint32(b)
	}
	return hash
}

// HashDJB64 implements the classic DJB hash algorithm for 64 bits.
func HashDJB64(str []byte) uint64 {
	const seed uint64 = 5381
	var hash = seed

	for _, b := range str {
		hash += (hash << 5) + uint64(b)
	}
	return hash
}
