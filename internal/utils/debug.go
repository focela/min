// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import "github.com/focela/ratcatcher/internal/cmd"

const (
	// Debug key for checking if in debug mode.
	commandEnvKeyForDebugKey = "rc.debug"
)

var (
	// isDebugEnabled marks whether Ratcatcher debug mode is enabled.
	isDebugEnabled = false
)

func init() {
	// Debugging configured.
	value := cmd.GetOptWithEnv(commandEnvKeyForDebugKey)
	if value == "" || value == "0" || value == "false" {
		isDebugEnabled = false
	} else {
		isDebugEnabled = true
	}
}

// IsDebugEnabled checks and returns whether debug mode is enabled.
// The debug mode is enabled when command argument "rc.debug" or environment "RC_DEBUG" is passed.
func IsDebugEnabled() bool {
	return isDebugEnabled
}

// SetDebugEnabled enables/disables the internal debug info.
func SetDebugEnabled(enabled bool) {
	isDebugEnabled = enabled
}
