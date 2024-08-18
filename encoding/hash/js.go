// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

const jsInitialHash32 uint32 = 1315423911
const jsInitialHash64 uint64 = 1315423911

// ComputeJSHash implements the classic JS hash algorithm, which produces a 32-bit hash value
// using bitwise operations for efficiency.
func ComputeJSHash(str []byte) uint32 {
	var hash = jsInitialHash32
	length := len(str)
	for i := 0; i < length; i++ {
		hash ^= (hash << 5) + uint32(str[i]) + (hash >> 2)
	}
	return hash
}

// ComputeJSHash64 implements the classic JS hash algorithm, which produces a 64-bit hash value
// using bitwise operations for efficiency.
func ComputeJSHash64(str []byte) uint64 {
	var hash = jsInitialHash64
	length := len(str)
	for i := 0; i < length; i++ {
		hash ^= (hash << 5) + uint64(str[i]) + (hash >> 2)
	}
	return hash
}
