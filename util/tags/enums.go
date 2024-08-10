// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package tags

import (
	"sync"

	"github.com/focela/aid/internal/json"
)

var (
	// Type name => enums json.
	enumsMap = make(map[string]json.RawMessage)

	// Mutex to ensure safe concurrent access to enumsMap.
	enumsMu sync.RWMutex
)

// SetGlobalEnums sets the global enums into the package.
// Note that this operation is not concurrent safe.
func SetGlobalEnums(enumsJson string) error {
	enumsMu.Lock()
	defer enumsMu.Unlock()
	return json.Unmarshal([]byte(enumsJson), &enumsMap)
}

// GetGlobalEnums retrieves and returns the global enums.
func GetGlobalEnums() (string, error) {
	enumsMu.RLock()
	defer enumsMu.RUnlock()

	enumsBytes, err := json.Marshal(enumsMap)
	if err != nil {
		return "", err
	}
	return string(enumsBytes), nil
}

// GetEnumsByType retrieves and returns the stored enums json by type name.
func GetEnumsByType(typeName string) string {
	enumsMu.RLock()
	defer enumsMu.RUnlock()

	if val, ok := enumsMap[typeName]; ok {
		return string(val)
	}
	return ""
}
