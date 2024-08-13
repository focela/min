// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package errors provides functionalities to manipulate errors for internal usage.
package errors

import (
	"github.com/focela/min/internal/cmd"
)

// StackMode defines the mode for printing stack information, either in brief or detailed mode.
type StackMode string

const (
	// commandEnvKeyForStackModeBrief is the command environment name for switch key for brief error stack.
	// Deprecated: use commandEnvKeyForStackMode instead.
	commandEnvKeyForStackModeBrief = "min.error.brief"

	// commandEnvKeyForStackMode is the command environment name for switch key for error stack mode.
	commandEnvKeyForStackMode = "min.error.stack.mode"
)

const (
	// StackModeBrief specifies that all error stacks print no framework error stacks.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies that all error stacks print detailed error stacks including framework stacks.
	StackModeDetail StackMode = "detail"
)

var (
	// currentStackMode is the configured error stack mode variable.
	// The default stack mode is brief.
	currentStackMode = StackModeBrief
)

func init() {
	// Deprecated.
	briefSetting := cmd.GetOptWithEnv(commandEnvKeyForStackModeBrief)
	if briefSetting == "1" || briefSetting == "true" {
		currentStackMode = StackModeBrief
	}

	// The error stack mode is configured using command line arguments or environments.
	stackModeSetting := cmd.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackModeSettingMode := StackMode(stackModeSetting)
		switch stackModeSettingMode {
		case StackModeBrief, StackModeDetail:
			currentStackMode = stackModeSettingMode
		}
	}
}

// IsStackModeBrief returns whether the current error stack mode is in brief mode.
func IsStackModeBrief() bool {
	return currentStackMode == StackModeBrief
}
