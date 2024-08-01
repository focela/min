// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package errors provides functionalities to manipulate errors for internal usage purposes.
package errors

import "github.com/focela/plume/internal/command"

// StackMode is the mode that determines how stack information is printed.
// It can be set to either StackModeBrief or StackModeDetail.
type StackMode string

const (
	// commandEnvKeyForBrief is the command environment name for the brief error stack switch key.
	// Deprecated: Use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "plume.error.brief"

	// commandEnvKeyForStackMode is the command environment name for the error stack mode switch key.
	commandEnvKeyForStackMode = "plume.error.stack.mode"
)

const (
	// StackModeBrief specifies that only brief error stacks are printed, excluding framework error stacks.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies that detailed error stacks, including framework stacks, are printed.
	StackModeDetail StackMode = "detail"
)

var (
	// stackModeConfigured is the variable that holds the currently configured error stack mode.
	// The default mode is StackModeBrief.
	stackModeConfigured = StackModeBrief
)

func init() {
	// Check for deprecated brief stack mode setting.
	briefSetting := command.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Configure error stack mode using command line arguments or environment variables.
	stackModeSetting := command.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackModeSettingMode := StackMode(stackModeSetting)
		switch stackModeSettingMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = stackModeSettingMode
		}
	}
}

// IsStackModeBrief returns whether the current error stack mode is set to brief mode.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
