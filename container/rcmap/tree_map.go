// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcmap

import "github.com/focela/ratcatcher/container/rctree"

// TreeMap based on red-black tree, alias of RedBlackTree.
type TreeMap = rctree.RedBlackTree

// NewTreeMap instantiates a tree map with the custom comparator.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
func NewTreeMap(comparator func(v1, v2 interface{}) int, safe ...bool) *TreeMap {
	return rctree.NewRedBlackTree(comparator, safe...)
}

// NewTreeMapFrom instantiates a tree map with the custom comparator and `data` map.
// Note that, the param `data` map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
// The parameter `safe` is used to specify whether using tree in concurrent-safety,
// which is false in default.
func NewTreeMapFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *TreeMap {
	return rctree.NewRedBlackTreeFrom(comparator, data, safe...)
}
