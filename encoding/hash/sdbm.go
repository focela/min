// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

// ComputeSDBMHash implements the classic SDBM hash algorithm for 32 bits.
func ComputeSDBMHash(str []byte) uint32 {
	var hash uint32
	length := len(str)
	for i := 0; i < length; i++ {
		// This is equivalent to: hash = 65599 * hash + uint32(str[i])
		hash = uint32(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

// ComputeSDBMHash64 implements the classic SDBM hash algorithm for 64 bits.
func ComputeSDBMHash64(str []byte) uint64 {
	var hash uint64
	length := len(str)
	for i := 0; i < length; i++ {
		// This is equivalent to: hash = 65599 * hash + uint64(str[i])
		hash = uint64(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}
