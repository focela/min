// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package minhash

// AP implements the classic AP hash algorithm for 32 bits.
func AP(str []byte) uint32 {
	var hash uint32
	for i, b := range str {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint32(b) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint32(b) ^ (hash >> 5)) + 1
		}
	}
	return hash
}

// AP64 implements the classic AP hash algorithm for 64 bits.
func AP64(str []byte) uint64 {
	var hash uint64
	for i, b := range str {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint64(b) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint64(b) ^ (hash >> 5)) + 1
		}
	}
	return hash
}
