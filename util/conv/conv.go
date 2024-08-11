// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package conv implements powerful and convenient converting functionality for any types of variables.
//
// This package should keep much fewer dependencies with other packages.
package conv

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/focela/aid/encoding/binary"
	"github.com/focela/aid/internal/intlog"
	"github.com/focela/aid/internal/json"
	"github.com/focela/aid/internal/reflection"
	"github.com/focela/aid/os/times"
	"github.com/focela/aid/util/tags"
)

var (
	// Empty strings map.
	emptyStrings = map[string]struct{}{
		"":      {},
		"0":     {},
		"no":    {},
		"off":   {},
		"false": {},
	}

	// StructTagPriorities defines the default priority tags for Map*/Struct* functions.
	// Note: The `conv/param` tags are from an older version of the package.
	// It's recommended to use the short tag `c/p` instead in the future.
	StructTagPriorities = []string{
		tags.Conv, tags.Param, tags.ConvShort, tags.ParamShort, tags.Json,
	}
)

// ToByte converts `any` to byte.
func ToByte(any interface{}) byte {
	if v, ok := any.(byte); ok {
		return v
	}
	return ToUint8(any)
}

// ToBytes converts `any` to []byte.
func ToBytes(any interface{}) []byte {
	if any == nil {
		return nil
	}
	switch value := any.(type) {
	case string:
		return []byte(value)
	case []byte:
		return value
	default:
		if f, ok := value.(iBytes); ok {
			return f.Bytes()
		}

		originValueAndKind := reflection.OriginValueAndKind(any)
		switch originValueAndKind.OriginKind {
		case reflect.Map:
			bytes, err := json.Marshal(any)
			if err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
				return nil
			}
			return bytes
		case reflect.Array, reflect.Slice:
			bytes := make([]byte, originValueAndKind.OriginValue.Len())
			for i := range bytes {
				int32Value := ToInt32(originValueAndKind.OriginValue.Index(i).Interface())
				if int32Value < 0 || int32Value > math.MaxUint8 {
					return binary.Encode(any)
				}
				bytes[i] = byte(int32Value)
			}
			return bytes
		}
		return binary.Encode(any)
	}
}

// ToRune converts `any` to rune.
func ToRune(any interface{}) rune {
	if v, ok := any.(rune); ok {
		return v
	}
	return ToInt32(any)
}

// ToRunes converts `any` to []rune.
func ToRunes(any interface{}) []rune {
	if v, ok := any.([]rune); ok {
		return v
	}
	return []rune(ToString(any))
}

// ToString converts `any` to string.
// It's the most commonly used converting function.
func ToString(any interface{}) string {
	if any == nil {
		return ""
	}
	switch value := any.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	case times.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *times.Time:
		if value == nil {
			return ""
		}
		return value.String()
	default:
		// Empty checks.
		if value == nil {
			return ""
		}
		if f, ok := value.(iString); ok {
			// If the variable implements the String() interface,
			// use that interface to perform the conversion.
			return f.String()
		}
		if f, ok := value.(iError); ok {
			// If the variable implements the Error() interface,
			// use that interface to perform the conversion.
			return f.Error()
		}
		// Reflect checks.
		rv := reflect.ValueOf(value)
		switch kind := rv.Kind(); kind {
		case reflect.Chan, reflect.Map, reflect.Slice, reflect.Func, reflect.Ptr, reflect.Interface, reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
		case reflect.String:
			return rv.String()
		}
		if kind == reflect.Ptr {
			return ToString(rv.Elem().Interface())
		}
		// Finally, use json.Marshal to convert.
		if jsonContent, err := json.Marshal(value); err != nil {
			return fmt.Sprint(value)
		} else {
			return string(jsonContent)
		}
	}
}

// ToBool converts `any` to bool.
// It returns false if `any` is: false, "", 0, "false", "off", "no", empty slice/map.
func ToBool(any interface{}) bool {
	if any == nil {
		return false
	}
	switch value := any.(type) {
	case bool:
		return value
	case []byte:
		_, exists := emptyStrings[strings.ToLower(string(value))]
		return !exists
	case string:
		_, exists := emptyStrings[strings.ToLower(value)]
		return !exists
	default:
		if f, ok := value.(iBool); ok {
			return f.Bool()
		}
		rv := reflect.ValueOf(any)
		switch rv.Kind() {
		case reflect.Ptr:
			return !rv.IsNil()
		case reflect.Map, reflect.Array, reflect.Slice:
			return rv.Len() != 0
		case reflect.Struct:
			return true
		default:
			_, exists := emptyStrings[strings.ToLower(ToString(any))]
			return !exists
		}
	}
}

// checkJsonAndUnmarshalUseNumber checks if `any` is a JSON formatted string value and does converting using `json.UnmarshalUseNumber`.
func checkJsonAndUnmarshalUseNumber(any interface{}, target interface{}) bool {
	switch r := any.(type) {
	case []byte:
		if json.Valid(r) {
			if err := json.UnmarshalUseNumber(r, &target); err != nil {
				return false
			}
			return true
		}
	case string:
		anyAsBytes := []byte(r)
		if json.Valid(anyAsBytes) {
			if err := json.UnmarshalUseNumber(anyAsBytes, &target); err != nil {
				return false
			}
			return true
		}
	}
	return false
}
