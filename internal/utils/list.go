// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
)

// ListToMapByKey converts `list` to a map[string]interface{} with the key specified by `key`.
// Note that the item value may be of type slice.
func ListToMapByKey(list []map[string]interface{}, key string) map[string]interface{} {
	m := make(map[string]interface{})
	for _, item := range list {
		if k, ok := item[key]; ok {
			s := fmt.Sprintf(`%v`, k)
			if existing, exists := m[s]; exists {
				if existingSlice, ok := existing.([]interface{}); ok {
					m[s] = append(existingSlice, item)
				} else {
					m[s] = []interface{}{existing, item}
				}
			} else {
				m[s] = item
			}
		}
	}
	return m
}
