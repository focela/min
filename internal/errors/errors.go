// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package errors provides functionalities to manipulate errors for internal usage.
package errors

import (
	"github.com/focela/min/internal/cmd"
)

// StackMode defines the mode for printing stack information: brief or detailed.
type StackMode string

const (
	// Deprecated: use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "min.error.brief"

	// commandEnvKeyForStackMode is the environment variable for setting the error stack mode.
	commandEnvKeyForStackMode = "min.error.stack.mode"
)

const (
	// StackModeBrief specifies printing of error stacks without framework details.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies printing detailed error stacks including framework details.
	StackModeDetail StackMode = "detail"
)

var (
	// stackModeConfigured holds the current error stack mode configuration.
	// The default is brief stack mode.
	stackModeConfigured = StackModeBrief
)

func init() {
	// Deprecated.
	briefSetting := cmd.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Configure error stack mode using command line arguments or environment variables.
	stackModeSetting := cmd.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackMode := StackMode(stackModeSetting)
		switch stackMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = stackMode
		}
	}
}

// IsStackModeBrief checks if the current error stack mode is brief.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
