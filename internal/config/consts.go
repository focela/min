// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package config defines constants that are shared among all packages of the framework.
package config

const (
	NodeDatabase        = "database"                 // NodeDatabase represents the database configuration node.
	NodeLogger          = "logger"                   // NodeLogger represents the logger configuration node.
	NodeRedis           = "redis"                    // NodeRedis represents the Redis configuration node.
	NodeViewer          = "viewer"                   // NodeViewer represents the viewer configuration node.
	NodeServer          = "server"                   // NodeServer represents the server configuration node.
	NodeServerSecondary = "httpserver"               // NodeServerSecondary represents the secondary server configuration node.
	StackFilterKey      = "github.com/focela/plume/" // StackFilterKey is the stack filter key for the GoPlume framework.

)
