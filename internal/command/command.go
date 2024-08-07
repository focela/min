// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package command provides console operations, like options/arguments reading.
package command

import (
	"os"
	"regexp"
	"strings"
)

var (
	parsedArgs    = make([]string, 0)
	parsedOptions = make(map[string]string)
	argRegex      = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=)?(.*)$`)
)

// Init initializes the argument and option parsing.
func Init(args ...string) {
	if len(args) == 0 {
		args = os.Args
		parsedArgs = make([]string, 0)
		parsedOptions = make(map[string]string)
	}
	// Parse os.Args with the default algorithm.
	parsedArgs, parsedOptions = parseArgs(args...)
}

// parseArgs parses arguments using the default algorithm.
func parseArgs(args ...string) ([]string, map[string]string) {
	parsedArgs := make([]string, 0)
	parsedOptions := make(map[string]string)
	for i := 0; i < len(args); {
		matches := argRegex.FindStringSubmatch(args[i])
		if len(matches) > 2 {
			if matches[2] == "=" {
				parsedOptions[matches[1]] = matches[3]
			} else if i < len(args)-1 {
				if len(args[i+1]) > 0 && args[i+1][0] == '-' {
					parsedOptions[matches[1]] = matches[3]
				} else {
					parsedOptions[matches[1]] = args[i+1]
					i += 2
					continue
				}
			} else {
				parsedOptions[matches[1]] = matches[3]
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}
	return parsedArgs, parsedOptions
}

// GetOpt returns the option value named `name`.
func GetOpt(name string, def ...string) string {
	Init()
	if v, ok := parsedOptions[name]; ok {
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
	return parsedOptions
}

// ContainsOpt checks whether the option named `name` exists in the arguments.
func ContainsOpt(name string) bool {
	Init()
	_, ok := parsedOptions[name]
	return ok
}

// GetArg returns the argument at `index`.
func GetArg(index int, def ...string) string {
	Init()
	if index < len(parsedArgs) {
		return parsedArgs[index]
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetArgAll returns all parsed arguments.
func GetArgAll() []string {
	Init()
	return parsedArgs
}

// GetOptWithEnv returns the command line argument of the specified `key`.
// If the argument does not exist, it returns the environment variable with the specified `key`.
// It returns the default value `def` if none of them exist.
//
// Fetching Rules:
// 1. Command line arguments are in lowercase format, e.g., muse.package.variable;
// 2. Environment arguments are in uppercase format, e.g., MUSE_PACKAGE_VARIABLE.
func GetOptWithEnv(key string, def ...string) string {
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	}
	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	if value, ok := os.LookupEnv(envKey); ok {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}
