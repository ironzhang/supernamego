package public

import "errors"

// logger name
const LoggerName = "supername"

// errors
var (
	ErrClusterNotFound = errors.New("can not find cluster")
)
