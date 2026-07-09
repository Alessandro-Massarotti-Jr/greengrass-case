package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type LogLevel string

const (
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
	LogLevelDebug LogLevel = "debug"
)

type LogInput struct {
	Action  string
	Message string
	Data    map[string]any
}

type logFormat struct {
	Timestamp string         `json:"timestamp"`
	Level     LogLevel       `json:"level"`
	Action    string         `json:"action"`
	Message   string         `json:"message"`
	Data      map[string]any `json:"data"`
}

type Logger struct{}

var (
	loggerInstance *Logger
	loggerOnce     sync.Once
)

func GetLogger() *Logger {
	loggerOnce.Do(func() {
		loggerInstance = &Logger{}
	})
	return loggerInstance
}

func (l *Logger) log(level LogLevel, input LogInput) {
	if input.Data == nil {
		input.Data = map[string]any{}
	}

	entry, err := json.Marshal(logFormat{
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
		Level:     level,
		Action:    input.Action,
		Message:   input.Message,
		Data:      input.Data,
	})
	if err != nil {
		return
	}

	fmt.Println(string(entry))
}

func (l *Logger) Info(input LogInput) {
	l.log(LogLevelInfo, input)
}

func (l *Logger) Warn(input LogInput) {
	l.log(LogLevelWarn, input)
}

func (l *Logger) Error(input LogInput) {
	l.log(LogLevelError, input)
}

func (l *Logger) Debug(input LogInput) {
	l.log(LogLevelDebug, input)
}
