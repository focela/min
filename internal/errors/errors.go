// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package errors provides functionalities to manipulate errors for internal usage purposes.
package errors

import (
	"github.com/focela/min/internal/cmd"
)

// StackMode defines the mode for printing stack information in either brief or detailed mode.
type StackMode string

const (
	// commandEnvKeyForBrief is the environment variable for switching to brief error stack mode.
	// Deprecated: use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "min.error.brief"

	// commandEnvKeyForStackMode is the environment variable for switching between brief and detailed error stack modes.
	commandEnvKeyForStackMode = "min.error.stack.mode"
)

const (
	// StackModeBrief specifies that error stacks should not include framework-related stack traces.
	StackModeBrief StackMode = "brief"

	// StackModeDetail specifies that error stacks should include detailed stack traces, including framework-related ones.
	StackModeDetail StackMode = "detail"
)

var (
	// stackModeConfigured holds the currently configured stack mode.
	// The default is brief stack mode.
	stackModeConfigured = StackModeBrief
)

func init() {
	// Deprecated handling for brief stack mode.
	briefSetting := cmd.GetOptWithEnv(commandEnvKeyForBrief)
	if briefSetting == "1" || briefSetting == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Configure stack mode based on command-line arguments or environment variables.
	stackModeSetting := cmd.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		stackMode := StackMode(stackModeSetting)
		switch stackMode {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = stackMode
		}
	}
}

// IsStackModeBrief returns whether the current error stack mode is set to brief.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
