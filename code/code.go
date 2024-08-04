// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package code provides universal error code definition and common error codes implementation.
package code

// ErrorCode is the universal error code interface definition.
type ErrorCode interface {
	// Code returns the integer number of the current error code.
	Code() int

	// Message returns the brief message for the current error code.
	Message() string

	// Detail returns the detailed information of the current error code,
	// which is mainly designed as an extension field for error codes.
	Detail() interface{}
}

// ================================================================================================================
// Common error code definitions.
// There are reserved internal error codes by the framework: code < 1000.
// ================================================================================================================

var (
	CODE_NIL                          = localCode{-1, "", nil}                             // No error code specified.
	CODE_OK                           = localCode{0, "OK", nil}                            // It is OK.
	CODE_INTERNAL_ERROR               = localCode{50, "Internal Error", nil}               // An error occurred internally.
	CODE_VALIDATION_FAILED            = localCode{51, "Validation Failed", nil}            // Data validation failed.
	CODE_DB_OPERATION_ERROR           = localCode{52, "Database Operation Error", nil}     // Database operation error.
	CODE_INVALID_PARAMETER            = localCode{53, "Invalid Parameter", nil}            // The given parameter for the current operation is invalid.
	CODE_MISSING_PARAMETER            = localCode{54, "Missing Parameter", nil}            // Parameter for the current operation is missing.
	CODE_INVALID_OPERATION            = localCode{55, "Invalid Operation", nil}            // The function cannot be used like this.
	CODE_INVALID_CONFIGURATION        = localCode{56, "Invalid Configuration", nil}        // The configuration is invalid for the current operation.
	CODE_MISSING_CONFIGURATION        = localCode{57, "Missing Configuration", nil}        // The configuration is missing for the current operation.
	CODE_NOT_IMPLEMENTED              = localCode{58, "Not Implemented", nil}              // The operation is not implemented yet.
	CODE_NOT_SUPPORTED                = localCode{59, "Not Supported", nil}                // The operation is not supported yet.
	CODE_OPERATION_FAILED             = localCode{60, "Operation Failed", nil}             // I tried, but I cannot give you what you want.
	CODE_NOT_AUTHORIZED               = localCode{61, "Not Authorized", nil}               // Not Authorized.
	CODE_SECURITY_REASON              = localCode{62, "Security Reason", nil}              // Security Reason.
	CODE_SERVER_BUSY                  = localCode{63, "Server Is Busy", nil}               // Server is busy, please try again later.
	CODE_UNKNOWN                      = localCode{64, "Unknown Error", nil}                // Unknown error.
	CODE_NOT_FOUND                    = localCode{65, "Not Found", nil}                    // Resource does not exist.
	CODE_INVALID_REQUEST              = localCode{66, "Invalid Request", nil}              // Invalid request.
	CODE_NECESSARY_PACKAGE_NOT_IMPORT = localCode{67, "Necessary Package Not Import", nil} // It needs necessary package import.
	CODE_INTERNAL_PANIC               = localCode{68, "Internal Panic", nil}               // A panic occurred internally.
	CODE_BUSINESS_VALIDATION_FAILED   = localCode{300, "Business Validation Failed", nil}  // Business validation failed.
)

// New creates and returns an error code.
// Note that it returns an interface object of ErrorCode.
func New(code int, message string, detail interface{}) ErrorCode {
	return localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode creates and returns a new error code based on the given ErrorCode.
// The code and message are from the given `code`, but the detail is from the given `detail`.
func WithCode(code ErrorCode, detail interface{}) ErrorCode {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
