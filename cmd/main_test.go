package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"os/exec"
	"testing"
)

func Test_package_main(t *testing.T) {

	cmd := exec.Command(os.Args[0], "-h")
	cmd.Env = append(os.Environ(), "ENV_WEB_AUTO_HOST=true")
	var outStd bytes.Buffer
	cmd.Stdout = &outStd
	var errStd bytes.Buffer
	cmd.Stderr = &errStd
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("-h result: %s\n%s", outStd.String(), errStd.String())

}

func Test_package_main_error(t *testing.T) {
	cmdFail := exec.Command(os.Args[0], "--error.arg")
	cmdFail.Env = append(os.Environ(), "ENV_WEB_AUTO_HOST=true")
	err := cmdFail.Run()
	if e, ok := err.(*exec.ExitError); ok {
		assert.False(t, e.Success())
		assert.Equal(t, 2, e.ExitCode())
		return
	}
	t.Fatalf("Process run with err %v, want os.Exit(2)", err)
}

func TestMainLog(t *testing.T) {
	lagerDefinition := slog.DefaultLagerDefinition()
	err := slog.InitWithConfig(lagerDefinition)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("-> env:ENV_WEB_AUTO_HOST %s", os.Getenv("ENV_WEB_AUTO_HOST"))
	flag.Parse()
	log.Printf("=> now version %v", cliVersion)

	slog.Debug("this is debug")
	slog.Infof("this is info %v", "some info")
	slog.Warn("this is warn")
	slog.Error("this is error", fmt.Errorf("some error"))
}
