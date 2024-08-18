// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

// ComputeAPHash implements the classic AP hash algorithm, which produces a 32-bit hash value.
func ComputeAPHash(str []byte) uint32 {
	var hash uint32
	length := len(str)
	for i := 0; i < length; i++ {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint32(str[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint32(str[i]) ^ (hash >> 5)) + 1
		}
	}
	return hash
}

// ComputeAPHash64 implements the classic AP hash algorithm, which produces a 64-bit hash value.
func ComputeAPHash64(str []byte) uint64 {
	var hash uint64
	length := len(str)
	for i := 0; i < length; i++ {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint64(str[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint64(str[i]) ^ (hash >> 5)) + 1
		}
	}
	return hash
}
