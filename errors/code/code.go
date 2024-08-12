// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package code provides universal error code definitions and common error code implementations.
package code

// Code is a universal error code interface definition.
type Code interface {
	// Code returns the integer value of the current error code.
	Code() int

	// Message returns a brief message for the current error code.
	Message() string

	// Detail returns detailed information for the current error code,
	// primarily designed as an extension field for the error code.
	Detail() interface{}
}

// ================================================================================================================
// Common error code definitions.
// Error codes reserved by the framework: code < 1000.
// ================================================================================================================

var (
	CodeNil                       = errorCode{-1, "", nil}                             // No error code specified.
	CodeOK                        = errorCode{0, "OK", nil}                            // Operation completed successfully.
	CodeInternalError             = errorCode{50, "Internal Error", nil}               // An internal error occurred.
	CodeValidationFailed          = errorCode{51, "Validation Failed", nil}            // Data validation failed.
	CodeDbOperationError          = errorCode{52, "Database Operation Error", nil}     // Database operation error.
	CodeInvalidParameter          = errorCode{53, "Invalid Parameter", nil}            // The provided parameter is invalid.
	CodeMissingParameter          = errorCode{54, "Missing Parameter", nil}            // A required parameter is missing.
	CodeInvalidOperation          = errorCode{55, "Invalid Operation", nil}            // The operation is not valid.
	CodeInvalidConfiguration      = errorCode{56, "Invalid Configuration", nil}        // The configuration is invalid.
	CodeMissingConfiguration      = errorCode{57, "Missing Configuration", nil}        // A required configuration is missing.
	CodeNotImplemented            = errorCode{58, "Not Implemented", nil}              // The operation is not yet implemented.
	CodeNotSupported              = errorCode{59, "Not Supported", nil}                // The operation is not supported.
	CodeOperationFailed           = errorCode{60, "Operation Failed", nil}             // The operation failed.
	CodeNotAuthorized             = errorCode{61, "Not Authorized", nil}               // Authorization is required.
	CodeSecurityReason            = errorCode{62, "Security Reason", nil}              // The operation is prohibited for security reasons.
	CodeServerBusy                = errorCode{63, "Server Is Busy", nil}               // The server is busy. Please try again later.
	CodeUnknown                   = errorCode{64, "Unknown Error", nil}                // An unknown error occurred.
	CodeNotFound                  = errorCode{65, "Not Found", nil}                    // The requested resource was not found.
	CodeInvalidRequest            = errorCode{66, "Invalid Request", nil}              // The request is invalid.
	CodeNecessaryPackageNotImport = errorCode{67, "Necessary Package Not Import", nil} // A necessary package has not been imported.
	CodeInternalPanic             = errorCode{68, "Internal Panic", nil}               // An internal panic occurred.
	CodeBusinessValidationFailed  = errorCode{300, "Business Validation Failed", nil}  // Business validation failed.
)

// New creates and returns an error code.
// It returns an interface object of Code.
func New(code int, message string, detail interface{}) Code {
	return errorCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode creates and returns a new error code based on the given Code.
// The code and message are from the given `code`, but the detail is from the given `detail`.
func WithCode(code Code, detail interface{}) Code {
	return errorCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
