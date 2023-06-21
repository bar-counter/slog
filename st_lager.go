package slog

import (
	"fmt"
	"github.com/bar-counter/slog/lager"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	//DEBUG is a constant of string type
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

// Config
//
//	is a struct which stores details for maintaining logs
type Config struct {
	LoggerLevel    string
	LoggerFile     string
	Writers        []string
	EnableRsyslog  bool
	RsyslogNetwork string
	RsyslogAddr    string

	LogFormatText bool
	LogHideLineno bool
}

var config = defaultConfig()
var m sync.RWMutex

// writers is a map
var writers = make(map[string]io.Writer)

// RegisterWriter is used to register a io writer
func RegisterWriter(name string, writer io.Writer) {
	m.Lock()
	writers[name] = writer
	m.Unlock()
}

// defaultConfig
//
//	is a function which
//	return config object with default configuration
func defaultConfig() *Config {
	return &Config{
		LoggerLevel:    INFO,
		LoggerFile:     "",
		EnableRsyslog:  false,
		RsyslogNetwork: "udp",
		RsyslogAddr:    "127.0.0.1:5140",
		LogFormatText:  false,
	}
}

// lagerInit
//
//	is a function which initializes all config struct variables
func lagerInit(c Config) {
	if c.LoggerLevel != "" {
		config.LoggerLevel = c.LoggerLevel
	}

	if c.LoggerFile != "" {
		config.LoggerFile = c.LoggerFile
		config.Writers = append(config.Writers, "file")
	}

	if c.EnableRsyslog {
		config.EnableRsyslog = c.EnableRsyslog
	}

	if c.RsyslogNetwork != "" {
		config.RsyslogNetwork = c.RsyslogNetwork
	}

	if c.RsyslogAddr != "" {
		config.RsyslogAddr = c.RsyslogAddr
	}
	if len(c.Writers) == 0 {
		config.Writers = append(config.Writers, "stdout")

	} else {
		config.Writers = c.Writers
	}
	config.LogFormatText = c.LogFormatText
	RegisterWriter("stdout", os.Stdout)
	var file io.Writer
	var err error
	if config.LoggerFile != "" {
		file, err = os.OpenFile(config.LoggerFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

	}
	for _, sink := range config.Writers {
		if sink == "file" {
			if file == nil {
				log.Panic("Must set file path")
			}
			RegisterWriter("file", file)
		}
	}
}

// newLogger
//
//	new st lager
func newLogger(component string, isLogFormatText, isHideLineno bool) lager.Logger {
	return NewLoggerExt(component, component, isLogFormatText, isHideLineno)
}

// NewLoggerExt is a function which is used to write new logs
func NewLoggerExt(component string, appGUID string, isLogFormatText, isHideLineno bool) lager.Logger {
	var lagerLogLevel lager.LogLevel
	switch strings.ToUpper(config.LoggerLevel) {
	case DEBUG:
		lagerLogLevel = lager.DEBUG
	case INFO:
		lagerLogLevel = lager.INFO
	case WARN:
		lagerLogLevel = lager.WARN
	case ERROR:
		lagerLogLevel = lager.ERROR
	case FATAL:
		lagerLogLevel = lager.FATAL
	default:
		panic(fmt.Errorf("unknown logger level: %s", config.LoggerLevel))
	}
	logger := lager.NewLoggerExt(component, isLogFormatText, isHideLineno)
	for _, sink := range config.Writers {

		writer, ok := writers[sink]
		if !ok {
			log.Panic("unknown writer: ", sink)
		}
		sink := lager.NewReconfigurableSink(lager.NewWriterSink(sink, writer, lager.DEBUG), lagerLogLevel)
		logger.RegisterSink(sink)
	}

	return logger
}

func Debug(action string, data ...lager.Data) {
	logger.Debug(action, data...)
}

func Info(action string, data ...lager.Data) {
	logger.Info(action, data...)
}

func Warn(action string, data ...lager.Data) {
	logger.Warn(action, data...)
}

func Error(action string, err error, data ...lager.Data) {
	logger.Error(action, err, data...)
}

func Fatal(action string, err error, data ...lager.Data) {
	logger.Fatal(action, err, data...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Errorf(err error, format string, args ...interface{}) {
	logger.Errorf(err, format, args...)
}

func Fatalf(err error, format string, args ...interface{}) {
	logger.Fatalf(err, format, args...)
}
