// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

const djbInitialHash32 uint32 = 5381
const djbInitialHash64 uint64 = 5381

// ComputeDJBHash implements the classic DJB hash algorithm, which produces a 32-bit hash value
// using a simple and efficient bitwise operation method.
func ComputeDJBHash(str []byte) uint32 {
	var hash = djbInitialHash32
	length := len(str)
	for i := 0; i < length; i++ {
		hash += (hash << 5) + uint32(str[i])
	}
	return hash
}

// ComputeDJBHash64 implements the classic DJB hash algorithm, which produces a 64-bit hash value
// using a simple and efficient bitwise operation method.
func ComputeDJBHash64(str []byte) uint64 {
	var hash = djbInitialHash64
	length := len(str)
	for i := 0; i < length; i++ {
		hash += (hash << 5) + uint64(str[i])
	}
	return hash
}
