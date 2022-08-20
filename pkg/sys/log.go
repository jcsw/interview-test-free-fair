package sys

import (
	"fmt"
	log "log"
	os "os"
	"time"
)

var logger = initialize()

func initialize() *log.Logger {

	logDate := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("./logs/%v.log", logDate)
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return log.New(file, "[app] ", log.LstdFlags)
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
