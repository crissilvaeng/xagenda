package support

import (
	"log"
	"os"
)

// LogLevel represents the logging level.
type LogLevel uint8

const (
	// LevelInfo is the level for info messages.
	LevelInfo LogLevel = 1 << iota
	// LevelError is the error level. It is used for error messages. (10)
	LevelError
)

// Logger provides logging functionality to info and error messages.
// Logger is thread-safe, can be used from multiple goroutines. But keep in mind that the usage of mutaxes increases the latency.
// About only support two log levels, please refer: https://dave.cheney.net/2015/11/05/lets-talk-about-logging
type Logger struct {
	level    LogLevel
	LogInfo  *log.Logger
	LogError *log.Logger
}

// Info logs an info message in stdout if logger level is set to LevelInfo.
func (l *Logger) Info(v ...interface{}) {
	if l.level&LevelInfo != 0 {
		l.LogInfo.Println(v...)
	}
}

// Infof logs an info message in stdout with format if logger level is set to LevelInfo.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level&LevelInfo != 0 {
		l.LogInfo.Printf(format, v...)
	}
}

// Error logs an error message in stderr if logger level is set to LevelError.
func (l *Logger) Error(v ...interface{}) {
	if l.level&LevelError != 0 {
		l.LogError.Println(v...)
	}
}

// Error logs an error message with format in stderr if logger level is set to LevelError.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level&LevelError != 0 {
		l.LogError.Printf(format, v...)
	}
}

// NewLogger creates a new logger. Accepts log level as parameter.
// Log info messages are printed to stdout and error messages to stderr.
func NewLogger(level LogLevel) *Logger {
	flags := log.Ldate | log.Ltime | log.Lshortfile | log.LUTC | log.Lmsgprefix
	return &Logger{
		level:    level,
		LogInfo:  log.New(os.Stdout, "[INFO] ", flags),
		LogError: log.New(os.Stderr, "[ERROR] ", flags),
	}
}
