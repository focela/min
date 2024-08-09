// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// HashBKDR32 implements the classic BKDR hash algorithm for 32 bits.
func HashBKDR32(str []byte) uint32 {
	const seed uint32 = 131 // 31 131 1313 13131 131313 etc.
	var hash uint32 = 0

	for _, b := range str {
		hash = hash*seed + uint32(b)
	}
	return hash
}

// HashBKDR64 implements the classic BKDR hash algorithm for 64 bits.
func HashBKDR64(str []byte) uint64 {
	const seed uint64 = 131 // 31 131 1313 13131 131313 etc.
	var hash uint64 = 0

	for _, b := range str {
		hash = hash*seed + uint64(b)
	}
	return hash
}
