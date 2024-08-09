// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"github.com/focela/aid/internal/command"
)

const (
	// Debug key for checking if in debug mode.
	commandEnvKeyForDebugKey = "aid.debug"
)

var (
	// isDebugEnabled marks whether debug mode is enabled.
	isDebugEnabled = false
)

func init() {
	// Configure debugging mode.
	switch value := command.GetOptWithEnv(commandEnvKeyForDebugKey); value {
	case "", "0", "false":
		isDebugEnabled = false
	default:
		isDebugEnabled = true
	}
}

// IsDebugEnabled checks and returns whether debug mode is enabled.
// The debug mode is enabled when the command argument "aid.debug" or
// the environment variable "AID_DEBUG" is set.
func IsDebugEnabled() bool {
	return isDebugEnabled
}

// SetDebugEnabled enables or disables internal debug information.
func SetDebugEnabled(enabled bool) {
	isDebugEnabled = enabled
}
