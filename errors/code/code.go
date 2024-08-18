// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package code provides universal error code definitions and common error code implementations.
package code

// Code defines the universal error code interface.
type Code interface {
	// Code returns the integer value of the current error code.
	Code() int

	// Message returns a brief message for the current error code.
	Message() string

	// Detail returns detailed information about the current error code,
	// primarily designed as an extension field for the error code.
	Detail() interface{}
}

// ================================================================================================================
// Common error code definitions.
// These internal error codes are reserved by the framework: code < 1000.
// ================================================================================================================

var (
	CodeNil                       = localCode{-1, "", nil}                             // No error code specified.
	CodeOK                        = localCode{0, "OK", nil}                            // Operation was successful.
	CodeInternalError             = localCode{50, "Internal Error", nil}               // An internal error occurred.
	CodeValidationFailed          = localCode{51, "Validation Failed", nil}            // Data validation failed.
	CodeDbOperationError          = localCode{52, "Database Operation Error", nil}     // Database operation error.
	CodeInvalidParameter          = localCode{53, "Invalid Parameter", nil}            // Invalid parameter for the current operation.
	CodeMissingParameter          = localCode{54, "Missing Parameter", nil}            // Parameter for the current operation is missing.
	CodeInvalidOperation          = localCode{55, "Invalid Operation", nil}            // Operation cannot be performed in this way.
	CodeInvalidConfiguration      = localCode{56, "Invalid Configuration", nil}        // Invalid configuration for the current operation.
	CodeMissingConfiguration      = localCode{57, "Missing Configuration", nil}        // Configuration is missing for the current operation.
	CodeNotImplemented            = localCode{58, "Not Implemented", nil}              // Operation not yet implemented.
	CodeNotSupported              = localCode{59, "Not Supported", nil}                // Operation not yet supported.
	CodeOperationFailed           = localCode{60, "Operation Failed", nil}             // The operation failed to complete successfully.
	CodeNotAuthorized             = localCode{61, "Not Authorized", nil}               // User is not authorized.
	CodeSecurityReason            = localCode{62, "Security Reason", nil}              // Security reason prevented operation.
	CodeServerBusy                = localCode{63, "Server Is Busy", nil}               // Server is busy, please try again later.
	CodeUnknown                   = localCode{64, "Unknown Error", nil}                // Unknown error occurred.
	CodeNotFound                  = localCode{65, "Not Found", nil}                    // Resource not found.
	CodeInvalidRequest            = localCode{66, "Invalid Request", nil}              // The request is invalid.
	CodeNecessaryPackageNotImport = localCode{67, "Necessary Package Not Import", nil} // Required package is not imported.
	CodeInternalPanic             = localCode{68, "Internal Panic", nil}               // An internal panic occurred.
	CodeBusinessValidationFailed  = localCode{300, "Business Validation Failed", nil}  // Business validation failed.
)

// New creates and returns a new error code.
// Note that it returns an interface object of type Code.
func New(code int, message string, detail interface{}) Code {
	return localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode creates and returns a new error code based on the given Code.
// The code and message are taken from the provided `code`, but the detail is from the provided `detail`.
func WithCode(code Code, detail interface{}) Code {
	if detail == nil {
		return code
	}
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
