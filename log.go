package slog

// Fatal logs an entry as fatal error
func Fatal(entry Entry) {
	if level >= LEVEL_FATAL {
		entry["log_level"] = "fatal"
		out(entry)
	}
}

// Error logs an entry as error
func Error(entry Entry) {
	if level >= LEVEL_ERROR {
		entry["log_level"] = "error"
		out(entry)
	}
}

// Warning logs an entry as warning
func Warning(entry Entry) {
	if level >= LEVEL_WARNING {
		entry["log_level"] = "warning"
		out(entry)
	}
}

// Info logs an entry as information
func Info(entry Entry) {
	if level >= LEVEL_INFO {
		entry["log_level"] = "info"
		out(entry)
	}
}

// Debug logs an entry as debug output
func Debug(entry Entry) {
	if level >= LEVEL_DEBUG {
		entry["log_level"] = "debug"
		out(entry)
	}
}

// Trace logs an entry as an detailed debug output
// e.g. to trace a behavior at low level
func Trace(entry Entry) {
	if level >= LEVEL_TRACE {
		entry["log_level"] = "trace"
		out(entry)
	}
}
