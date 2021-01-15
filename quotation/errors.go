package quotation

import "errors"

// ErrClientRedisFailedConnection connection is failed
var ErrClientRedisFailedConnection = errors.New("Client redis connection failed")

// ErrMissingTagOrAttibute tag or attribute not found
var ErrMissingTagOrAttibute = errors.New("Missing tag or attribute in the node tree")
