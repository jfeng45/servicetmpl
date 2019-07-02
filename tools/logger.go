// Package tools represents a wrapper to third party library. It may enhance or modify some of the functionalities
// of a third party library or create a interface for it.

package tools

// Log is a package level variable, every program should access logging function through "Log"
var Log Logger

// Logger represent common interface for logging function
type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info( args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

// SetLogger is the setter for log variable, it should be the only way to assign value to log
func SetLogger(newLogger Logger) {
	Log=newLogger
}


