package slog

import (
	"strings"
	"time"
)

// LogFormat defines the type to define the output format
// see enums below
type LogFormat int

const (
	FORMAT_JSON LogFormat = iota
	FORMAT_TEXT
)

// LogLevel defines the type for log level definition
// see enum below
type LogLevel int

const (
	LEVEL_NONE LogLevel = iota
	LEVEL_FATAL
	LEVEL_ERROR
	LEVEL_WARNING
	LEVEL_INFO
	LEVEL_DEBUG
	LEVEL_TRACE
)

// LogFiles define where to write the log files to
type LogFiles struct {
	Path    string
	Postfix string
}

// GetPath return current log file to use
// empty if not used
func (lf LogFiles) GetPath() string {
	if lf.Path == "" {
		return ""
	}

	timestamp := time.Now().Format("2006-01-02")
	if strings.HasSuffix(lf.Path, "/") || strings.HasSuffix(lf.Path, "\\") {
		return lf.Path + timestamp + lf.Postfix + ".log"
	}
	return lf.Path + "/" + timestamp + lf.Postfix + ".log"
}

// Some variables to store the configuration
var (
	level      = LEVEL_INFO
	format     = FORMAT_JSON
	timeFormat = "2006-01-02 15:04:05"
	logFiles   = LogFiles{Path: ""}
)

// Initialize is a function to set the three importent configuration at once
func Initialize(newLevel LogLevel, newFormat LogFormat, newLogFiles LogFiles) {
	level = newLevel
	format = newFormat
	logFiles = newLogFiles
}

// SetLogFiles can be used to set/change the log file configuration
func SetLogFiles(newLogFiles LogFiles) {
	logFiles = newLogFiles
}

// SetLogLevel can be used to set/change the log level
func SetLogLevel(newLevel LogLevel) {
	level = newLevel
}

// SetLogFormat can be used to set/change the log format
func SetLogFormat(newFormat LogFormat) {
	format = newFormat
}

// SetTimeFormat can be used to set/change the time formation
func SetTimeFormat(newTimeFormat string) {
	timeFormat = newTimeFormat
}
