// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package utils

// MapPossibleItemByKey tries to find the possible key-value pair for the given key, ignoring cases and symbols.
//
// Note that this function may have low performance.
func MapPossibleItemByKey(data map[string]interface{}, key string) (foundKey string, foundValue interface{}) {
	if len(data) == 0 {
		return
	}
	if v, ok := data[key]; ok {
		return key, v
	}
	// Loop checking.
	for k, v := range data {
		if EqualFoldWithoutChars(k, key) {
			return k, v
		}
	}
	return "", nil
}

// MapContainsPossibleKey checks if the given `key` is contained in the given map `data`.
// It checks the key, ignoring cases and symbols.
//
// Note that this function may have low performance.
func MapContainsPossibleKey(data map[string]interface{}, key string) bool {
	k, _ := MapPossibleItemByKey(data, key)
	return k != ""
}
