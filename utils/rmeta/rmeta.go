// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package rmeta provides embedded meta data feature for struct.
package rmeta

import (
	"github.com/focela/ratcatcher/container/rvar"
	"github.com/focela/ratcatcher/os/rstructs"
)

// Meta is used as an embedded attribute for struct to enabled metadata feature.
type Meta struct{}

const (
	metaAttributeName = "Meta"       // metaAttributeName is the attribute name of metadata in struct.
	metaTypeName      = "rmeta.Meta" // metaTypeName is for type string comparison.
)

// Data retrieves and returns all metadata from `object`.
func Data(object interface{}) map[string]string {
	reflectType, err := rstructs.StructType(object)
	if err != nil {
		return nil
	}
	if field, ok := reflectType.FieldByName(metaAttributeName); ok {
		if field.Type.String() == metaTypeName {
			return rstructs.ParseTag(string(field.Tag))
		}
	}
	return map[string]string{}
}

// Get retrieves and returns specified metadata by `key` from `object`.
func Get(object interface{}, key string) *rvar.Var {
	v, ok := Data(object)[key]
	if !ok {
		return nil
	}
	return rvar.New(v)
}
