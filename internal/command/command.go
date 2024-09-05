// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package command provides console operations, like options/arguments reading.
package command

import (
	"os"
	"regexp"
	"strings"
)

var (
	isInitialized        bool
	defaultParsedArgs    = make([]string, 0)
	defaultParsedOptions = make(map[string]string)
	argumentRegex        = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`)
)

// Init initializes argument parsing if not already initialized.
func Init(args ...string) {
	if isInitialized {
		return
	}
	isInitialized = true

	if len(args) == 0 {
		args = os.Args
	}

	defaultParsedArgs, defaultParsedOptions = ParseArgs(args...)
}

// ParseArgs parses arguments using the default algorithm.
func ParseArgs(args ...string) (parsedArgs []string, parsedOptions map[string]string) {
	parsedArgs = make([]string, 0)
	parsedOptions = make(map[string]string)

	for i := 0; i < len(args); i++ {
		arg := args[i]
		matches := argumentRegex.FindStringSubmatch(arg)

		if len(matches) > 2 {
			key := matches[1]
			if matches[2] == "=" {
				parsedOptions[key] = matches[3]
			} else if i < len(args)-1 && args[i+1][0] != '-' {
				parsedOptions[key] = args[i+1]
				i++ // skip the next arg
			} else {
				parsedOptions[key] = matches[3]
			}
		} else {
			parsedArgs = append(parsedArgs, arg)
		}
	}
	return
}

// GetOption returns the option value named `name`.
func GetOption(name string, def ...string) string {
	Init()
	if v, ok := defaultParsedOptions[name]; ok {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetAllOptions returns all parsed options.
func GetAllOptions() map[string]string {
	Init()
	return defaultParsedOptions
}

// HasOption checks whether option named `name` exists in the arguments.
func HasOption(name string) bool {
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

// GetAllArgs returns all parsed arguments.
func GetAllArgs() []string {
	Init()
	return defaultParsedArgs
}

// GetOptionWithEnv returns the command line argument or environment variable.
func GetOptionWithEnv(key string, def ...string) string {
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if HasOption(cmdKey) {
		return GetOption(cmdKey)
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
