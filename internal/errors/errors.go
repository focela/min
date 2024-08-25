// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package errors provides functionalities to manipulate errors for internal usage purpose.
package errors

import (
	"strings"

	"github.com/focela/min/internal/cmd"
)

// StackMode is the mode that controls the printing of stack information
// in StackModeBrief or StackModeDetail mode.
type StackMode string

const (
	// commandEnvKeyForBrief is the command environment name for switching to brief error stack.
	// Deprecated: Use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "min.error.brief"

	// commandEnvKeyForStackMode is the command environment name for controlling error stack mode.
	commandEnvKeyForStackMode = "min.error.stack.mode"
)

const (
	// StackModeBrief specifies that error stacks are printed without framework error stacks.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies that all error stacks, including framework stacks, are printed.
	StackModeDetail StackMode = "detail"
)

var (
	// stackModeConfigured is the configured error stack mode variable.
	// It is set to brief stack mode by default.
	stackModeConfigured = StackModeBrief
)

func init() {
	// Deprecated: Checking for brief mode using old environment key.
	briefSetting := cmd.GetOptWithEnv(commandEnvKeyForBrief)
	if strings.EqualFold(briefSetting, "1") || strings.EqualFold(briefSetting, "true") {
		stackModeConfigured = StackModeBrief
	}

	// The error stack mode is configured using command line arguments or environment variables.
	stackModeSetting := cmd.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackModeSettingMode := StackMode(stackModeSetting)
		if stackModeSettingMode == StackModeBrief || stackModeSettingMode == StackModeDetail {
			stackModeConfigured = stackModeSettingMode
		}
	}
}

// IsStackModeBrief returns whether the current error stack mode is set to brief mode.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
