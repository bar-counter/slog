//go:build !test

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/bar-counter/slog"
	"log"
	"os"
)

const cliVersion = "0.1.2"

var serverPort = flag.String("serverPort", "49002", "http service address")

func main() {
	err := slog.InitWithFile("log.yaml")
	if err != nil {
		panic(err)
	}
	log.Printf("-> env:ENV_WEB_AUTO_HOST %s", os.Getenv("ENV_WEB_AUTO_HOST"))
	flag.Parse()
	log.Printf("-> run serverPort %v", *serverPort)
	log.Printf("=> now version %v", cliVersion)

	slog.Debug("this is debug")
	slog.Infof("this is info %v", "some info")
	slog.Warn("this is warn")
	slog.Error("this is error", fmt.Errorf("some error"))
}
