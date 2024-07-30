// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package hash

// DJB implements the classic DJB hash algorithm for 32 bits.
func DJB(str []byte) uint32 {
	var hash uint32 = 5381
	for i := 0; i < len(str); i++ {
		hash = (hash << 5) + hash + uint32(str[i])
	}
	return hash
}

// DJB64 implements the classic DJB hash algorithm for 64 bits.
func DJB64(str []byte) uint64 {
	var hash uint64 = 5381
	for i := 0; i < len(str); i++ {
		hash = (hash << 5) + hash + uint64(str[i])
	}
	return hash
}
