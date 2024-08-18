// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package hash

const bkdrSeed32 uint32 = 131
const bkdrSeed64 uint64 = 131

// ComputeBKDRHash implements the classic BKDR hash algorithm, which produces a 32-bit hash value
// using a simple and effective polynomial accumulation method.
func ComputeBKDRHash(str []byte) uint32 {
	var hash uint32 = 0
	length := len(str)
	for i := 0; i < length; i++ {
		hash = hash*bkdrSeed32 + uint32(str[i])
	}
	return hash
}

// ComputeBKDRHash64 implements the classic BKDR hash algorithm, which produces a 64-bit hash value
// using a simple and effective polynomial accumulation method.
func ComputeBKDRHash64(str []byte) uint64 {
	var hash uint64 = 0
	length := len(str)
	for i := 0; i < length; i++ {
		hash = hash*bkdrSeed64 + uint64(str[i])
	}
	return hash
}
