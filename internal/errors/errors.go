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
	// commandEnvKeyForBrief is the environment variable name for enabling brief error stack mode.
	// Deprecated: use commandEnvKeyForStackMode instead.
	commandEnvKeyForBrief = "min.error.brief"

	// commandEnvKeyForStackMode is the environment variable name for setting stack mode (brief or detailed).
	commandEnvKeyForStackMode = "min.error.stack.mode"
)

const (
	// StackModeBrief prints only brief error stacks, excluding framework error stacks.
	StackModeBrief StackMode = "brief"

	// StackModeDetail prints detailed error stacks, including framework stacks.
	StackModeDetail StackMode = "detail"
)

var (
	// stackModeConfigured holds the currently configured stack mode.
	// Defaults to brief stack mode.
	stackModeConfigured = StackModeBrief
)

func init() {
	// Deprecated setting for brief mode.
	isBriefModeEnabled := cmd.GetOptWithEnv(commandEnvKeyForBrief)
	if isBriefModeEnabled == "1" || isBriefModeEnabled == "true" {
		stackModeConfigured = StackModeBrief
	}

	// Set stack mode based on command line arguments or environment variables.
	stackModeSetting := cmd.GetOptWithEnv(commandEnvKeyForStackMode)
	if stackModeSetting != "" {
		switch StackMode(stackModeSetting) {
		case StackModeBrief, StackModeDetail:
			stackModeConfigured = StackMode(stackModeSetting)
		}
	}
}

// IsStackModeBrief returns whether the current stack mode is set to brief.
func IsStackModeBrief() bool {
	return stackModeConfigured == StackModeBrief
}
