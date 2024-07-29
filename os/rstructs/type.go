// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rstructs

// Signature returns a unique string as this type.
func (t Type) Signature() string {
	return t.PkgPath() + "/" + t.String()
}

// FieldKeys returns the keys of current struct/map.
func (t Type) FieldKeys() []string {
	keys := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		keys[i] = t.Field(i).Name
	}
	return keys
}
