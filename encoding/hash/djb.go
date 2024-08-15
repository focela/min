// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

// DJB calculates a 32-bit hash value using the classic DJB hash algorithm.
func DJB(str []byte) uint32 {
	var hash uint32 = 5381
	for _, b := range str {
		hash += (hash << 5) + uint32(b)
	}
	return hash
}

// DJB64 calculates a 64-bit hash value using the classic DJB hash algorithm.
func DJB64(str []byte) uint64 {
	var hash uint64 = 5381
	for _, b := range str {
		hash += (hash << 5) + uint64(b)
	}
	return hash
}
