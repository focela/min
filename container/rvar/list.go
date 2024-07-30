// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rvar

import "github.com/focela/ratcatcher/utils/rutil"

// ListItemValues retrieves and returns the elements of all item struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.
func (v *Var) ListItemValues(key interface{}) (values []interface{}) {
	return rutil.ListItemValues(v.Val(), key)
}

// ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.
func (v *Var) ListItemValuesUnique(key string) []interface{} {
	return rutil.ListItemValuesUnique(v.Val(), key)
}
