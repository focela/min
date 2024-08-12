// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// BKDR implements the classic BKDR hash algorithm for 32-bit integers.
// This algorithm is known for its simplicity and efficiency in handling strings.
func BKDR(str []byte) uint32 {
	var (
		seed uint32 = 131 // 31, 131, 1313, 13131, 131313 etc.
		hash uint32 = 0
	)
	for _, b := range str {
		hash = hash*seed + uint32(b)
	}
	return hash
}

// BKDR64 implements the classic BKDR hash algorithm for 64-bit integers.
// This version extends the BKDR algorithm to operate on 64-bit values.
func BKDR64(str []byte) uint64 {
	var (
		seed uint64 = 131 // 31, 131, 1313, 13131, 131313 etc.
		hash uint64 = 0
	)
	for _, b := range str {
		hash = hash*seed + uint64(b)
	}
	return hash
}
