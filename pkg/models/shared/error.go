// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Code - The status code is a short, machine parsable string, which uniquely identifies the error type. Tigris to HTTP code mapping [here](/reference/http-code)
type Code string

const (
	CodeOk                 Code = "OK"
	CodeCancelled          Code = "CANCELLED"
	CodeUnknown            Code = "UNKNOWN"
	CodeInvalidArgument    Code = "INVALID_ARGUMENT"
	CodeDeadlineExceeded   Code = "DEADLINE_EXCEEDED"
	CodeNotFound           Code = "NOT_FOUND"
	CodeAlreadyExists      Code = "ALREADY_EXISTS"
	CodePermissionDenied   Code = "PERMISSION_DENIED"
	CodeResourceExhausted  Code = "RESOURCE_EXHAUSTED"
	CodeFailedPrecondition Code = "FAILED_PRECONDITION"
	CodeAborted            Code = "ABORTED"
	CodeOutOfRange         Code = "OUT_OF_RANGE"
	CodeUnimplemented      Code = "UNIMPLEMENTED"
	CodeInternal           Code = "INTERNAL"
	CodeUnavailable        Code = "UNAVAILABLE"
	CodeDataLoss           Code = "DATA_LOSS"
	CodeUnauthenticated    Code = "UNAUTHENTICATED"
	CodeConflict           Code = "CONFLICT"
	CodeBadGateway         Code = "BAD_GATEWAY"
	CodeMethodNotAllowed   Code = "METHOD_NOT_ALLOWED"
)

func (e Code) ToPointer() *Code {
	return &e
}

func (e *Code) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OK":
		fallthrough
	case "CANCELLED":
		fallthrough
	case "UNKNOWN":
		fallthrough
	case "INVALID_ARGUMENT":
		fallthrough
	case "DEADLINE_EXCEEDED":
		fallthrough
	case "NOT_FOUND":
		fallthrough
	case "ALREADY_EXISTS":
		fallthrough
	case "PERMISSION_DENIED":
		fallthrough
	case "RESOURCE_EXHAUSTED":
		fallthrough
	case "FAILED_PRECONDITION":
		fallthrough
	case "ABORTED":
		fallthrough
	case "OUT_OF_RANGE":
		fallthrough
	case "UNIMPLEMENTED":
		fallthrough
	case "INTERNAL":
		fallthrough
	case "UNAVAILABLE":
		fallthrough
	case "DATA_LOSS":
		fallthrough
	case "UNAUTHENTICATED":
		fallthrough
	case "CONFLICT":
		fallthrough
	case "BAD_GATEWAY":
		fallthrough
	case "METHOD_NOT_ALLOWED":
		*e = Code(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Code: %v", v)
	}
}

// The Error type defines a logical error model
type Error struct {
	// The status code is a short, machine parsable string, which uniquely identifies the error type. Tigris to HTTP code mapping [here](/reference/http-code)
	Code *Code `json:"code,omitempty"`
	// A developer-facing descriptive error message
	Message *string `json:"message,omitempty"`
}

func (o *Error) GetCode() *Code {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *Error) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}
