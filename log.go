package ilog

// Log log interface
type Log interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
}

// SetLogger set log
func SetLogger(logimp Log) {
	if logimp == nil {
		logger = &stdLog{}
		return
	}
	logger = logimp
}

// Info log info
func Info(format string, args ...interface{}) {
	logger.Info(format, args...)
}

// Error log error
func Error(format string, args ...interface{}) {
	logger.Error(format, args...)
}

func init() {
	SetLogger(nil)
}
