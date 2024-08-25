// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package cmd provides console operations, like options/arguments reading.
package cmd

import (
	"os"
	"regexp"
	"strings"
	"sync"
)

var (
	defaultParsedArgs    []string
	defaultParsedOptions map[string]string
	argumentRegex        = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=)?(.*)$`)
	initOnce             sync.Once
)

// Init does custom initialization.
func Init(args ...string) {
	initOnce.Do(func() {
		if len(args) == 0 {
			args = os.Args
		}
		defaultParsedArgs, defaultParsedOptions = ParseUsingDefaultAlgorithm(args...)
	})
}

// ParseUsingDefaultAlgorithm parses arguments using the default algorithm.
func ParseUsingDefaultAlgorithm(args ...string) (parsedArgs []string, parsedOptions map[string]string) {
	parsedArgs = make([]string, 0)
	parsedOptions = make(map[string]string)

	for i := 0; i < len(args); i++ {
		array := argumentRegex.FindStringSubmatch(args[i])
		if len(array) > 2 {
			if array[2] == "=" || (i == len(args)-1 || len(args[i+1]) > 0 && args[i+1][0] == '-') {
				parsedOptions[array[1]] = array[3]
			} else {
				parsedOptions[array[1]] = args[i+1]
				i++
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
	}

	return
}

// GetOpt returns the option value named `name`.
func GetOpt(name string, def ...string) string {
	Init()
	if v, ok := defaultParsedOptions[name]; ok {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetOptAll returns all parsed options.
func GetOptAll() map[string]string {
	Init()
	return defaultParsedOptions
}

// ContainsOpt checks whether the option named `name` exists in the arguments.
func ContainsOpt(name string) bool {
	Init()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetArg returns the argument at `index`.
func GetArg(index int, def ...string) string {
	Init()
	if index < len(defaultParsedArgs) {
		return defaultParsedArgs[index]
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetArgAll returns all parsed arguments.
func GetArgAll() []string {
	Init()
	return defaultParsedArgs
}

// GetOptWithEnv returns the command line argument of the specified `key`.
// If the argument does not exist, it returns the environment variable with the specified `key`.
// It returns the default value `def` if none of them exists.
//
// Fetching Rules:
// 1. Command line arguments are in lowercase format, e.g., min.package.variable;
// 2. Environment arguments are in uppercase format, e.g., MIN_PACKAGE_VARIABLE.
func GetOptWithEnv(key string, def ...string) string {
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	}

	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	if r, ok := os.LookupEnv(envKey); ok {
		return r
	}

	if len(def) > 0 {
		return def[0]
	}

	return ""
}
