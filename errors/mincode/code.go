// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package mincode provides universal error code definition and common error codes implements.
package mincode

// Code is universal error code interface definition.
type Code interface {
	// Code returns the integer representation of the error code.
	Code() int

	// Message returns the brief message for the error code.
	Message() string

	// Detail returns detailed information of the error code.
	Detail() interface{}
}

// ================================================================================================================
// Common error code definition.
// There are reserved internal error codes by framework: code < 1000.
// ================================================================================================================

var (
	CodeNil                       = errorCode{-1, "", nil}                               // No error code specified.
	CodeOK                        = errorCode{0, "OK", nil}                              // It is OK.
	CodeInternalError             = errorCode{50, "Internal Error", nil}                 // An internal error occurred.
	CodeValidationFailed          = errorCode{51, "Validation Failed", nil}              // Data validation failed.
	CodeDbOperationError          = errorCode{52, "Database Operation Error", nil}       // Database operation error.
	CodeInvalidParameter          = errorCode{53, "Invalid Parameter", nil}              // The given parameter is invalid.
	CodeMissingParameter          = errorCode{54, "Missing Parameter", nil}              // Parameter is missing.
	CodeInvalidOperation          = errorCode{55, "Invalid Operation", nil}              // The function cannot be used this way.
	CodeInvalidConfiguration      = errorCode{56, "Invalid Configuration", nil}          // Invalid configuration.
	CodeMissingConfiguration      = errorCode{57, "Missing Configuration", nil}          // Missing configuration.
	CodeNotImplemented            = errorCode{58, "Not Implemented", nil}                // The operation is not implemented.
	CodeNotSupported              = errorCode{59, "Not Supported", nil}                  // The operation is not supported.
	CodeOperationFailed           = errorCode{60, "Operation Failed", nil}               // Operation failed.
	CodeNotAuthorized             = errorCode{61, "Not Authorized", nil}                 // Not authorized.
	CodeSecurityReason            = errorCode{62, "Security Reason", nil}                // Security reason.
	CodeServerBusy                = errorCode{63, "Server Is Busy", nil}                 // Server is busy.
	CodeUnknown                   = errorCode{64, "Unknown Error", nil}                  // Unknown error.
	CodeNotFound                  = errorCode{65, "Not Found", nil}                      // Resource not found.
	CodeInvalidRequest            = errorCode{66, "Invalid Request", nil}                // Invalid request.
	CodeNecessaryPackageNotImport = errorCode{67, "Necessary Package Not Imported", nil} // Necessary package not imported.
	CodeInternalPanic             = errorCode{68, "Internal Panic", nil}                 // Internal panic.
	CodeBusinessValidationFailed  = errorCode{300, "Business Validation Failed", nil}    // Business validation failed.
)

// New creates and returns an error code.
// It returns a Code interface object.
func New(code int, message string, detail interface{}) Code {
	return errorCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode creates and returns a new error code based on a given Code.
// The code and message are from the given code, but the detail is from the given detail.
func WithCode(code Code, detail interface{}) Code {
	return errorCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
