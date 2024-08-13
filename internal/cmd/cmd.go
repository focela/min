// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package cmd provides console operations, such as options/arguments parsing.
package cmd

import (
	"os"
	"regexp"
	"strings"
)

// minCmd provides console operations for the Min project.
var (
	defaultParsedArgs    []string
	defaultParsedOptions map[string]string
	argumentRegex        = regexp.MustCompile(`^\-{1,2}([\w\?\.\-]+)(=){0,1}(.*)$`)
)

// Initialize does custom initialization.
func Initialize(args ...string) {
	if len(args) == 0 {
		if len(defaultParsedArgs) == 0 && len(defaultParsedOptions) == 0 {
			args = os.Args
		} else {
			return
		}
	} else {
		defaultParsedArgs = []string{}
		defaultParsedOptions = make(map[string]string)
	}
	// Parses the provided arguments using the default algorithm.
	defaultParsedArgs, defaultParsedOptions = ParseUsingDefaultAlgorithm(args...)
}

// ParseUsingDefaultAlgorithm parses arguments using the default algorithm.
func ParseUsingDefaultAlgorithm(args ...string) (parsedArgs []string, parsedOptions map[string]string) {
	parsedArgs = []string{}
	parsedOptions = make(map[string]string)
	for i := 0; i < len(args); {
		matchedParts := argumentRegex.FindStringSubmatch(args[i])
		if len(matchedParts) > 2 {
			switch {
			case matchedParts[2] == "=":
				parsedOptions[matchedParts[1]] = matchedParts[3]
			case i < len(args)-1 && len(args[i+1]) > 0 && args[i+1][0] != '-':
				parsedOptions[matchedParts[1]] = args[i+1]
				i += 2
				continue
			default:
				parsedOptions[matchedParts[1]] = matchedParts[3]
			}
		} else {
			parsedArgs = append(parsedArgs, args[i])
		}
		i++
	}
	return
}

// GetOpt returns the option value named `name`.
func GetOpt(name string, def ...string) string {
	Initialize()
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
	Initialize()
	return defaultParsedOptions
}

// ContainsOpt checks whether option named `name` exists in the arguments.
func ContainsOpt(name string) bool {
	Initialize()
	_, ok := defaultParsedOptions[name]
	return ok
}

// GetArg returns the argument at `index`.
func GetArg(index int, def ...string) string {
	Initialize()
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
	Initialize()
	return defaultParsedArgs
}

// GetOptWithEnv returns the command line argument of the specified `key`.
// If the argument does not exist, then it returns the environment variable with specified `key`.
// It returns the default value `def` if none of them exists.
//
// Fetching Rules:
// 1. Command line arguments are in lowercase format, eg: min.package.variable;
// 2. Environment arguments are in uppercase format, eg: MIN_PACKAGE_VARIABLE;
func GetOptWithEnv(key string, def ...string) string {
	cmdKey := strings.ToLower(strings.ReplaceAll(key, "_", "."))
	if ContainsOpt(cmdKey) {
		return GetOpt(cmdKey)
	} else {
		envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		if r, ok := os.LookupEnv(envKey); ok {
			return r
		} else {
			if len(def) > 0 {
				return def[0]
			}
		}
	}
	return ""
}
