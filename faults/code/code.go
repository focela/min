// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package code provides universal error code definition and common error code implementations.
package code

// Code defines a universal error code interface.
type Code interface {
	// Code returns the current error code as an integer.
	Code() int

	// Message returns a brief message associated with the current error code.
	Message() string

	// Detail returns detailed information about the current error code.
	// This is mainly designed as an extension field for the error code.
	Detail() interface{}
}

// ================================================================================================================
// Common error code definitions.
// The framework reserves internal error codes where code < 1000.
// ================================================================================================================

var (
	CodeNil                       = &localCode{-1, "", nil}                             // No error code specified.
	CodeOK                        = &localCode{0, "OK", nil}                            // It is OK.
	CodeInternalError             = &localCode{50, "Internal Error", nil}               // An error occurred internally.
	CodeValidationFailed          = &localCode{51, "Validation Failed", nil}            // Data validation failed.
	CodeDbOperationError          = &localCode{52, "Database Operation Error", nil}     // Database operation error.
	CodeInvalidParameter          = &localCode{53, "Invalid Parameter", nil}            // The given parameter for current operation is invalid.
	CodeMissingParameter          = &localCode{54, "Missing Parameter", nil}            // Parameter for current operation is missing.
	CodeInvalidOperation          = &localCode{55, "Invalid Operation", nil}            // The function cannot be used like this.
	CodeInvalidConfiguration      = &localCode{56, "Invalid Configuration", nil}        // The configuration is invalid for current operation.
	CodeMissingConfiguration      = &localCode{57, "Missing Configuration", nil}        // The configuration is missing for current operation.
	CodeNotImplemented            = &localCode{58, "Not Implemented", nil}              // The operation is not implemented yet.
	CodeNotSupported              = &localCode{59, "Not Supported", nil}                // The operation is not supported yet.
	CodeOperationFailed           = &localCode{60, "Operation Failed", nil}             // I tried, but I cannot give you what you want.
	CodeNotAuthorized             = &localCode{61, "Not Authorized", nil}               // Not Authorized.
	CodeSecurityReason            = &localCode{62, "Security Reason", nil}              // Security Reason.
	CodeServerBusy                = &localCode{63, "Server Is Busy", nil}               // Server is busy, please try again later.
	CodeUnknown                   = &localCode{64, "Unknown Error", nil}                // Unknown error.
	CodeNotFound                  = &localCode{65, "Not Found", nil}                    // Resource does not exist.
	CodeInvalidRequest            = &localCode{66, "Invalid Request", nil}              // Invalid request.
	CodeNecessaryPackageNotImport = &localCode{67, "Necessary Package Not Import", nil} // It needs necessary package import.
	CodeInternalPanic             = &localCode{68, "Internal Panic", nil}               // An internal panic occurred.
	CodeBusinessValidationFailed  = &localCode{300, "Business Validation Failed", nil}  // Business validation failed.
)

// New creates and returns a new error code.
// Note that it returns an interface object of Code.
func New(code int, message string, detail interface{}) Code {
	return &localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode creates and returns a new error code based on the given Code.
// The code and message are from the given `code`, but the detail is from the given `detail`.
func WithCode(code Code, detail interface{}) Code {
	return &localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
