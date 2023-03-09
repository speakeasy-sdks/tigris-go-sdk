package shared

type ErrorCodeEnum string

const (
	ErrorCodeEnumOk                 ErrorCodeEnum = "OK"
	ErrorCodeEnumCancelled          ErrorCodeEnum = "CANCELLED"
	ErrorCodeEnumUnknown            ErrorCodeEnum = "UNKNOWN"
	ErrorCodeEnumInvalidArgument    ErrorCodeEnum = "INVALID_ARGUMENT"
	ErrorCodeEnumDeadlineExceeded   ErrorCodeEnum = "DEADLINE_EXCEEDED"
	ErrorCodeEnumNotFound           ErrorCodeEnum = "NOT_FOUND"
	ErrorCodeEnumAlreadyExists      ErrorCodeEnum = "ALREADY_EXISTS"
	ErrorCodeEnumPermissionDenied   ErrorCodeEnum = "PERMISSION_DENIED"
	ErrorCodeEnumResourceExhausted  ErrorCodeEnum = "RESOURCE_EXHAUSTED"
	ErrorCodeEnumFailedPrecondition ErrorCodeEnum = "FAILED_PRECONDITION"
	ErrorCodeEnumAborted            ErrorCodeEnum = "ABORTED"
	ErrorCodeEnumOutOfRange         ErrorCodeEnum = "OUT_OF_RANGE"
	ErrorCodeEnumUnimplemented      ErrorCodeEnum = "UNIMPLEMENTED"
	ErrorCodeEnumInternal           ErrorCodeEnum = "INTERNAL"
	ErrorCodeEnumUnavailable        ErrorCodeEnum = "UNAVAILABLE"
	ErrorCodeEnumDataLoss           ErrorCodeEnum = "DATA_LOSS"
	ErrorCodeEnumUnauthenticated    ErrorCodeEnum = "UNAUTHENTICATED"
	ErrorCodeEnumConflict           ErrorCodeEnum = "CONFLICT"
	ErrorCodeEnumBadGateway         ErrorCodeEnum = "BAD_GATEWAY"
	ErrorCodeEnumMethodNotAllowed   ErrorCodeEnum = "METHOD_NOT_ALLOWED"
)

// Error
// The Error type defines a logical error model
type Error struct {
	Code    *ErrorCodeEnum `json:"code,omitempty"`
	Message *string        `json:"message,omitempty"`
}
