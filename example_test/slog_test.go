package example_test

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/bar-counter/slog/lager"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func Test_DefaultLagerDefinition(t *testing.T) {
	currentFolderPath, err := getCurrentFolderPath()
	if err != nil {
		t.Fatal(err)
	}

	logPath := filepath.Join(currentFolderPath, "testdata", "chassis.log")
	lagerDefinition := slog.DefaultLagerDefinition()
	lagerDefinition.Writers = "stdout,file"
	lagerDefinition.LoggerFile = "testdata/chassis.log"
	err = slog.InitWithConfig(lagerDefinition)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		slog.Infof("Hi %s, system is starting up ...", "paas-bot")
		slog.Info("check-info", lager.Data{
			"info": "something",
		})

		slog.Debug("check-info", lager.Data{
			"info": "something",
		})

		slog.Warn("failed-to-do-something", lager.Data{
			"info": "something",
		})

		err = fmt.Errorf("this is an error")
		slog.Error("failed-to-do-something", err)

		slog.Info("shutting-down")
		t.Logf("~> mock _slog")
		// do _slog
		t.Logf("~> do _slog")
		// verify _slog
		assert.Equal(t, "", "")
	}

	time.Sleep(5 * time.Second)

	assert.True(t, pathExistsFast(logPath))
}

func TestHideLineno(t *testing.T) {
	t.Logf("~> mock HideLineno")
	// mock HideLineno
	currentFolderPath, err := getCurrentFolderPath()
	if err != nil {
		t.Fatal(err)
	}

	logPath := filepath.Join(currentFolderPath, "testdata", "hideLineno.log")
	lagerDefinition := slog.DefaultLagerDefinition()
	lagerDefinition.Writers = "stdout,file"
	lagerDefinition.LogHideLineno = true
	lagerDefinition.LoggerFile = "testdata/hideLineno.log"
	err = slog.InitWithConfig(lagerDefinition)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		slog.Infof("Hi %s, system is starting up ...", "paas-bot")
		slog.Info("check-info", lager.Data{
			"info": "something",
		})

		slog.Debug("check-info", lager.Data{
			"info": "something",
		})

		slog.Warn("failed-to-do-something", lager.Data{
			"info": "something",
		})

		err = fmt.Errorf("this is an error")
		slog.Error("failed-to-do-something", err)

		slog.Info("shutting-down")
		t.Logf("~> mock _slog")
		// do _slog
		t.Logf("~> do _slog")
		// verify _slog
		assert.Equal(t, "", "")
	}
	t.Logf("~> do HideLineno")
	// do HideLineno

	// verify HideLineno
	time.Sleep(5 * time.Second)

	assert.True(t, pathExistsFast(logPath))
}

func TestNoFileOut(t *testing.T) {
	t.Logf("~> mock NoFileOut")
	// mock NoFileOut
	currentFolderPath, err := getCurrentFolderPath()
	if err != nil {
		t.Fatal(err)
	}

	logsPath := filepath.Join(currentFolderPath, "testdata", "logs")
	lagerDefinition := slog.DefaultLagerDefinition()
	err = slog.InitWithConfig(lagerDefinition)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("~> do NoFileOut")
	// do NoFileOut
	for i := 0; i < 10; i++ {
		slog.Infof("Hi %s, system is starting up ...", "paas-bot")
		slog.Info("check-info", lager.Data{
			"info": "something",
		})

		slog.Debug("check-info", lager.Data{
			"info": "something",
		})

		slog.Warn("failed-to-do-something", lager.Data{
			"info": "something",
		})

		err = fmt.Errorf("this is an error")
		slog.Error("failed-to-do-something", err)

		slog.Info("shutting-down")
		t.Logf("~> mock _slog")
		// do _slog
		t.Logf("~> do _slog")
		// verify _slog
		assert.Equal(t, "", "")
	}

	// verify NoFileOut
	time.Sleep(5 * time.Second)

	assert.False(t, pathExistsFast(logsPath))
}

func Test_yml_slog(t *testing.T) {
	// mock _slog
	currentFolderPath, err := getCurrentFolderPath()
	if err != nil {
		t.Fatal(err)
	}

	logPath := filepath.Join(currentFolderPath, "testdata", "foo.log")

	err = slog.InitWithFile(filepath.Join(currentFolderPath, "log.yaml"))
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		slog.Infof("Hi %s, system is starting up ...", "paas-bot")
		slog.Info("check-info", lager.Data{
			"info": "something",
		})

		slog.Debug("check-info", lager.Data{
			"info": "something",
		})

		slog.Warn("failed-to-do-something", lager.Data{
			"info": "something",
		})

		err = fmt.Errorf("this is an error")
		slog.Error("failed-to-do-something", err)

		slog.Info("shutting-down")
		t.Logf("~> mock _slog")
		// do _slog
		t.Logf("~> do _slog")
		// verify _slog
		assert.Equal(t, "", "")
	}

	time.Sleep(5 * time.Second)

	assert.Truef(t, pathExistsFast(logPath), "want %s to exist", logPath)

}

func TestLogSingleLineJson(t *testing.T) {
	// mock LogJson

	currentFolderPath, err := getCurrentFolderPath()
	if err != nil {
		t.Fatal(err)
	}
	jsonLogPath := filepath.Join(currentFolderPath, "testdata", "single_line_json.json")
	if pathExistsFast(jsonLogPath) {
		errRm := os.Remove(jsonLogPath)
		if errRm != nil {
			t.Fatal(errRm)
		}
	}

	slogCfg := slog.PassLagerCfg{
		Writers:        "file,stdout",
		LoggerLevel:    "INFO",
		LoggerFile:     "testdata/single_line_json.json",
		LogHideLineno:  false,
		LogFormatText:  false,
		RollingPolicy:  "size",
		LogRotateDate:  1,
		LogRotateSize:  8,
		LogBackupCount: 7,
	}
	err = slog.InitWithConfig(&slogCfg)
	if err != nil {
		t.Fatal(err)
	}

	slog.Infof("one %v", slogCfg)

	time.Sleep(5 * time.Second)

	var logContent lager.LogContent

	err = readFileAsJson(jsonLogPath, &logContent)
	if err != nil {
		t.Fatalf("want read json out log as one line err: %v", err)
	}

	assert.Equal(t, "INFO", logContent.Level)
}

func TestPanicConfigErrorByWriters(t *testing.T) {
	// mock TestPanicConfigErrorByWriters

	errString := "[ logger_file ] is empty, but writers contains [ file ], please check the configuration"

	if !assert.PanicsWithError(t, errString, func() {
		// do TestPanicConfigErrorByWriters
		lagerDefinition := slog.DefaultLagerDefinition()
		lagerDefinition.Writers = "stdout,file"
		_ = slog.InitWithConfig(lagerDefinition)
	}) {
		// verify TestPanicConfigErrorByWriters
		t.Fatalf("TestPanicConfigErrorByWriters should panic")
	}
}

func TestPanicConfigErrorByLogFile(t *testing.T) {
	// mock TestPanicConfigErrorByLogFile

	errString := "[ logger_file ] is not empty, but writers does not contain [ file ], please check the configuration"

	if !assert.PanicsWithError(t, errString, func() {
		// do TestPanicConfigErrorByLogFile
		lagerDefinition := slog.DefaultLagerDefinition()
		lagerDefinition.LoggerFile = "testdata/bar.log"
		_ = slog.InitWithConfig(lagerDefinition)
	}) {
		// verify TestPanicConfigErrorByLogFile
		t.Fatalf("TestPanicConfigErrorByLogFile should panic")
	}
}
