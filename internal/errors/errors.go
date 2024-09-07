// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package errors provides functionalities to manipulate errors for internal usage purpose.
package errors

import (
	"github.com/focela/min/internal/command"
)

// StackMode is the mode that printing stack information in StackModeBrief or StackModeDetail mode.
type StackMode string

const (
	// commandEnvKeyForBrief is the command environment name for switch key for brief error stack.
	// Deprecated: use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "min.error.brief"

	// commandEnvKeyForStackMode is the command environment name for switching error stack modes (brief or detail).
	commandEnvKeyForStackMode = "min.error.stack.mode"
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
	// Check and set the brief mode from command or environment variables.
	if briefSetting := command.GetOptionWithEnv(commandEnvKeyForBrief); briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Set the stack mode based on command line arguments or environment variables.
	if stackModeSetting := command.GetOptionWithEnv(commandEnvKeyForStackMode); stackModeSetting != "" {
		switch StackMode(stackModeSetting) {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = StackMode(stackModeSetting)
		}
	}
}

// IsStackModeBrief checks if the current error stack mode is set to brief mode.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
