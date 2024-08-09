// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package hash

// HashRS32 implements the classic RS hash algorithm for 32 bits.
func HashRS32(str []byte) uint32 {
	const (
		b     uint32 = 378551
		aInit uint32 = 63689
	)
	var (
		a           = aInit
		hash uint32 = 0
	)

	for _, c := range str {
		hash = hash*a + uint32(c)
		a *= b
	}
	return hash
}

// HashRS64 implements the classic RS hash algorithm for 64 bits.
func HashRS64(str []byte) uint64 {
	const (
		b     uint64 = 378551
		aInit uint64 = 63689
	)
	var (
		a           = aInit
		hash uint64 = 0
	)

	for _, c := range str {
		hash = hash*a + uint64(c)
		a *= b
	}
	return hash
}
