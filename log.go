package vlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

// Logging level.
const (
	Off = iota
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
)

//goland:noinspection ALL
const (
	info  = "info"
	debug = "debug"
	off   = "off"
	trace = "trace"
	warn  = "warn"
	error = "error"
	fatal = "fatal"
)

type LevelEnum struct {
	INFO  LevelType
	DEBUG LevelType
	OFF   LevelType
	TRACE LevelType
	WARN  LevelType
	ERROR LevelType
	FATAL LevelType
}

type LevelType string

var Level = LevelEnum{
	INFO:  info,
	DEBUG: debug,
	OFF:   off,
	TRACE: trace,
	WARN:  warn,
	ERROR: error,
	FATAL: fatal,
}

var syncOutput = false

var mu sync.Mutex

var writer *io.Writer

// New creates a logger.
func New(out io.Writer) *Logger {
	ret := &Logger{level: logLevel, logger: log.New(out, "", log.Ldate|log.Ltime|log.Lshortfile)}

	loggers = append(loggers, ret)

	return ret
}

// Default create default logger
func Default() *Logger {
	if writer != nil {
		return New(*writer)
	}
	return New(os.Stdout)
}

// SetLevel sets the logging level of all loggers. (global)
func SetLevel(level LevelType) {
	mu.Lock()

	defer mu.Unlock()
	logLevel = getLevel(level)

	for _, l := range loggers {
		l.SetLevel(level)
	}
}

// SetOutput sets the output destination for the standard logger. (global)
func SetOutput(w io.Writer) {
	mu.Lock()
	defer mu.Unlock()
	writer = &w
	if syncOutput {
		log.SetOutput(w)
	}
	for _, logger := range loggers {
		logger.SetOutput(w)
	}
}

// SetSyncOutput synchronize output settings to the default log library
func SetSyncOutput(b bool) {
	mu.Lock()
	defer mu.Unlock()
	syncOutput = b
}

// all loggers.
var loggers []*Logger

// the global default logging level, it will be used for creating logger.
var logLevel = Info

// Logger represents a simple logger with level.
// The underlying logger is the standard Go logging "log".
type Logger struct {
	level  int
	logger *log.Logger
}

// getLevel gets logging level int value corresponding to the specified level.
func getLevel(level LevelType) int {

	switch level {
	case off:
		return Off
	case trace:
		return Trace
	case debug:
		return Debug
	case info:
		return Info
	case warn:
		return Warn
	case error:
		return Error
	case fatal:
		return Fatal
	default:
		return Info
	}
}

// SetOutput sets the output destination for the standard logger.
func (l *Logger) SetOutput(w io.Writer) {
	l.logger.SetOutput(w)
}

// SetLevel sets the logging level of a logger.
func (l *Logger) SetLevel(level LevelType) {
	l.level = getLevel(level)
}

// IsTraceEnabled determines whether the trace level is enabled.
func (l *Logger) IsTraceEnabled() bool {
	return l.level <= Trace
}

// IsDebugEnabled determines whether the debug level is enabled.
func (l *Logger) IsDebugEnabled() bool {
	return l.level <= Debug
}

// IsWarnEnabled determines whether the debug level is enabled.
func (l *Logger) IsWarnEnabled() bool {
	return l.level <= Warn
}

// Trace prints trace level message.
func (l *Logger) Trace(v ...interface{}) {
	if Trace < l.level {
		return
	}

	l.logger.SetPrefix("TRACE ")
	err := l.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Tracef prints trace level message with format.
func (l *Logger) Tracef(format string, v ...interface{}) {
	if Trace < l.level {
		return
	}
	l.logger.SetPrefix("TRACE ")
	err := l.logger.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Debug prints debug level message.
func (l *Logger) Debug(v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.logger.SetPrefix("DEBUG ")
	err := l.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Debugf prints debug level message with format.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.logger.SetPrefix("DEBUG ")
	err := l.logger.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Info prints info level message.
func (l *Logger) Info(v ...interface{}) {
	if Info < l.level {
		return
	}

	l.logger.SetPrefix("INFO ")
	err := l.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Infof prints info level message with format.
func (l *Logger) Infof(format string, v ...interface{}) {
	if Info < l.level {
		return
	}

	l.logger.SetPrefix("INFO ")
	err := l.logger.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Warn prints warning level message.
func (l *Logger) Warn(v ...interface{}) {
	if Warn < l.level {
		return
	}

	l.logger.SetPrefix("WARN ")
	err := l.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Warnf prints warning level message with format.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if Warn < l.level {
		return
	}

	l.logger.SetPrefix("WARN ")
	err := l.logger.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Error prints error level message.
func (l *Logger) Error(v ...interface{}) {
	if Error < l.level {
		return
	}

	l.logger.SetPrefix("ERROR ")
	err := l.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Errorf prints error level message with format.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if Error < l.level {
		return
	}

	l.logger.SetPrefix("ERROR ")
	err := l.logger.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Fatal prints fatal level message and exit process with code 1.
func (l *Logger) Fatal(v ...interface{}) {
	if Fatal < l.level {
		return
	}

	l.logger.SetPrefix("FATAL ")
	err := l.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Exit(1)
}

// Fatalf prints fatal level message with format and exit process with code 1.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if Fatal < l.level {
		return
	}

	l.logger.SetPrefix("FATAL ")
	err := l.logger.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Exit(1)
}
