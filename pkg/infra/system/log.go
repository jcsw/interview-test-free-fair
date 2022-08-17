package system

import (
	log "log"
	os "os"
)

var logger = initialize()

func initialize() *log.Logger {
	return log.New(os.Stdout, "[app] ", log.LstdFlags)
}

// Debug - Logging in level DEBUG
func Debug(log string, v ...interface{}) {
	logger.Printf("DEBUG "+log, v...)
}

// Info - Logging in level INFO
func Info(log string, v ...interface{}) {
	logger.Printf("INFO "+log, v...)
}

// Warn - Logging in level WARN
func Warn(log string, v ...interface{}) {
	logger.Printf("WARN "+log, v...)
}

// Error - Logging in level ERROR
func Error(log string, v ...interface{}) {
	logger.Printf("ERROR "+log, v...)
}

// Fatal - Logging in level FATAL
func Fatal(log string, v ...interface{}) {
	logger.Fatalf("FATAL "+log, v...)
}
