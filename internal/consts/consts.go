// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package consts defines constants that are shared all among packages of framework.
package consts

const (
	ConfigNodeNameDatabase        = "database"
	ConfigNodeNameLogger          = "logger"
	ConfigNodeNameRedis           = "redis"
	ConfigNodeNameViewer          = "viewer"
	ConfigNodeNameServer          = "server"     // General version configuration item name.
	ConfigNodeNameServerSecondary = "httpserver" // New version configuration item name support from v2.
	StackFilterKeyForRatcatcher   = "github.com/focela/ratcatcher/"
)
