// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// HashJS32 implements the classic JS hash algorithm for 32 bits.
func HashJS32(str []byte) uint32 {
	const initialHash uint32 = 1315423911
	var hash = initialHash

	for _, b := range str {
		hash ^= (hash << 5) + uint32(b) + (hash >> 2)
	}
	return hash
}

// HashJS64 implements the classic JS hash algorithm for 64 bits.
func HashJS64(str []byte) uint64 {
	const initialHash uint64 = 1315423911
	var hash = initialHash

	for _, b := range str {
		hash ^= (hash << 5) + uint64(b) + (hash >> 2)
	}
	return hash
}
