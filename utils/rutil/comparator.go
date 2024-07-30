// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rutil

import (
	"strings"

	"github.com/focela/ratcatcher/utils/rconv"
)

// Comparator is a function that compare a and b, and returns the result as int.
//
// Should return a number:
//
//	negative , if a < b
//	zero     , if a == b
//	positive , if a > b
type Comparator func(a, b interface{}) int

// ComparatorString provides a fast comparison on strings.
func ComparatorString(a, b interface{}) int {
	return strings.Compare(rconv.String(a), rconv.String(b))
}

// ComparatorInt provides a basic comparison on int.
func ComparatorInt(a, b interface{}) int {
	return rconv.Int(a) - rconv.Int(b)
}

// ComparatorInt8 provides a basic comparison on int8.
func ComparatorInt8(a, b interface{}) int {
	return int(rconv.Int8(a) - rconv.Int8(b))
}

// ComparatorInt16 provides a basic comparison on int16.
func ComparatorInt16(a, b interface{}) int {
	return int(rconv.Int16(a) - rconv.Int16(b))
}

// ComparatorInt32 provides a basic comparison on int32.
func ComparatorInt32(a, b interface{}) int {
	return int(rconv.Int32(a) - rconv.Int32(b))
}

// ComparatorInt64 provides a basic comparison on int64.
func ComparatorInt64(a, b interface{}) int {
	return int(rconv.Int64(a) - rconv.Int64(b))
}

// ComparatorUint provides a basic comparison on uint.
func ComparatorUint(a, b interface{}) int {
	return int(rconv.Uint(a) - rconv.Uint(b))
}

// ComparatorUint8 provides a basic comparison on uint8.
func ComparatorUint8(a, b interface{}) int {
	return int(rconv.Uint8(a) - rconv.Uint8(b))
}

// ComparatorUint16 provides a basic comparison on uint16.
func ComparatorUint16(a, b interface{}) int {
	return int(rconv.Uint16(a) - rconv.Uint16(b))
}

// ComparatorUint32 provides a basic comparison on uint32.
func ComparatorUint32(a, b interface{}) int {
	return int(rconv.Uint32(a) - rconv.Uint32(b))
}

// ComparatorUint64 provides a basic comparison on uint64.
func ComparatorUint64(a, b interface{}) int {
	return int(rconv.Uint64(a) - rconv.Uint64(b))
}

// ComparatorFloat32 provides a basic comparison on float32.
func ComparatorFloat32(a, b interface{}) int {
	aFloat := rconv.Float32(a)
	bFloat := rconv.Float32(b)
	if aFloat == bFloat {
		return 0
	}
	if aFloat > bFloat {
		return 1
	}
	return -1
}

// ComparatorFloat64 provides a basic comparison on float64.
func ComparatorFloat64(a, b interface{}) int {
	aFloat := rconv.Float64(a)
	bFloat := rconv.Float64(b)
	if aFloat == bFloat {
		return 0
	}
	if aFloat > bFloat {
		return 1
	}
	return -1
}

// ComparatorByte provides a basic comparison on byte.
func ComparatorByte(a, b interface{}) int {
	return int(rconv.Byte(a) - rconv.Byte(b))
}

// ComparatorRune provides a basic comparison on rune.
func ComparatorRune(a, b interface{}) int {
	return int(rconv.Rune(a) - rconv.Rune(b))
}

// ComparatorTime provides a basic comparison on time.Time.
func ComparatorTime(a, b interface{}) int {
	aTime := rconv.Time(a)
	bTime := rconv.Time(b)
	switch {
	case aTime.After(bTime):
		return 1
	case aTime.Before(bTime):
		return -1
	default:
		return 0
	}
}
