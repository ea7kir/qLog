/*
 *  Custom Logger
 *  Copyright (c) 2023 Michael Naylor EA7KIR (https://michaelnaylor.es)
 */

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

func Close() {
	logFile.Close()
}

func SetDebugLevel() {
	logLevel = kDEBUG
}

func SetInfoLevel() {
	logLevel = kINFO
}

func SetWarnLevel() {
	logLevel = kWARN
}

func SetErrorLevel() {
	logLevel = kERROR
}

func SetFatalLevel() {
	logLevel = kFATAL
}

func SetNoneLevel() {
	logLevel = kNONE
}

func Debug(msg string, args ...interface{}) {
	if logLevel != kDEBUG {
		return
	}
	pDebug.Output(2, fmt.Sprintf(msg, args...))
}

func Info(msg string, args ...interface{}) {
	if logLevel <= kINFO {
		pInfo.Output(2, fmt.Sprintf(msg, args...))
	}
}

func Warn(msg string, args ...interface{}) {
	if logLevel <= kWARN {
		pWarn.Output(2, fmt.Sprintf(msg, args...))
	}
}

func Error(msg string, args ...interface{}) {
	if logLevel <= kERROR {
		pError.Output(2, fmt.Sprintf(msg, args...))
	}
}

func Fatal(msg string, args ...interface{}) {
	if logLevel <= kFATAL {
		pFatal.Output(2, fmt.Sprintf(msg, args...))
	}
}
