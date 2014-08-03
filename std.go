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

// Error if error, similar to panic but just log the error and returns
// true if we got an error. Useful in if statements.
func Errorif(err error) bool {
	if err == nil {
		return false
	}
	std.Write(3, ERROR, err.Error())
	return true
}

// Panic if error
func Panicif(err error) {
	if err == nil {
		return
	}
	std.Write(3, ERROR, err.Error())
	panic(err)
}

// Fatal log.
func Fatalif(err error) {
	if err == nil {
		return
	}
	std.Write(3, FATAL, err.Error())
	os.Exit(1)
}
