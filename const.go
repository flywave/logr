package logr

import "time"

// Defaults.
const (
	// DefaultMaxQueueSize is the default maximum queue size for Logr instances.
	DefaultMaxQueueSize = 1000

	// DefaultMaxStackFrames is the default maximum max number of stack frames collected
	// when generating stack traces for logging.
	DefaultMaxStackFrames = 30

	// DefaultEnqueueTimeout is the default amount of time a log record can take to be queued.
	// This only applies to blocking enqueue which happen after `logr.OnQueueFull` is called
	// and returns false.
	DefaultEnqueueTimeout = time.Second * 30

	// DefaultShutdownTimeout is the default amount of time `logr.Shutdown` can execute before
	// timing out.
	DefaultShutdownTimeout = time.Second * 30

	// DefaultFlushTimeout is the default amount of time `logr.Flush` can execute before
	// timing out.
	DefaultFlushTimeout = time.Second * 30
)
