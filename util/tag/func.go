// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package tag

import (
	"regexp"

	"github.com/focela/aid/errors"
)

var (
	data  = make(map[string]string)
	regex = regexp.MustCompile(`\{(.+?)\}`)
)

// Set stores the tag content for the specified name.
// It will panic if the `name` already exists.
func Set(name, value string) {
	if _, ok := data[name]; ok {
		panic(errors.Newf(`value for tag name "%s" already exists`, name))
	}
	data[name] = value
}

// SetOver stores the tag content, overwriting the existing value if `name` already exists.
func SetOver(name, value string) {
	data[name] = value
}

// Sets stores multiple tag contents from the provided map.
func Sets(m map[string]string) {
	for k, v := range m {
		Set(k, v)
	}
}

// SetsOver stores multiple tag contents, overwriting existing values if the `name` already exists.
func SetsOver(m map[string]string) {
	for k, v := range m {
		SetOver(k, v)
	}
}

// Get retrieves and returns the stored tag content for the specified name.
func Get(name string) string {
	return data[name]
}

// Parse parses the given content, replacing all tag name variables with their corresponding content.
// For example:
// tag.Set("demo", "content")
// Parse(`This is {demo}`) -> `This is content`.
func Parse(content string) string {
	return regex.ReplaceAllStringFunc(content, func(s string) string {
		if v, ok := data[s[1:len(s)-1]]; ok {
			return v
		}
		return s
	})
}
