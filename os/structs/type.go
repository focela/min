// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package structs

// Signature returns a unique string representing the fully qualified name of the type.
func (t Type) Signature() string {
	return t.PkgPath() + "/" + t.String()
}

// FieldKeys returns the names of all the fields in the current struct type.
func (t Type) FieldKeys() []string {
	keys := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		keys[i] = t.Field(i).Name
	}
	return keys
}
