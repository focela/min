// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package code provides universal error code definitions and implementations of common error codes.
package code

// Code is a universal error code interface definition.
type Code interface {
	// Code returns the integer number of the current error code.
	Code() int

	// Message returns the brief message for the current error code.
	Message() string

	// Detail returns the detailed information of the current error code,
	// which is mainly designed as an extension field for the error code.
	Detail() interface{}
}

// ================================================================================================================
// Common error code definitions.
// The framework reserves internal error codes: code < 1000.
// ================================================================================================================

var (
	CodeNil                       = localCode{-1, "", nil}                               // No error code specified.
	CodeOK                        = localCode{0, "OK", nil}                              // Everything is OK.
	CodeInternalError             = localCode{50, "Internal Error", nil}                 // An error occurred internally.
	CodeValidationFailed          = localCode{51, "Validation Failed", nil}              // Data validation failed.
	CodeDbOperationError          = localCode{52, "Database Operation Error", nil}       // Database operation error.
	CodeInvalidParameter          = localCode{53, "Invalid Parameter", nil}              // The given parameter for the current operation is invalid.
	CodeMissingParameter          = localCode{54, "Missing Parameter", nil}              // Parameter for the current operation is missing.
	CodeInvalidOperation          = localCode{55, "Invalid Operation", nil}              // The function cannot be used like this.
	CodeInvalidConfiguration      = localCode{56, "Invalid Configuration", nil}          // The configuration is invalid for the current operation.
	CodeMissingConfiguration      = localCode{57, "Missing Configuration", nil}          // The configuration is missing for the current operation.
	CodeNotImplemented            = localCode{58, "Not Implemented", nil}                // The operation is not implemented yet.
	CodeNotSupported              = localCode{59, "Not Supported", nil}                  // The operation is not supported yet.
	CodeOperationFailed           = localCode{60, "Operation Failed", nil}               // The operation failed.
	CodeNotAuthorized             = localCode{61, "Not Authorized", nil}                 // Not authorized.
	CodeSecurityReason            = localCode{62, "Security Reason", nil}                // Security reasons.
	CodeServerBusy                = localCode{63, "Server Is Busy", nil}                 // The server is busy, please try again later.
	CodeUnknown                   = localCode{64, "Unknown Error", nil}                  // Unknown error.
	CodeNotFound                  = localCode{65, "Not Found", nil}                      // Resource does not exist.
	CodeInvalidRequest            = localCode{66, "Invalid Request", nil}                // Invalid request.
	CodeNecessaryPackageNotImport = localCode{67, "Necessary Package Not Imported", nil} // A necessary package is not imported.
	CodeInternalPanic             = localCode{68, "Internal Panic", nil}                 // An internal panic occurred.
	CodeBusinessValidationFailed  = localCode{300, "Business Validation Failed", nil}    // Business validation failed.
)

// New creates and returns a new error code.
// It returns an interface object of type Code.
func New(code int, message string, detail interface{}) Code {
	return localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode creates and returns a new error code based on the given Code.
// The code and message are taken from the given `code`, but the detail is taken from the given `detail`.
func WithCode(code Code, detail interface{}) Code {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
