// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// HashAP32 implements the classic AP hash algorithm for 32 bits.
func HashAP32(str []byte) uint32 {
	var hash uint32 = 0xABCDEF // Initialize with a non-zero value
	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint32(str[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint32(str[i]) ^ (hash >> 5)) + 1
		}
	}
	return hash
}

// HashAP64 implements the classic AP hash algorithm for 64 bits.
func HashAP64(str []byte) uint64 {
	var hash uint64 = 0x123456789ABCDEF // Initialize with a non-zero value
	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint64(str[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint64(str[i]) ^ (hash >> 5)) + 1
		}
	}
	return hash
}
