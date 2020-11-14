package main

import (
	"fmt"
	"log"
	"os"
)

// Level These are the integer logging levels used by the logger
type Level int

// Comment
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

//全局变量
var (
	logPrefix  = ""
	levelFlags = []string{"DEBG", "INFO", "WARN", "ERRO", "FATL"}

	logger  *log.Logger
	loggerf *log.Logger

	// curLevel ...
	curLevel Level
	logfile  *os.File
)

func init() {
	curLevel = DEBUG
	logger = log.New(os.Stdout, "[default] ", log.LstdFlags)
	logger.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

// printline ..
func printline(l *log.Logger, v ...interface{}) {
	if l != nil {
		l.Output(3, fmt.Sprintln(v...))
	}
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func fatalline(l *log.Logger, v ...interface{}) {
	if l != nil {
		l.Output(3, fmt.Sprintln(v...))
		os.Exit(1)
	}
}

// Debug ...
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	if DEBUG >= curLevel {
		printline(logger, v...)
		printline(loggerf, v...)
	}

}

// Info ...
func Info(v ...interface{}) {
	setPrefix(INFO)
	if INFO >= curLevel {
		printline(logger, v...)
		printline(loggerf, v...)
	}
}

// Warn ...
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	if WARNING >= curLevel {
		printline(logger, v...)
		printline(loggerf, v...)
	}
}

// Error Warn
func Error(v ...interface{}) {
	setPrefix(ERROR)
	if ERROR >= curLevel {
		printline(logger, v...)
		printline(loggerf, v...)
	}
}

// Fatal ...
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	if FATAL >= curLevel {
		fatalline(logger, v...)
		fatalline(loggerf, v...)
	}

}
func setPrefix(level Level) {
	logPrefix = fmt.Sprintf("[%s] ", levelFlags[level])
	logger.SetPrefix(logPrefix)
	if loggerf != nil {
		loggerf.SetPrefix(logPrefix)
	}
}

// Config ..
func config(level Level, lfile *os.File) {
	curLevel = level
	loggerf = log.New(lfile, "[default] ", log.LstdFlags)
	loggerf.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func testLogger() {
	lgfile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		Error("Failed to open log file:" + err.Error())
	}

	config(DEBUG, lgfile)

	Debug("message")
	Info("message")
	Warn("message")
	Error("message")
}
