//
// Default logger:
//
//   log.Info("something %s", "here")
//
package log

import "os"

var std = New(os.Stderr, INFO, "")

// SetPrefix wrapper.
func SetPrefix(str string) {
	std.SetPrefix(str)
}

// SetLevel wrapper.
func SetLevel(level Level) {
	std.SetLevel(level)
}

// SetLevelString wrapper.
func SetLevelString(level string) {
	std.SetLevelString(level)
}

// Debug log.
func Debug(msg string, args ...interface{}) error {
	return std.Debug(msg, args...)
}

// Info log.
func Info(msg string, args ...interface{}) error {
	return std.Info(msg, args...)
}

// Notice log.
func Notice(msg string, args ...interface{}) error {
	return std.Notice(msg, args...)
}

// Warning log.
func Warning(msg string, args ...interface{}) error {
	return std.Warning(msg, args...)
}

// Error log.
func Error(msg string, args ...interface{}) error {
	return std.Error(msg, args...)
}

// Critical log.
func Critical(msg string, args ...interface{}) error {
	return std.Critical(msg, args...)
}

// Alert log.
func Alert(msg string, args ...interface{}) error {
	return std.Alert(msg, args...)
}

// Emergency log.
func Emergency(msg string, args ...interface{}) error {
	return std.Emergency(msg, args...)
}

// Check if there's an `err` and exit, useful for bootstrapping.
func Check(err error) {
	std.Check(err)
}
