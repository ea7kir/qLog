/*
 *  Custom Logger
 *  Copyright (c) 2023 Michael Naylor EA7KIR (https://michaelnaylor.es)
 */

// Package qLog implements a simple logging package, based entirely on
// the Standard log Library Pakage, but adds preformatted log levels.
// DEGUG, INFO, WARN, ERROR and FATAL.  The required log level is set
// by function calls SetDebg, SetInfo, SetWarn, SetError and SetFatal.
// In addtion, SetNone can be called to disable all log messages.
//
// Ouput messages are formatted as follows:
//
// ERROR 2023/09/03 18:02:24 file.go:82: something went wrong
package qLog

import (
	"fmt"
	"log"
	"os"
)

const (
	kDEBUG = -1
	kINFO  = 0
	kWARN  = 1
	kERROR = 2
	kFATAL = 3
	kNONE  = 4
)

var (
	logLevel int
	logFile  *os.File
	pDebug   *log.Logger
	pInfo    *log.Logger
	pWarn    *log.Logger
	pError   *log.Logger
	pFatal   *log.Logger
)

// SetOutput must be called before use.  Output can be any standard
// output or file path.
func SetOutput(file *os.File) {
	logFile = file
	flags := log.Ldate | log.Ltime | log.Lshortfile
	pDebug = log.New(logFile, "DEBUG ", flags)
	pInfo = log.New(logFile, "INFO ", flags)
	pWarn = log.New(logFile, "WARN ", flags)
	pError = log.New(logFile, "ERROR ", flags)
	pFatal = log.New(logFile, "FATAL ", flags)
}

// func Open(output string) {
// 	f, err := os.OpenFile(output, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	SetOutput(f)
// 	// flags := log.Ldate | log.Ltime | log.Lshortfile
// 	// pDebug = log.New(logFile, "DEBUG ", flags)
// 	// pInfo = log.New(logFile, "INFO ", flags)
// 	// pWarn = log.New(logFile, "WARN ", flags)
// 	// pError = log.New(logFile, "ERROR ", flags)
// 	// pFatal = log.New(logFile, "FATAL ", flags)
// }

// Close closes the logging ouput file.
func Close() {
	logFile.Close()
}

// SetDebugLevel sets messages to DEBUG and above.
func SetDebugLevel() {
	logLevel = kDEBUG
}

// SetInfoLevel retricts messages to INFO and above.
func SetInfoLevel() {
	logLevel = kINFO
}

// SetWarnLevel retricts messages to WARN and above.
func SetWarnLevel() {
	logLevel = kWARN
}

// SetWarnLevel retricts messages to ERROR and above.
func SetErrorLevel() {
	logLevel = kERROR
}

// SetFatalLevel retricts messages to FATAL.
func SetFatalLevel() {
	logLevel = kFATAL
}

// SetNoneLevel disables all log messages.
func SetNoneLevel() {
	logLevel = kNONE
}

// Debug prints the arguments, prefixed with DEBUG.
func Debug(msg string, args ...interface{}) {
	if logLevel != kDEBUG {
		return
	}
	pDebug.Output(2, fmt.Sprintf(msg, args...))
}

// Indo prints the arguments, prefixed with INFO.
func Info(msg string, args ...interface{}) {
	if logLevel <= kINFO {
		pInfo.Output(2, fmt.Sprintf(msg, args...))
	}
}

// Warn prints the arguments, prefixed with WARN.
func Warn(msg string, args ...interface{}) {
	if logLevel <= kWARN {
		pWarn.Output(2, fmt.Sprintf(msg, args...))
	}
}

// Error prints the arguments, prefixed with ERROR.
func Error(msg string, args ...interface{}) {
	if logLevel <= kERROR {
		pError.Output(2, fmt.Sprintf(msg, args...))
	}
}

// Fatal prints the arguments, prifexed with FATAL and followed by a call to os.Exit(1).
func Fatal(msg string, args ...interface{}) {
	if logLevel <= kFATAL {
		pFatal.Output(2, fmt.Sprintf(msg, args...))
		os.Exit(1)
	}
}
