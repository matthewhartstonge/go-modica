package modica

import "errors"

const (
	// constant error codes as returned by the API
	errCodeSendFailed     = "send_failed"
	errCodeInvalidJson    = "invalid_json"
	errCodeMissingAttrib  = "missing_attrib"
	errCodeInvalidAttrib  = "invalid_attrib"
	errCodeBroadcastLimit = "broadcast_limit"
	errCode400            = "400"
	errCode422            = "422"
)

// Modica Mobile Gateway Errors
var (
	// ErrMobileGatewaySendFailed is returned when the API could not queue the
	// message due to an unknown error (API error code: send_failed).
	ErrMobileGatewaySendFailed = errors.New("could not queue message due to an unknown error")

	// ErrMobileGatewayInvalidJSON is returned when the request body contains
	// invalid JSON (API error code: invalid_json).
	ErrMobileGatewayInvalidJSON = errors.New("invalid json data in the request body")

	// ErrMobileGatewayMissingAttribute is returned when a required message
	// attribute is absent from the request (API error code: missing_attrib).
	ErrMobileGatewayMissingAttribute = errors.New("missing a required attribute")

	// ErrMobileGatewayInvalidAttribute is returned when a message attribute
	// contains an invalid value (API error code: invalid_attrib).
	ErrMobileGatewayInvalidAttribute = errors.New("invalid attribute value")

	// ErrMobileGatewayBroadcastLimit is returned when the account's broadcast
	// quota has been exceeded (API error code: broadcast_limit).
	ErrMobileGatewayBroadcastLimit = errors.New("broadcast limit has been exceeded")

	// ErrMobileGatewayInvalidTimestampFormat is returned when the scheduled
	// timestamp is not in RFC3339 format (HTTP 400).
	ErrMobileGatewayInvalidTimestampFormat = errors.New("invalid scheduled timestamp (must be rfc3339)")

	// ErrMobileGatewayInvalidTimestamp is returned when the scheduled timestamp
	// is in the past (HTTP 422).
	ErrMobileGatewayInvalidTimestamp = errors.New("invalid scheduled timestamp (must not be in the past)")

	// ErrMobileGatewayMessageIDNotFound is returned when a message id is not
	// returned from the API, but the request to create a new message was
	// successful.
	ErrMobileGatewayMessageIDNotFound = errors.New("message id not found")
)

var mobileGatewayErrorMap = map[string]error{
	errCodeSendFailed:     ErrMobileGatewaySendFailed,
	errCodeInvalidJson:    ErrMobileGatewayInvalidJSON,
	errCodeMissingAttrib:  ErrMobileGatewayMissingAttribute,
	errCodeInvalidAttrib:  ErrMobileGatewayInvalidAttribute,
	errCodeBroadcastLimit: ErrMobileGatewayBroadcastLimit,
	errCode400:            ErrMobileGatewayInvalidTimestampFormat,
	errCode422:            ErrMobileGatewayInvalidTimestamp,
}
