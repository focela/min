// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package binary provides useful API for handling binary/bytes data.
//
// Note that package binary encodes the data using LittleEndian by default.
package binary

func Encode(values ...interface{}) []byte {
	return LEEncode(values...)
}

func EncodeByLength(length int, values ...interface{}) []byte {
	return LEEncodeByLength(length, values...)
}

func Decode(b []byte, values ...interface{}) error {
	return LEDecode(b, values...)
}

func EncodeString(s string) []byte {
	return LEEncodeString(s)
}

func DecodeToString(b []byte) string {
	return LEDecodeToString(b)
}

func EncodeBool(b bool) []byte {
	return LEEncodeBool(b)
}

func EncodeInt(i int) []byte {
	return LEEncodeInt(i)
}

func EncodeUint(i uint) []byte {
	return LEEncodeUint(i)
}

func EncodeInt8(i int8) []byte {
	return LEEncodeInt8(i)
}

func EncodeUint8(i uint8) []byte {
	return LEEncodeUint8(i)
}

func EncodeInt16(i int16) []byte {
	return LEEncodeInt16(i)
}

func EncodeUint16(i uint16) []byte {
	return LEEncodeUint16(i)
}

func EncodeInt32(i int32) []byte {
	return LEEncodeInt32(i)
}

func EncodeUint32(i uint32) []byte {
	return LEEncodeUint32(i)
}

func EncodeInt64(i int64) []byte {
	return LEEncodeInt64(i)
}

func EncodeUint64(i uint64) []byte {
	return LEEncodeUint64(i)
}

func EncodeFloat32(f float32) []byte {
	return LEEncodeFloat32(f)
}

func EncodeFloat64(f float64) []byte {
	return LEEncodeFloat64(f)
}

func DecodeToInt(b []byte) int {
	return LEDecodeToInt(b)
}

func DecodeToUint(b []byte) uint {
	return LEDecodeToUint(b)
}

func DecodeToBool(b []byte) bool {
	return LEDecodeToBool(b)
}

func DecodeToInt8(b []byte) int8 {
	return LEDecodeToInt8(b)
}

func DecodeToUint8(b []byte) uint8 {
	return LEDecodeToUint8(b)
}

func DecodeToInt16(b []byte) int16 {
	return LEDecodeToInt16(b)
}

func DecodeToUint16(b []byte) uint16 {
	return LEDecodeToUint16(b)
}

func DecodeToInt32(b []byte) int32 {
	return LEDecodeToInt32(b)
}

func DecodeToUint32(b []byte) uint32 {
	return LEDecodeToUint32(b)
}

func DecodeToInt64(b []byte) int64 {
	return LEDecodeToInt64(b)
}

func DecodeToUint64(b []byte) uint64 {
	return LEDecodeToUint64(b)
}

func DecodeToFloat32(b []byte) float32 {
	return LEDecodeToFloat32(b)
}

func DecodeToFloat64(b []byte) float64 {
	return LEDecodeToFloat64(b)
}
