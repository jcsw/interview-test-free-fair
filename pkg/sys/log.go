package sys

import (
	log "log"
	os "os"
)

var logger = initialize()

func initialize() *log.Logger {
	return log.New(os.Stdout, "[app] ", log.LstdFlags)
}

// LogDebug - Logging in level DEBUG
func LogDebug(log string, v ...interface{}) {
	logger.Printf("DEBUG "+log, v...)
}

// LogInfo - Logging in level INFO
func LogInfo(log string, v ...interface{}) {
	logger.Printf("INFO "+log, v...)
}

// LogWarn - Logging in level WARN
func LogWarn(log string, v ...interface{}) {
	logger.Printf("WARN "+log, v...)
}

// LogError - Logging in level ERROR
func LogError(log string, v ...interface{}) {
	logger.Printf("ERROR "+log, v...)
}

// LogFatal - Logging in level FATAL
func LogFatal(log string, v ...interface{}) {
	logger.Fatalf("FATAL "+log, v...)
}
