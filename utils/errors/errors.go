package errors

import (
	"fmt"
	"net/http"
)

const (
	ErrCodeBadRequest                    = "BAD_REQUEST"
	ErrCodeUnauthorized                  = "UNAUTHORIZED"
	ErrCodeInsufficientAccess            = "INSUFFICIENT_ACCESS"
	ErrCodeNotFound                      = "NOT_FOUND"
	ErrCodePathNotFound                  = "PATH_NOT_FOUND"
	ErrCodeMethodNotAllowed              = "METHOD_NOT_ALLOWED"
	ErrCodeConflict                      = "CONFLICT"
	ErrCodeInternalServerError           = "INTERNAL_SERVER_ERROR"
	ErrCodeNotImplementedError           = "NOT_IMPLEMENTED"
	ErrCodeMissingMandatoryParameter     = "MISSING_MANDATORY_PARAMETER"
	ErrCodeMissingMandatoryConfiguration = "MISSING_MANDATORY_CONFIGURATION"
	ErrCodeConfigLoad                    = "CONFIG_LOAD_ERROR"
	ErrCodeServerStartupFailed           = "SERVER_STARTUP_FAILED"
	ErrCodeInvalidServerProtocol         = "SERVER_PROTOCOL_INVALID"
	ErrCodeInvalidServerLogLevel         = "SERVER_LOG_LEVEL_INVALID"
	ErrCodeInvalidPassword               = "PASSWORD_INVALID"
	ErrCodePasswordEncryptionError       = "PASSWORD_ENCRYPTION_FAILED"
	ErrCodePasswordDecryptionError       = "PASSWORD_DECRYPTION_FAILED"
	ErrCodeRegexCompileError             = "REGEX_COMPILE_ERROR"
	ErrCodeJSONMarshalError              = "JSON_MARSHAL_ERROR"
	ErrCodeJSONUnmarshalError            = "JSON_UNMARSHAL_ERROR"
	ErrCodeYAMLMarshalError              = "YAML_MARSHAL_ERROR"
	ErrCodeYAMLUnmarshalError            = "YAML_UNMARSHAL_ERROR"
	ErrCodeBase64DecodeError             = "BASE64_DECODE_ERROR"
	ErrCodeFileNotFound                  = "FILE_NOT_FOUND"
	ErrCodeFileCreateError               = "FILE_CREATE_ERROR"
	ErrCodeFileRemoveError               = "FILE_REMOVE_ERROR"
	ErrCodeFileOpenError                 = "FILE_OPEN_ERROR"
	ErrCodeFileReadError                 = "FILE_READ_ERROR"
	ErrCodeFileWriteError                = "FILE_WRITE_ERROR"
	ErrCodeBasicAuthMissing              = "BASIC_AUTH_MISSING"
	ErrCodeInvalidCredentials            = "CREDENTIALS_INVALID"
	ErrCodeInvalidPayload                = "PAYLOAD_INVALID"
	ErrCodeTraceIdMissing                = "TRACE_ID_MISSING"
	ErrCodeSubjectTokenTypeMissing       = "SUBJECT_TOKEN_TYPE_MISSING"
	ErrCodeSubjectTokenMissing           = "SUBJECT_TOKEN_MISSING"
	ErrCodeSubjectTokenTypeInvalid       = "SUBJECT_TOKEN_TYPE_INVALID"
	ErrCodeSubjectUnauthenticated        = "SUBJECT_UNAUTHENTICATED"
	ErrCodeSubjectNotAllowed             = "SUBJECT_NOT_ALLOWED"
)

var (
	ErrMsg = map[string]string{
		ErrCodeBadRequest:                    http.StatusText(http.StatusBadRequest),
		ErrCodeUnauthorized:                  http.StatusText(http.StatusUnauthorized),
		ErrCodeInsufficientAccess:            "Insufficient access",
		ErrCodeNotFound:                      http.StatusText(http.StatusNotFound),
		ErrCodePathNotFound:                  "Path Not Found",
		ErrCodeMethodNotAllowed:              http.StatusText(http.StatusMethodNotAllowed),
		ErrCodeConflict:                      http.StatusText(http.StatusConflict),
		ErrCodeInternalServerError:           "Unable to process the request due to an internal error. Please contact the system administrator",
		ErrCodeNotImplementedError:           "Not Implemented",
		ErrCodeMissingMandatoryParameter:     "Missing mandatory parameters : %v",
		ErrCodeMissingMandatoryConfiguration: "Missing mandatory configuration : %v",
		ErrCodeConfigLoad:                    "Error loading configuration: %s",
		ErrCodeServerStartupFailed:           "Server startup failed: %v",
		ErrCodeInvalidServerProtocol:         "Invalid server protocol '%s'",
		ErrCodeInvalidServerLogLevel:         "Invalid server log Level '%s'",
		ErrCodeInvalidPassword:               "Password should be at least 8 characters long with at least one number, one uppercase letter, one lowercase letter and one special character",
		ErrCodePasswordEncryptionError:       "Password encryption errors: %v",
		ErrCodePasswordDecryptionError:       "Password decryption errors: %v",
		ErrCodeRegexCompileError:             "Unable to compile regex : %v",
		ErrCodeJSONMarshalError:              "JSON marshal errors : %v",
		ErrCodeJSONUnmarshalError:            "JSON unmarshal errors : %v",
		ErrCodeYAMLMarshalError:              "YAML marshal errors : %v",
		ErrCodeYAMLUnmarshalError:            "YAML unmarshal errors : %v",
		ErrCodeFileNotFound:                  "File '%s' was not found",
		ErrCodeFileCreateError:               "Unable to create file '%s' : %v",
		ErrCodeFileRemoveError:               "Unable to remove file '%s' : %v",
		ErrCodeFileOpenError:                 "Unable to open file '%s' : %v",
		ErrCodeFileReadError:                 "Unable to read file '%s' : %v",
		ErrCodeFileWriteError:                "Unable to write to the file '%s' : %v",
		ErrCodeBasicAuthMissing:              "Basic authentication is required",
		ErrCodeInvalidCredentials:            "Invalid credentials",
		ErrCodeInvalidPayload:                "Payload is not valid",
		ErrCodeTraceIdMissing:                "Trace-Id header must be set",
		ErrCodeSubjectTokenTypeMissing:       "Subject-Token-Type header must be set",
		ErrCodeSubjectTokenMissing:           "Subject-Token header must be set",
		ErrCodeSubjectTokenTypeInvalid:       "Value of header Subject-Token-Type (%s) must be one of %v",
		ErrCodeSubjectUnauthenticated:        "Subject token is not valid",
		ErrCodeSubjectNotAllowed:             "Insufficient access",
	}
)

// Errors represents a list of Error.
type Errors struct {
	Errors []Error `json:"errors"`
}

// Error represents the error information.
type Error struct {
	Code    string `json:"code"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	TraceId string `json:"traceId"`
}

// New returns an Error
func New(code string, status int, message string) *Error {
	return &Error{
		Code:    code,
		Status:  status,
		Message: message,
		TraceId: "",
	}
}

// Newf formats according to a format specifier for the message and returns the Error.
func Newf(code string, status int, format string, a ...any) *Error {
	return &Error{
		Code:    code,
		Status:  status,
		Message: fmt.Sprintf(format, a...),
		TraceId: "",
	}
}

// BadRequestError returns a new bad request Error.
func BadRequestError(message string) *Error {
	return &Error{
		Code:    ErrCodeBadRequest,
		Status:  http.StatusBadRequest,
		Message: message,
		TraceId: "",
	}
}

// BadRequestErrorf formats according to a format specifier for the message and returns a bad request Error.
func BadRequestErrorf(format string, a ...any) *Error {
	return &Error{
		Code:    ErrCodeBadRequest,
		Status:  http.StatusBadRequest,
		Message: fmt.Sprintf(format, a...),
		TraceId: "",
	}
}

// UnauthorizedError returns a new unauthorized Error.
func UnauthorizedError(message string) *Error {
	return &Error{
		Code:    ErrCodeUnauthorized,
		Status:  http.StatusUnauthorized,
		Message: message,
		TraceId: "",
	}
}

// UnauthorizedErrorf returns a new unauthorized Error.
func UnauthorizedErrorf(format string, a ...any) *Error {
	return &Error{
		Code:    ErrCodeUnauthorized,
		Status:  http.StatusUnauthorized,
		Message: fmt.Sprintf(format, a...),
		TraceId: "",
	}
}

// ForbiddenError returns a new forbidden Error.
func ForbiddenError(message string) *Error {
	return &Error{
		Code:    ErrCodeInsufficientAccess,
		Status:  http.StatusForbidden,
		Message: message,
		TraceId: "",
	}
}

// ForbiddenErrorf returns a new forbidden Error.
func ForbiddenErrorf(format string, a ...any) *Error {
	return &Error{
		Code:    ErrCodeInsufficientAccess,
		Status:  http.StatusForbidden,
		Message: fmt.Sprintf(format, a...),
		TraceId: "",
	}
}

// NotFoundError returns a new not found Error.
func NotFoundError(message string) *Error {
	return &Error{
		Code:    ErrCodeNotFound,
		Status:  http.StatusNotFound,
		Message: message,
		TraceId: "",
	}
}

// NotFoundErrorf returns a new not found Error.
func NotFoundErrorf(format string, a ...any) *Error {
	return &Error{
		Code:    ErrCodeNotFound,
		Status:  http.StatusNotFound,
		Message: fmt.Sprintf(format, a...),
		TraceId: "",
	}
}

// ConflictError returns a new conflict Error.
func ConflictError(message string) *Error {
	return &Error{
		Code:    ErrCodeConflict,
		Status:  http.StatusConflict,
		Message: message,
		TraceId: "",
	}
}

// ConflictErrorf returns a new conflict Error.
func ConflictErrorf(format string, a ...any) *Error {
	return &Error{
		Code:    ErrCodeConflict,
		Status:  http.StatusConflict,
		Message: fmt.Sprintf(format, a...),
		TraceId: "",
	}
}

// InternalServerError returns a new internal server Error.
func InternalServerError(message string) *Error {
	return &Error{
		Code:    ErrCodeInternalServerError,
		Status:  http.StatusInternalServerError,
		Message: message,
		TraceId: "",
	}
}

// InternalServerErrorf returns a new internal server Error.
func InternalServerErrorf(format string, a ...any) *Error {
	return &Error{
		Code:    ErrCodeInternalServerError,
		Status:  http.StatusInternalServerError,
		Message: fmt.Sprintf(format, a...),
		TraceId: "",
	}
}
