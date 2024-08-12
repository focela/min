// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package errors provides functionalities to manipulate errors for internal usage purposes.
package errors

import (
	"github.com/focela/aid/internal/command"
)

// StackMode is the mode for printing stack information in either brief or detailed mode.
type StackMode string

const (
	// commandEnvKeyBrief is the command environment name for the brief error stack mode.
	// Deprecated: use commandEnvKeyStackMode instead for more flexibility.
	commandEnvKeyBrief = "aid.errors.brief"

	// commandEnvKeyStackMode is the command environment name for the error stack mode.
	commandEnvKeyStackMode = "aid.errors.stack.mode"
)

const (
	// StackModeBrief specifies printing error stacks without framework-specific stacks.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies printing detailed error stacks including framework-specific stacks.
	StackModeDetail StackMode = "detail"
)

var (
	// stackModeConfigured is the configured error stack mode.
	// Defaults to StackModeBrief.
	stackModeConfigured = StackModeBrief
)

func init() {
	// Check the deprecated brief setting.
	briefSetting := command.GetOptWithEnv(commandEnvKeyBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Configure the error stack mode using command line arguments or environment variables.
	stackModeSetting := command.GetOptWithEnv(commandEnvKeyStackMode)
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
