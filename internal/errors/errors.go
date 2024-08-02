// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package errors provides functionalities to manipulate errors for internal usage purposes.
package errors

import (
	"github.com/focela/plume/internal/command"
)

// StackMode is the mode that prints stack information in StackModeBrief or StackModeDetail mode.
type StackMode string

const (
	// commandEnvKeyForBrief is the command environment name for the switch key for brief error stacks.
	// Deprecated: use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "plume.error.brief"

	// commandEnvKeyForStackMode is the command environment name for the switch key for error stack mode.
	commandEnvKeyForStackMode = "plume.error.stack.mode"
)

const (
	// StackModeBrief specifies printing error stacks without framework error stacks.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies printing detailed error stacks including framework stacks.
	StackModeDetail StackMode = "detail"
)

var (
	// stackModeConfigured is the configured error stack mode variable.
	// It is brief stack mode by default.
	stackModeConfigured = StackModeBrief
)

func init() {
	// Deprecated.
	briefSetting := command.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// The error stack mode is configured using command line arguments or environment variables.
	stackModeSetting := command.GetOptWithEnv(commandEnvKeyForStackMode)
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
