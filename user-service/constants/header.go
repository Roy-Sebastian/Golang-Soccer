package constants

import "net/textproto"

var (
	XApiKey       = textproto.CanonicalMIMEHeaderKey("x-api-key")
	XServiceName  = textproto.CanonicalMIMEHeaderKey("x-service-name")
	XRequestAt    = textproto.CanonicalMIMEHeaderKey("x-request-at")
	Authorization = textproto.CanonicalMIMEHeaderKey("authorization")
)
