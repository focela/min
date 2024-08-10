// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package json provides JSON operations wrapping, ignoring stdlib or third-party JSON libraries.
package json

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/focela/aid/errors"
)

// RawMessage is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
type RawMessage = json.RawMessage

// Marshal adapts to json/encoding Marshal API.
// It returns the JSON encoding of v, adapting to the json/encoding Marshal API.
// Refer to https://godoc.org/encoding/json#Marshal for more information.
func Marshal(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, errors.Wrap(err, `json.Marshal failed`)
	}
	return b, nil
}

// MarshalIndent is the same as json.MarshalIndent.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	b, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return nil, errors.Wrap(err, `json.MarshalIndent failed`)
	}
	return b, nil
}

// Unmarshal adapts to json/encoding Unmarshal API.
// It parses the JSON-encoded data and stores the result in the value pointed to by v.
// Refer to https://godoc.org/encoding/json#Unmarshal for more information.
func Unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return errors.Wrap(err, `json.Unmarshal failed`)
	}
	return nil
}

// UnmarshalUseNumber decodes the JSON data bytes to target interface using the number option.
func UnmarshalUseNumber(data []byte, v interface{}) error {
	decoder := NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(v); err != nil {
		return errors.Wrap(err, `json.UnmarshalUseNumber failed`)
	}
	return nil
}

// NewEncoder is the same as json.NewEncoder.
func NewEncoder(writer io.Writer) *json.Encoder {
	return json.NewEncoder(writer)
}

// NewDecoder adapts to the json/stream NewDecoder API.
// It returns a new decoder that reads from r.
func NewDecoder(reader io.Reader) *json.Decoder {
	return json.NewDecoder(reader)
}

// Valid reports whether data is a valid JSON encoding.
func Valid(data []byte) bool {
	return json.Valid(data)
}
