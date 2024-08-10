// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package structs provides functions for struct information retrieving.
package structs

import (
	"reflect"

	"github.com/focela/aid/internal/empty"
	"github.com/focela/aid/internal/utils"
	"github.com/focela/aid/util/tags"
)

// Tag returns the value associated with key in the tag string.
// If there is no such key in the tag, Tag returns the empty string.
func (f *Field) Tag(key string) string {
	s := f.Field.Tag.Get(key)
	if s != "" {
		s = tags.Parse(s)
	}
	return s
}

// TagLookup returns the value associated with key in the tag string.
// If the key is present in the tag the value (which may be empty) is returned.
// Otherwise, the returned value will be the empty string.
func (f *Field) TagLookup(key string) (value string, ok bool) {
	value, ok = f.Field.Tag.Lookup(key)
	if ok && value != "" {
		value = tags.Parse(value)
	}
	return
}

// IsEmbedded returns true if the given field is an anonymous field (embedded).
func (f *Field) IsEmbedded() bool {
	return f.Field.Anonymous
}

// TagStr returns the tag string of the field.
func (f *Field) TagStr() string {
	return string(f.Field.Tag)
}

// TagMap returns all the tags of the field along with their value strings as a map.
func (f *Field) TagMap() map[string]string {
	data := ParseTag(f.TagStr())
	for k, v := range data {
		data[k] = utils.StripSlashes(tags.Parse(v))
	}
	return data
}

// IsExported returns true if the given field is exported.
func (f *Field) IsExported() bool {
	return f.Field.PkgPath == ""
}

// Name returns the name of the given field.
func (f *Field) Name() string {
	return f.Field.Name
}

// Type returns the type of the given field.
// Note that this Type is not reflect.Type. If you need reflect.Type, please use Field.Type().Type.
func (f *Field) Type() Type {
	return Type{
		Type: f.Field.Type,
	}
}

// Kind returns the reflect.Kind for the Value of Field `f`.
func (f *Field) Kind() reflect.Kind {
	return f.Value.Kind()
}

// OriginalKind retrieves and returns the original reflect.Kind for the Value of Field `f`.
func (f *Field) OriginalKind() reflect.Kind {
	reflectType := f.Value.Type()
	for reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	return reflectType.Kind()
}

// OriginalValue retrieves and returns the original reflect.Value of Field `f`.
func (f *Field) OriginalValue() reflect.Value {
	reflectValue := f.Value
	for reflectValue.Kind() == reflect.Ptr && !f.IsNil() {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

// IsEmpty checks and returns whether the value of this Field is empty.
func (f *Field) IsEmpty() bool {
	return empty.IsEmpty(f.Value)
}

// IsNil checks and returns whether the value of this Field is nil.
func (f *Field) IsNil(traceSource ...bool) bool {
	return empty.IsNil(f.Value, traceSource...)
}
