// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package hash

// JS implements the classic JS hash algorithm for 32 bits.
func JS(str []byte) uint32 {
	var hash uint32 = 1315423911
	for i := 0; i < len(str); i++ {
		hash ^= (hash << 5) + uint32(str[i]) + (hash >> 2)
	}
	return hash
}

// JS64 implements the classic JS hash algorithm for 64 bits.
func JS64(str []byte) uint64 {
	var hash uint64 = 1315423911
	for i := 0; i < len(str); i++ {
		hash ^= (hash << 5) + uint64(str[i]) + (hash >> 2)
	}
	return hash
}
