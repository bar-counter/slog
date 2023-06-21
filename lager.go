package slog

import (
	"encoding/json"
	"fmt"
	"github.com/bar-counter/slog/lager"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// constant values for logrotate parameters
const (
	RollingPolicySize  = "size"
	RollingPolicyDaily = "daily"
	LogRotateDate      = 1
	LogRotateSize      = 10
	LogBackupCount     = 7
	//DefaultLagerLogsFile = "logs/chassis.log"
)

// Lager struct for logger parameters
type Lager struct {
	Writers        string `yaml:"writers"`
	LoggerLevel    string `yaml:"logger_level"`
	LoggerFile     string `yaml:"logger_file"`
	LogHideLineno  bool   `yaml:"log_hide_lineno"`
	LogFormatText  bool   `yaml:"log_format_text"`
	RollingPolicy  string `yaml:"rolling_policy"`
	LogRotateDate  int    `yaml:"log_rotate_date"`
	LogRotateSize  int    `yaml:"log_rotate_size"`
	LogBackupCount int    `yaml:"log_backup_count"`
}

// PassLagerCfg is the struct for lager information(passlager.yaml)
type PassLagerCfg struct {
	Writers        string `yaml:"writers"`
	LoggerLevel    string `yaml:"logger_level"`
	LoggerFile     string `yaml:"logger_file"`
	LogHideLineno  bool   `yaml:"log_hide_lineno"`
	LogFormatText  bool   `yaml:"log_format_text"`
	RollingPolicy  string `yaml:"rolling_policy"`
	LogRotateDate  int    `yaml:"log_rotate_date"`
	LogRotateSize  int    `yaml:"log_rotate_size"`
	LogBackupCount int    `yaml:"log_backup_count"`
}

// logger
//
//	is the global variable for the object of lager.logger
var logger lager.Logger

// logFilePath log file path
var logFilePath string

// PassLagerDefinition
//
//	is having the information about logging
var PassLagerDefinition *PassLagerCfg = DefaultLagerDefinition()

// initialize
//
//	Build constructs a *Lager.logger with the configured parameters.
func initialize(writers, loggerLevel, loggerFile, rollingPolicy string, logFormatText, logHideLineno bool,
	LogRotateDate, LogRotateSize, LogBackupCount int) {
	lag := &Lager{
		Writers:        writers,
		LoggerLevel:    loggerLevel,
		LoggerFile:     loggerFile,
		LogHideLineno:  logHideLineno,
		LogFormatText:  logFormatText,
		RollingPolicy:  rollingPolicy,
		LogRotateDate:  LogRotateDate,
		LogRotateSize:  LogRotateSize,
		LogBackupCount: LogBackupCount,
	}

	logger = newLog(lag)
	initLogRotate(logFilePath, lag)
}

// newLog
//
//	new log
func newLog(lag *Lager) lager.Logger {
	checkPassLagerDefinition(lag)

	if lag.LoggerFile != "" {
		if !strings.Contains(lag.Writers, "file") {
			panic(fmt.Errorf("[ logger_file ] is not empty, but writers does not contain [ file ], please check the configuration"))
		}
		if filepath.IsAbs(lag.LoggerFile) {
			createLogFile("", lag.LoggerFile)
			logFilePath = filepath.Join("", lag.LoggerFile)
		} else {
			createLogFile(os.Getenv("CHASSIS_HOME"), lag.LoggerFile)
			logFilePath = filepath.Join(os.Getenv("CHASSIS_HOME"), lag.LoggerFile)
		}
	} else {
		if strings.Contains(lag.Writers, "file") {
			panic(fmt.Errorf("[ logger_file ] is empty, but writers contains [ file ], please check the configuration"))
		}
	}

	writers := strings.Split(strings.TrimSpace(lag.Writers), ",")
	if len(strings.TrimSpace(lag.Writers)) == 0 {
		writers = []string{"stdout"}
	}
	lagerInit(Config{
		Writers:       writers,
		LoggerLevel:   lag.LoggerLevel,
		LoggerFile:    logFilePath,
		LogHideLineno: lag.LogHideLineno,
		LogFormatText: lag.LogFormatText,
	})

	logger := newLogger(lag.LoggerFile, lag.LogFormatText, lag.LogHideLineno)
	return logger
}

// checkPassLagerDefinition
//
//	check pass lager definition
func checkPassLagerDefinition(lag *Lager) {
	if lag.LoggerLevel == "" {
		lag.LoggerLevel = "DEBUG"
	}

	if lag.LoggerFile != "" {
		if lag.RollingPolicy == "" {
			log.Println("RollingPolicy is empty, use default policy[size]")
			lag.RollingPolicy = RollingPolicySize
		} else if lag.RollingPolicy != RollingPolicyDaily && lag.RollingPolicy != RollingPolicySize {
			log.Printf("RollingPolicy is error, RollingPolicy=%s, use default policy[size].", lag.RollingPolicy)
			lag.RollingPolicy = RollingPolicySize
		}

		if lag.LogRotateDate <= 0 || lag.LogRotateDate > 10 {
			lag.LogRotateDate = LogRotateDate
		}

		if lag.LogRotateSize <= 0 || lag.LogRotateSize > 64 {
			lag.LogRotateSize = LogRotateSize
		}

		if lag.LogBackupCount < 0 || lag.LogBackupCount > 100 {
			lag.LogBackupCount = LogBackupCount
		}
	}

}

// createLogFile create log file
func createLogFile(localPath, outputpath string) {
	_, err := os.Stat(strings.Replace(filepath.Dir(filepath.Join(localPath, outputpath)), "\\", "/", -1))
	if err != nil && os.IsNotExist(err) {
		errDir := os.MkdirAll(strings.Replace(filepath.Dir(filepath.Join(localPath, outputpath)), "\\", "/", -1), os.ModePerm)
		if errDir != nil {
			panic(errDir)
		}
	} else if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(strings.Replace(filepath.Join(localPath, outputpath), "\\", "/", -1), os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

// InitWithFile
//
//	readPassLagerConfigFile is unmarshal the paas lager configuration file(log.yaml)
func InitWithFile(lagerFile string) error {
	if lagerFile == "" {
		log.Printf("log config file is empty, use default config: `%s`\n", marshalDefinition())
		return initPassLager()
	}

	passLagerDef := PassLagerCfg{}
	yamlFile, err := ioutil.ReadFile(lagerFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v, use default config: `%s`\n", err, marshalDefinition())
		return initPassLager()
	}

	err = yaml.Unmarshal(yamlFile, &passLagerDef)
	if err != nil {
		log.Printf("Unmarshal: %v, use default config: `%s`\n", err, marshalDefinition())
		return initPassLager()
	}

	PassLagerDefinition = &passLagerDef
	return initPassLager()
}

// InitWithConfig
//
//	use pass lager configuration
func InitWithConfig(passLagerDef *PassLagerCfg) error {
	PassLagerDefinition = passLagerDef
	return initPassLager()
}

// DefaultLagerDefinition
//
//	use default lager definition
func DefaultLagerDefinition() *PassLagerCfg {
	cfg := PassLagerCfg{
		Writers:        "stdout",
		LoggerLevel:    "DEBUG",
		LoggerFile:     "",
		LogHideLineno:  false,
		LogFormatText:  true,
		RollingPolicy:  RollingPolicySize,
		LogRotateDate:  1,
		LogRotateSize:  10,
		LogBackupCount: 7,
	}

	return &cfg
}

func initPassLager() error {
	initialize(PassLagerDefinition.Writers, PassLagerDefinition.LoggerLevel,
		PassLagerDefinition.LoggerFile,
		PassLagerDefinition.RollingPolicy,
		PassLagerDefinition.LogFormatText, PassLagerDefinition.LogHideLineno,
		PassLagerDefinition.LogRotateDate, PassLagerDefinition.LogRotateSize, PassLagerDefinition.LogBackupCount)

	return nil
}

func marshalDefinition() string {
	data, _ := json.Marshal(PassLagerDefinition)
	return string(data)
}
