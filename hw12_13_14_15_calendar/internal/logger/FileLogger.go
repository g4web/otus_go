package logger

import (
	"io"
	"log"
	"os"
)

type FileLogger struct {
	level    int32
	filePath string
	logFile  *os.File
}

const (
	LogLevelDebug   = "debug"
	LogLevelInfo    = "info"
	LogLevelWarning = "warn"
	LogLevelError   = "error"
)

var logLevels = map[string]int32{
	LogLevelDebug:   10,
	LogLevelInfo:    20,
	LogLevelWarning: 30,
	LogLevelError:   40,
}

func New(minLevel string, filePath string) *FileLogger {
	logLevel := logLevels[minLevel]

	logFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return &FileLogger{logLevel, filePath, logFile}
}

func (l *FileLogger) Debug(msg string) {
	l.save(msg, LogLevelDebug)
}

func (l *FileLogger) Info(msg string) {
	l.save(msg, LogLevelInfo)
}

func (l *FileLogger) Warning(msg string) {
	l.save(msg, LogLevelWarning)
}

func (l *FileLogger) Error(msg string) {
	l.save(msg, LogLevelError)
}

func (l *FileLogger) save(msg string, level string) {
	logLevel := logLevels[level]

	if logLevel == 0 || logLevel < l.level {
		return
	}

	wrt := io.MultiWriter(os.Stdout, l.logFile)
	log.SetOutput(wrt)

	log.Println(level + ": " + msg)
}

func (l *FileLogger) Close() {
	err := l.logFile.Close()
	if err != nil {
		log.Fatalf("can't close log file: %v", err)
	}
}
