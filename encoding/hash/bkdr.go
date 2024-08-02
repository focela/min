// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// BKDR implements the classic BKDR hash algorithm for 32 bits.
func BKDR(str []byte) uint32 {
	var (
		seed uint32 = 131 // Seed value: 31, 131, 1313, 13131, 131313 etc.
		hash uint32
	)
	for _, c := range str {
		hash = hash*seed + uint32(c)
	}
	return hash
}

// BKDR64 implements the classic BKDR hash algorithm for 64 bits.
func BKDR64(str []byte) uint64 {
	var (
		seed uint64 = 131 // Seed value: 31, 131, 1313, 13131, 131313 etc.
		hash uint64
	)
	for _, c := range str {
		hash = hash*seed + uint64(c)
	}
	return hash
}
