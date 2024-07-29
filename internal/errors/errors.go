// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package errors provides functionalities to manipulate errors for internal usage purpose.
package errors

import "github.com/focela/ratcatcher/internal/cmd"

// StackMode is the mode that printing stack information in StackModeBrief or StackModeDetail mode.
type StackMode string

const (
	// commandEnvKeyForBrief is the command environment name for switch key for brief error stack.
	// Deprecated: use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "r.rerror.brief"

	// commandEnvKeyForStackMode is the command environment name for switch key for brief error stack.
	commandEnvKeyForStackMode = "r.rerror.stack.mode"
)

const (
	// StackModeBrief specifies all error stacks printing no framework error stacks.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies all error stacks printing detailed error stacks including framework stacks.
	StackModeDetail StackMode = "detail"
)

var (
	// stackModeConfigured is the configured error stack mode variable.
	// It is brief stack mode in default.
	stackModeConfigured = StackModeBrief
)

func init() {
	// Deprecated.
	briefSetting := cmd.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// The error stack mode is configured using command line arguments or environments.
	stackModeSetting := cmd.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackModeSettingMode := StackMode(stackModeSetting)
		switch stackModeSettingMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = stackModeSettingMode
		}
	}
}

// IsStackModeBrief returns whether current error stack mode is in brief mode.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
