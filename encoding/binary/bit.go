// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package binary

// Bit represents a binary bit (0 or 1).
type Bit int8

// EncodeBits encodes an integer `i` into a slice of `l` bits and appends it to `bits`.
func EncodeBits(bits []Bit, i int, l int) []Bit {
	return EncodeBitsWithUint(bits, uint(i), l)
}

// EncodeBitsWithUint encodes a uint `ui` into a slice of `l` bits and appends it to `bits`.
func EncodeBitsWithUint(bits []Bit, ui uint, l int) []Bit {
	a := make([]Bit, l)
	for i := l - 1; i >= 0; i-- {
		a[i] = Bit(ui & 1)
		ui >>= 1
	}
	return append(bits, a...)
}

// EncodeBitsToBytes encodes a slice of bits into a slice of bytes.
// If the number of bits is not a multiple of 8, it pads the remaining bits with 0.
func EncodeBitsToBytes(bits []Bit) []byte {
	padLength := (8 - len(bits)%8) % 8
	for i := 0; i < padLength; i++ {
		bits = append(bits, 0)
	}
	b := make([]byte, 0, len(bits)/8)
	for i := 0; i < len(bits); i += 8 {
		b = append(b, byte(DecodeBitsToUint(bits[i:i+8])))
	}
	return b
}

// DecodeBits decodes a slice of bits into an integer.
func DecodeBits(bits []Bit) int {
	return int(DecodeBitsToUint(bits))
}

// DecodeBitsToUint decodes a slice of bits into an unsigned integer.
func DecodeBitsToUint(bits []Bit) uint {
	var v uint
	for _, bit := range bits {
		v = (v << 1) | uint(bit)
	}
	return v
}

// DecodeBytesToBits decodes a slice of bytes into a slice of bits.
func DecodeBytesToBits(bs []byte) []Bit {
	var bits []Bit
	for _, b := range bs {
		bits = EncodeBitsWithUint(bits, uint(b), 8)
	}
	return bits
}
