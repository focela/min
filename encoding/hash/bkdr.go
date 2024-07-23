// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package hash

// BKDR implements the classic BKDR hash algorithm for 32 bits.
func BKDR(str []byte) uint32 {
	var (
		seed uint32 = 131 // 31 131 1313 13131 131313 etc..
		hash uint32 = 0
	)
	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint32(str[i])
	}
	return hash
}

// BKDR64 implements the classic BKDR hash algorithm for 64 bits.
func BKDR64(str []byte) uint64 {
	var (
		seed uint64 = 131 // 31 131 1313 13131 131313 etc..
		hash uint64 = 0
	)
	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint64(str[i])
	}
	return hash
}
