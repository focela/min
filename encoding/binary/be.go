// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package binary

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/focela/aid/errors"
	"github.com/focela/aid/internal/intlog"
)

// BEEncode encodes one or multiple `values` into bytes using BigEndian.
func BEEncode(values ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, value := range values {
		if value == nil {
			return buf.Bytes()
		}

		switch v := value.(type) {
		case int:
			buf.Write(BEEncodeInt(v))
		case int8:
			buf.Write(BEEncodeInt8(v))
		case int16:
			buf.Write(BEEncodeInt16(v))
		case int32:
			buf.Write(BEEncodeInt32(v))
		case int64:
			buf.Write(BEEncodeInt64(v))
		case uint:
			buf.Write(BEEncodeUint(v))
		case uint8:
			buf.Write(BEEncodeUint8(v))
		case uint16:
			buf.Write(BEEncodeUint16(v))
		case uint32:
			buf.Write(BEEncodeUint32(v))
		case uint64:
			buf.Write(BEEncodeUint64(v))
		case bool:
			buf.Write(BEEncodeBool(v))
		case string:
			buf.Write(BEEncodeString(v))
		case []byte:
			buf.Write(v)
		case float32:
			buf.Write(BEEncodeFloat32(v))
		case float64:
			buf.Write(BEEncodeFloat64(v))
		default:
			if err := binary.Write(buf, binary.BigEndian, v); err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
				buf.Write(BEEncodeString(fmt.Sprintf("%v", v)))
			}
		}
	}
	return buf.Bytes()
}

func BEEncodeByLength(length int, values ...interface{}) []byte {
	b := BEEncode(values...)
	if len(b) < length {
		b = append(b, make([]byte, length-len(b))...)
	} else if len(b) > length {
		b = b[:length]
	}
	return b
}

func BEDecode(b []byte, values ...interface{}) error {
	buf := bytes.NewBuffer(b)
	for _, value := range values {
		if err := binary.Read(buf, binary.BigEndian, value); err != nil {
			return errors.Wrap(err, `binary.Read failed`)
		}
	}
	return nil
}

func BEEncodeString(s string) []byte {
	return []byte(s)
}

func BEDecodeToString(b []byte) string {
	return string(b)
}

func BEEncodeBool(b bool) []byte {
	if b {
		return []byte{1}
	}
	return []byte{0}
}

func BEEncodeInt(i int) []byte {
	switch {
	case i <= math.MaxInt8:
		return BEEncodeInt8(int8(i))
	case i <= math.MaxInt16:
		return BEEncodeInt16(int16(i))
	case i <= math.MaxInt32:
		return BEEncodeInt32(int32(i))
	default:
		return BEEncodeInt64(int64(i))
	}
}

func BEEncodeUint(i uint) []byte {
	switch {
	case i <= math.MaxUint8:
		return BEEncodeUint8(uint8(i))
	case i <= math.MaxUint16:
		return BEEncodeUint16(uint16(i))
	case i <= math.MaxUint32:
		return BEEncodeUint32(uint32(i))
	default:
		return BEEncodeUint64(uint64(i))
	}
}

func BEEncodeInt8(i int8) []byte {
	return []byte{byte(i)}
}

func BEEncodeUint8(i uint8) []byte {
	return []byte{i}
}

func BEEncodeInt16(i int16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(i))
	return b
}

func BEEncodeUint16(i uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}

func BEEncodeInt32(i int32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}

func BEEncodeUint32(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

func BEEncodeInt64(i int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func BEEncodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

func BEEncodeFloat32(f float32) []byte {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, bits)
	return b
}

func BEEncodeFloat64(f float64) []byte {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, bits)
	return b
}

func BEDecodeToInt(b []byte) int {
	switch len(b) {
	case 1:
		return int(BEDecodeToUint8(b))
	case 2:
		return int(BEDecodeToUint16(b))
	case 4:
		return int(BEDecodeToUint32(b))
	default:
		return int(BEDecodeToUint64(b))
	}
}

func BEDecodeToUint(b []byte) uint {
	switch len(b) {
	case 1:
		return uint(BEDecodeToUint8(b))
	case 2:
		return uint(BEDecodeToUint16(b))
	case 4:
		return uint(BEDecodeToUint32(b))
	default:
		return uint(BEDecodeToUint64(b))
	}
}

func BEDecodeToBool(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	return !bytes.Equal(b, make([]byte, len(b)))
}

func BEDecodeToInt8(b []byte) int8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return int8(b[0])
}

func BEDecodeToUint8(b []byte) uint8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return b[0]
}

func BEDecodeToInt16(b []byte) int16 {
	return int16(binary.BigEndian.Uint16(BEFillUpSize(b, 2)))
}

func BEDecodeToUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(BEFillUpSize(b, 2))
}

func BEDecodeToInt32(b []byte) int32 {
	return int32(binary.BigEndian.Uint32(BEFillUpSize(b, 4)))
}

func BEDecodeToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(BEFillUpSize(b, 4))
}

func BEDecodeToInt64(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(BEFillUpSize(b, 8)))
}

func BEDecodeToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(BEFillUpSize(b, 8))
}

func BEDecodeToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(BEFillUpSize(b, 4)))
}

func BEDecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(BEFillUpSize(b, 8)))
}

// BEFillUpSize fills up the bytes `b` to given length `l` using BigEndian.
func BEFillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c[l-len(b):], b)
	return c
}
