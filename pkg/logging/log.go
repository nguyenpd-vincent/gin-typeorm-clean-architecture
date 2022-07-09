package logging

import (
	"fmt"
	"github.com/pdnguyen1503/base-go/pkg/file"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	ResetColor  = "\033[0m"
	logger      *log.Logger
	logPrefix   = ""
	levelFlags  = []string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}
	levelColors = []string{"\033[34m", "\033[32m", "\033[33m", "\033[31m", ""}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// setup initialize the log instance
func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatal("logging.Setup Error: %v", err)
	}
	mw := io.MultiWriter(os.Stdout, F)
	logger = log.New(mw, DefaultPrefix, log.LstdFlags)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
	//utils.Info(fmt.Sprintf("Log: %v", v))
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
	//utils.Info(fmt.Sprintf("Log: %v", v))
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
	//utils.Warn(fmt.Sprintf("Log: %v", v))
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
	//utils.Error(fmt.Sprintf("Log: %v", v))
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
	//utils.Fatal(fmt.Sprintf("Log: %v", v))
}

// getColor
func getColor(level Level) string {
	if runtime.GOOS == "windows" {
		return ""
	}
	return levelColors[level]
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	colo := getColor(level)
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("%s[%s]%s[%s:%d]", colo, levelFlags[level], ResetColor, filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("%s[%s]%s", colo, levelFlags[level], ResetColor)
	}

	logger.SetPrefix(logPrefix)
}
