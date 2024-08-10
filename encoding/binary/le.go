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

// LEEncode encodes one or multiple `values` into bytes using LittleEndian.
// It checks the type of each value in `values` and internally calls the corresponding function for byte conversion.
func LEEncode(values ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, value := range values {
		if value == nil {
			continue
		}
		switch v := value.(type) {
		case int:
			buf.Write(LEEncodeInt(v))
		case int8:
			buf.Write(LEEncodeInt8(v))
		case int16:
			buf.Write(LEEncodeInt16(v))
		case int32:
			buf.Write(LEEncodeInt32(v))
		case int64:
			buf.Write(LEEncodeInt64(v))
		case uint:
			buf.Write(LEEncodeUint(v))
		case uint8:
			buf.Write(LEEncodeUint8(v))
		case uint16:
			buf.Write(LEEncodeUint16(v))
		case uint32:
			buf.Write(LEEncodeUint32(v))
		case uint64:
			buf.Write(LEEncodeUint64(v))
		case bool:
			buf.Write(LEEncodeBool(v))
		case string:
			buf.Write(LEEncodeString(v))
		case []byte:
			buf.Write(v)
		case float32:
			buf.Write(LEEncodeFloat32(v))
		case float64:
			buf.Write(LEEncodeFloat64(v))

		default:
			if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
				intlog.Errorf(context.TODO(), "%+v", err)
				buf.Write(LEEncodeString(fmt.Sprintf("%v", v)))
			}
		}
	}
	return buf.Bytes()
}

func LEEncodeByLength(length int, values ...interface{}) []byte {
	b := LEEncode(values...)
	if len(b) < length {
		b = append(b, make([]byte, length-len(b))...)
	} else if len(b) > length {
		b = b[:length]
	}
	return b
}

func LEDecode(b []byte, values ...interface{}) error {
	buf := bytes.NewBuffer(b)
	for _, value := range values {
		if err := binary.Read(buf, binary.LittleEndian, value); err != nil {
			return errors.Wrap(err, "binary.Read failed")
		}
	}
	return nil
}

func LEEncodeString(s string) []byte {
	return []byte(s)
}

func LEDecodeToString(b []byte) string {
	return string(b)
}

func LEEncodeBool(b bool) []byte {
	if b {
		return []byte{1}
	}
	return []byte{0}
}

func LEEncodeInt(i int) []byte {
	switch {
	case i <= math.MaxInt8:
		return LEEncodeInt8(int8(i))
	case i <= math.MaxInt16:
		return LEEncodeInt16(int16(i))
	case i <= math.MaxInt32:
		return LEEncodeInt32(int32(i))
	default:
		return LEEncodeInt64(int64(i))
	}
}

func LEEncodeUint(i uint) []byte {
	switch {
	case i <= math.MaxUint8:
		return LEEncodeUint8(uint8(i))
	case i <= math.MaxUint16:
		return LEEncodeUint16(uint16(i))
	case i <= math.MaxUint32:
		return LEEncodeUint32(uint32(i))
	default:
		return LEEncodeUint64(uint64(i))
	}
}

func LEEncodeInt8(i int8) []byte {
	return []byte{byte(i)}
}

func LEEncodeUint8(i uint8) []byte {
	return []byte{i}
}

func LEEncodeInt16(i int16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(i))
	return b
}

func LEEncodeUint16(i uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, i)
	return b
}

func LEEncodeInt32(i int32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(i))
	return b
}

func LEEncodeUint32(i uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	return b
}

func LEEncodeInt64(i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

func LEEncodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}

func LEEncodeFloat32(f float32) []byte {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, bits)
	return b
}

func LEEncodeFloat64(f float64) []byte {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, bits)
	return b
}

func LEDecodeToInt(b []byte) int {
	switch len(b) {
	case 1:
		return int(LEDecodeToUint8(b))
	case 2:
		return int(LEDecodeToUint16(b))
	case 3, 4:
		return int(LEDecodeToUint32(b))
	default:
		return int(LEDecodeToUint64(b))
	}
}

func LEDecodeToUint(b []byte) uint {
	switch len(b) {
	case 1:
		return uint(LEDecodeToUint8(b))
	case 2:
		return uint(LEDecodeToUint16(b))
	case 3, 4:
		return uint(LEDecodeToUint32(b))
	default:
		return uint(LEDecodeToUint64(b))
	}
}

func LEDecodeToBool(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	return b[0] != 0
}

func LEDecodeToInt8(b []byte) int8 {
	if len(b) == 0 {
		panic("empty slice given")
	}
	return int8(b[0])
}

func LEDecodeToUint8(b []byte) uint8 {
	if len(b) == 0 {
		panic("empty slice given")
	}
	return b[0]
}

func LEDecodeToInt16(b []byte) int16 {
	return int16(binary.LittleEndian.Uint16(LEFillUpSize(b, 2)))
}

func LEDecodeToUint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(LEFillUpSize(b, 2))
}

func LEDecodeToInt32(b []byte) int32 {
	return int32(binary.LittleEndian.Uint32(LEFillUpSize(b, 4)))
}

func LEDecodeToUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(LEFillUpSize(b, 4))
}

func LEDecodeToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(LEFillUpSize(b, 8)))
}

func LEDecodeToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(LEFillUpSize(b, 8))
}

func LEDecodeToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(LEFillUpSize(b, 4)))
}

func LEDecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(LEFillUpSize(b, 8)))
}

// LEFillUpSize fills up the bytes `b` to the given length `l` using LittleEndian.
// It creates a new byte slice only if the original slice is shorter than the required length.
func LEFillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c, b)
	return c
}
