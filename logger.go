package epilog

import (
	"fmt"
	"time"

	"github.com/go-st/logger"
)

// DefaultHandler is a default StreamHandler with popular message formatter
var DefaultHandler = NewStreamHandler(logger.LevelDebug, NewTextFormatter("[:time:] (:level:) :message:"))

// Logger is a default implementation of ILogger interface
type Logger struct {
	name       string
	handler    IHandler
	processors []IProcessor
}

// New Returns new Logger instance
func New(name string, handler IHandler) *Logger {
	return &Logger{
		name:       name,
		handler:    handler,
		processors: make([]IProcessor, 0),
	}
}

// Copy creates copy of current logger
func (l *Logger) Copy() *Logger {
	return &Logger{
		name:       l.name,
		handler:    l.handler.Copy(),
		processors: l.processors,
	}
}

// AddProcessor adds entry processor to logger
func (l *Logger) AddProcessor(processors ...IProcessor) {
	l.processors = append(l.processors, processors...)
}

// Log logs new entry with specified level
func (l *Logger) Log(level logger.Level, args ...interface{}) {
	l.log(level, args...)
}

// Debug alias for log with debug level
func (l *Logger) Debug(args ...interface{}) {
	l.log(logger.LevelDebug, args...)
}

// Info alias for log with Info level
func (l *Logger) Info(args ...interface{}) {
	l.log(logger.LevelInfo, args...)
}

// Notice alias for log with notice level
func (l *Logger) Notice(args ...interface{}) {
	l.log(logger.LevelNotice, args...)
}

// Warning alias for log with warning level
func (l *Logger) Warning(args ...interface{}) {
	l.log(logger.LevelWarning, args...)
}

// Error alias for log with error level
func (l *Logger) Error(args ...interface{}) {
	l.log(logger.LevelError, args...)
}

// Critical alias for log with critical level
func (l *Logger) Critical(args ...interface{}) {
	l.log(logger.LevelCritical, args...)
}

// Alert alias for log with alert level
func (l *Logger) Alert(args ...interface{}) {
	l.log(logger.LevelAlert, args...)
}

// Emergency alias for log with emergency level
func (l *Logger) Emergency(args ...interface{}) {
	l.log(logger.LevelEmergency, args...)
}

// Logf logs new entry with specified level
func (l *Logger) Logf(level logger.Level, format string, args ...interface{}) {
	l.logf(level, format, args...)
}

// Debugf alias for log with debug level
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logf(logger.LevelDebug, format, args...)
}

// Infof alias for log with info level
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logf(logger.LevelInfo, format, args...)
}

// Noticef alias for log with notice level
func (l *Logger) Noticef(format string, args ...interface{}) {
	l.logf(logger.LevelNotice, format, args...)
}

// Warningf alias for log with warning level
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.logf(logger.LevelWarning, format, args...)
}

// Errorf alias for log with error level
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logf(logger.LevelError, format, args...)
}

// Criticalf alias for log with critical level
func (l *Logger) Criticalf(format string, args ...interface{}) {
	l.logf(logger.LevelCritical, format, args...)
}

// Alertf alias for log with alert level
func (l *Logger) Alertf(format string, args ...interface{}) {
	l.logf(logger.LevelAlert, format, args...)
}

// Emergencyf alias for log with emergency level
func (l *Logger) Emergencyf(format string, args ...interface{}) {
	l.logf(logger.LevelEmergency, format, args...)
}

// IsDebugEnabled returns true if logger is enabled for logger.LevelDebug and false otherwise
func (l *Logger) IsDebugEnabled() bool {
	return l.handler.IsEnabledFor(logger.LevelDebug)
}

// IsInfoEnabled returns true if logger is enabled for logger.LevelInfo and false otherwise
func (l *Logger) IsInfoEnabled() bool {
	return l.handler.IsEnabledFor(logger.LevelInfo)
}

// IsNoticeEnabled returns true if logger is enabled for logger.LevelNotice and false otherwise
func (l *Logger) IsNoticeEnabled() bool {
	return l.handler.IsEnabledFor(logger.LevelNotice)
}

// IsWarningEnabled returns true if logger is enabled for logger.LevelWarning and false otherwise
func (l *Logger) IsWarningEnabled() bool {
	return l.handler.IsEnabledFor(logger.LevelWarning)
}

// IsErrorEnabled returns true if logger is enabled for logger.LevelWarning and false otherwise
func (l *Logger) IsErrorEnabled() bool {
	return l.handler.IsEnabledFor(logger.LevelError)
}

// IsAlertEnabled returns true if logger is enabled for logger.LevelAlert and false otherwise
func (l *Logger) IsAlertEnabled() bool {
	return l.handler.IsEnabledFor(logger.LevelAlert)
}

// IsCriticalEnabled returns true if logger is enabled for logger.LevelCritical and false otherwise
func (l *Logger) IsCriticalEnabled() bool {
	return l.handler.IsEnabledFor(logger.LevelCritical)
}

// IsEmergencyEnabled returns true if logger is enabled for logger.LevelEmergency and false otherwise
func (l *Logger) IsEmergencyEnabled() bool {
	return l.handler.IsEnabledFor(logger.LevelEmergency)
}

// log message with level check
func (l *Logger) log(level logger.Level, args ...interface{}) {
	if l.handler.IsEnabledFor(level) {
		l.handle(level, args...)
	}
}

// logf message with level check (important to do heavy Sprintf after level check)
func (l *Logger) logf(level logger.Level, format string, args ...interface{}) {
	if l.handler.IsEnabledFor(level) {
		l.handle(level, fmt.Sprintf(format, args...))
	}
}

// handle puts entry to handle
func (l *Logger) handle(level logger.Level, args ...interface{}) {
	entry := NewEntry(level, time.Now(), fmt.Sprint(args...))
	for _, processor := range l.processors {
		processor.Process(entry)
	}
	entry.Fields["_module"] = l.name
	l.handler.Handle(entry)
}
