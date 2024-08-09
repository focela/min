// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package code

// Code is universal error code interface definition.
type Code interface {
	// Code returns the integer number of current error code.
	Code() int

	// Message returns the brief message for current error code.
	Message() string

	// Detail returns the detailed information of current error code,
	// which is mainly designed as an extension field for error code.
	Detail() interface{}
}

type errorCode struct {
	code    int
	message string
	detail  interface{}
}

func (e errorCode) Code() int {
	return e.code
}

func (e errorCode) Message() string {
	return e.message
}

func (e errorCode) Detail() interface{} {
	return e.detail
}

// ================================================================================================================
// Common error code definition.
// There are reserved internal error codes by the framework: code < 1000.
// ================================================================================================================

const (
	CodeNilCode                       = -1
	CodeOKCode                        = 0
	CodeInternalErrorCode             = 50
	CodeValidationFailedCode          = 51
	CodeDbOperationErrorCode          = 52
	CodeInvalidParameterCode          = 53
	CodeMissingParameterCode          = 54
	CodeInvalidOperationCode          = 55
	CodeInvalidConfigurationCode      = 56
	CodeMissingConfigurationCode      = 57
	CodeNotImplementedCode            = 58
	CodeNotSupportedCode              = 59
	CodeOperationFailedCode           = 60
	CodeNotAuthorizedCode             = 61
	CodeSecurityReasonCode            = 62
	CodeServerBusyCode                = 63
	CodeUnknownCode                   = 64
	CodeNotFoundCode                  = 65
	CodeInvalidRequestCode            = 66
	CodeNecessaryPackageNotImportCode = 67
	CodeInternalPanicCode             = 68
	CodeBusinessValidationFailedCode  = 300
)

var (
	CodeNil                       = errorCode{CodeNilCode, "", nil}                                                   // No error code specified.
	CodeOK                        = errorCode{CodeOKCode, "OK", nil}                                                  // It is OK.
	CodeInternalError             = errorCode{CodeInternalErrorCode, "Internal Error", nil}                           // An error occurred internally.
	CodeValidationFailed          = errorCode{CodeValidationFailedCode, "Validation Failed", nil}                     // Data validation failed.
	CodeDbOperationError          = errorCode{CodeDbOperationErrorCode, "Database Operation Error", nil}              // Database operation error.
	CodeInvalidParameter          = errorCode{CodeInvalidParameterCode, "Invalid Parameter", nil}                     // The given parameter for current operation is invalid.
	CodeMissingParameter          = errorCode{CodeMissingParameterCode, "Missing Parameter", nil}                     // Parameter for current operation is missing.
	CodeInvalidOperation          = errorCode{CodeInvalidOperationCode, "Invalid Operation", nil}                     // The function cannot be used like this.
	CodeInvalidConfiguration      = errorCode{CodeInvalidConfigurationCode, "Invalid Configuration", nil}             // The configuration is invalid for current operation.
	CodeMissingConfiguration      = errorCode{CodeMissingConfigurationCode, "Missing Configuration", nil}             // The configuration is missing for current operation.
	CodeNotImplemented            = errorCode{CodeNotImplementedCode, "Not Implemented", nil}                         // The operation is not implemented yet.
	CodeNotSupported              = errorCode{CodeNotSupportedCode, "Not Supported", nil}                             // The operation is not supported yet.
	CodeOperationFailed           = errorCode{CodeOperationFailedCode, "Operation Failed", nil}                       // I tried, but I cannot give you what you want.
	CodeNotAuthorized             = errorCode{CodeNotAuthorizedCode, "Not Authorized", nil}                           // Not Authorized.
	CodeSecurityReason            = errorCode{CodeSecurityReasonCode, "Security Reason", nil}                         // Security Reason.
	CodeServerBusy                = errorCode{CodeServerBusyCode, "Server Is Busy", nil}                              // Server is busy, please try again later.
	CodeUnknown                   = errorCode{CodeUnknownCode, "Unknown Error", nil}                                  // Unknown error.
	CodeNotFound                  = errorCode{CodeNotFoundCode, "Not Found", nil}                                     // Resource does not exist.
	CodeInvalidRequest            = errorCode{CodeInvalidRequestCode, "Invalid Request", nil}                         // Invalid request.
	CodeNecessaryPackageNotImport = errorCode{CodeNecessaryPackageNotImportCode, "Necessary Package Not Import", nil} // It needs necessary package import.
	CodeInternalPanic             = errorCode{CodeInternalPanicCode, "A panic occurred internally.", nil}             // A panic occurred internally.
	CodeBusinessValidationFailed  = errorCode{CodeBusinessValidationFailedCode, "Business Validation Failed", nil}    // Business validation failed.
)

// New creates and returns an error code.
// Note that it returns an interface object of Code.
func New(code int, message string, detail interface{}) Code {
	return errorCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode creates and returns a new error code based on given Code.
// The code and message is from given `code`, but the detail is from given `detail`.
func WithCode(code Code, detail interface{}) Code {
	return errorCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
