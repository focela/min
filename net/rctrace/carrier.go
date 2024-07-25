// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rctrace

import (
	"github.com/focela/ratcatcher/internal/json"
	"github.com/focela/ratcatcher/utils/rcconv"
)

// Carrier is the storage medium used by a TextMapPropagator.
type Carrier map[string]interface{}

// NewCarrier creates and returns a Carrier.
func NewCarrier(data ...map[string]interface{}) Carrier {
	if len(data) > 0 && data[0] != nil {
		return data[0]
	}
	return make(map[string]interface{})
}

// Get returns the value associated with the passed key.
func (c Carrier) Get(k string) string {
	return rcconv.String(c[k])
}

// Set stores the key-value pair.
func (c Carrier) Set(k, v string) {
	c[k] = v
}

// Keys lists the keys stored in this carrier.
func (c Carrier) Keys() []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}

// MustMarshal .returns the JSON encoding of c
func (c Carrier) MustMarshal() []byte {
	b, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return b
}

// String converts and returns current Carrier as string.
func (c Carrier) String() string {
	return string(c.MustMarshal())
}

// UnmarshalJSON implements interface UnmarshalJSON for package json.
func (c Carrier) UnmarshalJSON(b []byte) error {
	carrier := NewCarrier(nil)
	return json.UnmarshalUseNumber(b, carrier)
}
