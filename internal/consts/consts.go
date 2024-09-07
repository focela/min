// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package consts defines constants shared across all framework packages.
package consts

const (
	NodeDatabase        = "database"               // Database configuration node
	NodeLogger          = "logger"                 // Logger configuration node
	NodeRedis           = "redis"                  // Redis configuration node
	NodeViewer          = "viewer"                 // Viewer configuration node
	NodeServer          = "server"                 // Server configuration node
	NodeServerSecondary = "httpserver"             // Secondary server configuration node
	StackFilterKey      = "github.com/focela/min/" // Stack filter key for Min
)
