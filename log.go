//
// Simple logger similar to Go's standard logger with log level.
//
//   l := log.New(os.Stderr, INFO, "myapp")
//   l.Error("something exploded")
//
package log

import "time"
import "sync"
import "fmt"
import "io"
import "github.com/jehiah/go-strftime"

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var Levels = map[Level]string{
	DEBUG:   "DEBUG",
	INFO:    "INFO",
	WARNING: "WARNING",
	ERROR:   "ERROR",
	FATAL:   "FATAL",
}

func (l Level) String() string { return Levels[l] }

type Formatter func(ctx string, level Level, msg string) string

type Logger struct {
	Writer io.Writer
	Level  Level
	Prefix string
	Format Formatter
	sync.Mutex
}

// New logger which writes to `w` at the given `level`. Optionally
// provide a `prefix` for the logger.
func New(w io.Writer, level Level, prefix string) *Logger {
	l := &Logger{Writer: w, Level: level, Prefix: prefix, Format: StandardFormater}
	l.SetPrefix(prefix)
	return l
}

// SetPrefix changes the prefix to `str`.
func (l *Logger) SetPrefix(str string) {
	l.Lock()
	defer l.Unlock()

	l.Prefix = str
}

// SetLevel changes the log `level`.
func (l *Logger) SetLevel(level Level) {
	l.Lock()
	defer l.Unlock()

	l.Level = level
}

// Standard Formatter
func StandardFormater(prefix string, level Level, msg string) string {
	ts := strftime.Format("%Y-%m-%d %H:%M:%S", time.Now())
	return fmt.Sprintf("%s %s %s - %s", ts, prefix, level, msg)
}

// Write a message.
func (l *Logger) Write(level Level, msg string, args ...interface{}) error {
	l.Lock()
	defer l.Unlock()

	// return early
	if l.Level > level {
		return nil
	}

	// append a "\n" only when necessary
	if l := len(msg); l == 0 || msg[l-1] != '\n' {
		msg += "\n" // not super performant...
	}

	// format the output using a "custom" function
	f := l.Format(l.Prefix, level, msg)

	_, err := fmt.Fprintf(l.Writer, f, args...)
	return err
}

// Debug log.
func (l *Logger) Debug(msg string, args ...interface{}) error {
	return l.Write(DEBUG, msg, args...)
}

// Info log.
func (l *Logger) Info(msg string, args ...interface{}) error {
	return l.Write(INFO, msg, args...)
}

// Warning log.
func (l *Logger) Warning(msg string, args ...interface{}) error {
	return l.Write(WARNING, msg, args...)
}

// Error log.
func (l *Logger) Error(msg string, args ...interface{}) error {
	return l.Write(ERROR, msg, args...)
}

// Fatal log.
func (l *Logger) Fatal(msg string, args ...interface{}) error {
	return l.Write(FATAL, msg, args...)
}
