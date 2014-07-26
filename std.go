//
// Default logger:
//
//   log.Info("something %s", "here")
//
package log

import "os"

var std = New(os.Stderr, INFO, "")

func SetPrefix(str string) {
	std.SetPrefix(str)
}

func SetLevel(level Level) {
	std.SetLevel(level)
}

// Debug log.
func Debug(msg string, args ...interface{}) error {
	return std.Write(3, DEBUG, msg, args...)
}

// Info log.
func Info(msg string, args ...interface{}) error {
	return std.Write(3, INFO, msg, args...)
}

// Warning log.
func Warning(msg string, args ...interface{}) error {
	return std.Write(3, WARNING, msg, args...)
}

// Error log.
func Error(msg string, args ...interface{}) error {
	return std.Write(3, ERROR, msg, args...)
}

// Fatal log.
func Fatal(msg string, args ...interface{}) error {
	return std.Write(3, FATAL, msg, args...)
}
