// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcvar

import "github.com/focela/ratcatcher/utils/rcconv"

// Scan automatically checks the type of `pointer` and converts `params` to `pointer`. It supports `pointer`
// with type of `*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct` for converting.
//
// See rcconv.Scan.
func (v *Var) Scan(pointer interface{}, mapping ...map[string]string) error {
	return rcconv.Scan(v.Val(), pointer, mapping...)
}
