// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package code provides universal error code definitions and common error code implementations.
package code

// Code is the universal error code interface definition.
type Code interface {
	// Code returns the integer value of the current error code.
	Code() int

	// Message returns a brief message for the current error code.
	Message() string

	// Detail returns detailed information of the current error code,
	// which is mainly designed as an extension field for the error code.
	Detail() interface{}
}

// ================================================================================================================
// Common error code definitions.
// The framework reserves internal error codes: code < 1000.
// ================================================================================================================

var (
	CodeNil                       = localCode{-1, "", nil}                             // No error code specified.
	CodeOK                        = localCode{0, "OK", nil}                            // Operation was successful.
	CodeInternalError             = localCode{50, "Internal Error", nil}               // An internal error occurred.
	CodeValidationFailed          = localCode{51, "Validation Failed", nil}            // Data validation failed.
	CodeDbOperationError          = localCode{52, "Database Operation Error", nil}     // A database operation error occurred.
	CodeInvalidParameter          = localCode{53, "Invalid Parameter", nil}            // The given parameter for the current operation is invalid.
	CodeMissingParameter          = localCode{54, "Missing Parameter", nil}            // A required parameter is missing for the current operation.
	CodeInvalidOperation          = localCode{55, "Invalid Operation", nil}            // The operation is invalid.
	CodeInvalidConfiguration      = localCode{56, "Invalid Configuration", nil}        // The configuration is invalid for the current operation.
	CodeMissingConfiguration      = localCode{57, "Missing Configuration", nil}        // The configuration is missing for the current operation.
	CodeNotImplemented            = localCode{58, "Not Implemented", nil}              // The operation is not implemented yet.
	CodeNotSupported              = localCode{59, "Not Supported", nil}                // The operation is not supported.
	CodeOperationFailed           = localCode{60, "Operation Failed", nil}             // The operation failed.
	CodeNotAuthorized             = localCode{61, "Not Authorized", nil}               // Not authorized to perform the operation.
	CodeSecurityReason            = localCode{62, "Security Reason", nil}              // The operation is prohibited for security reasons.
	CodeServerBusy                = localCode{63, "Server Is Busy", nil}               // The server is busy, please try again later.
	CodeUnknown                   = localCode{64, "Unknown Error", nil}                // An unknown error occurred.
	CodeNotFound                  = localCode{65, "Not Found", nil}                    // The requested resource does not exist.
	CodeInvalidRequest            = localCode{66, "Invalid Request", nil}              // The request is invalid.
	CodeNecessaryPackageNotImport = localCode{67, "Necessary Package Not Import", nil} // A necessary package was not imported.
	CodeInternalPanic             = localCode{68, "Internal Panic", nil}               // An internal panic occurred.
	CodeBusinessValidationFailed  = localCode{300, "Business Validation Failed", nil}  // Business validation failed.
)

// New creates and returns a new error code.
// Note that it returns an interface object of Code.
func New(code int, message string, detail interface{}) Code {
	return localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode creates and returns a new error code based on the given Code.
// The code and message are from the given `code`, but the detail is from the given `detail`.
func WithCode(code Code, detail interface{}) Code {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
