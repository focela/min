// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rvar

import "github.com/focela/ratcatcher/utils/rconv"

// Struct maps value of `v` to `pointer`.
// The parameter `pointer` should be a pointer to a struct instance.
// The parameter `mapping` is used to specify the key-to-attribute mapping rules.
func (v *Var) Struct(pointer interface{}, mapping ...map[string]string) error {
	return rconv.Struct(v.Val(), pointer, mapping...)
}

// Structs converts and returns `v` as given struct slice.
func (v *Var) Structs(pointer interface{}, mapping ...map[string]string) error {
	return rconv.Structs(v.Val(), pointer, mapping...)
}
