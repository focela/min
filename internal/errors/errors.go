// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package errors provides functionalities to manipulate errors for internal usage purposes.
package errors

import "github.com/focela/orca/internal/cmd"

// StackMode is the mode for printing stack information, either in StackModeBrief or StackModeDetail.
type StackMode string

const (
	// commandEnvKeyForStackMode is the command environment name for the switch key for error stack mode.
	commandEnvKeyForStackMode = "go.error.stack.mode"
)

const (
	// StackModeBrief specifies printing error stacks without framework error stacks.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies printing detailed error stacks including framework stacks.
	StackModeDetail StackMode = "detail"
)

var (
	// stackModeConfigured is the configured error stack mode variable.
	// It is set to brief stack mode by default.
	stackModeConfigured = StackModeBrief
)

func init() {
	// The error stack mode is configured using command line arguments or environment variables.
	stackModeSetting := cmd.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackModeSettingMode := StackMode(stackModeSetting)
		switch stackModeSettingMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = stackModeSettingMode
		}
	}
}

// IsStackModeBrief returns whether the current error stack mode is in brief mode.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
