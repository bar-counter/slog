[![golang-ci](https://github.com/bar-counter/slog/actions/workflows/golang-ci.yml/badge.svg)](https://github.com/bar-counter/slog/actions/workflows/golang-ci.yml)
[![go mod version](https://img.shields.io/github/go-mod/go-version/bar-counter/slog?label=go.mod)](https://github.com/bar-counter/slog)
[![GoDoc](https://godoc.org/github.com/bar-counter/slog?status.png)](https://godoc.org/github.com/bar-counter/slog/)
[![GoReportCard](https://goreportcard.com/badge/github.com/bar-counter/slog)](https://goreportcard.com/report/github.com/bar-counter/slog)
[![codecov](https://codecov.io/gh/bar-counter/slog/branch/main/graph/badge.svg)](https://codecov.io/gh/bar-counter/slog)
[![github release](https://img.shields.io/github/v/release/bar-counter/slog?style=social)](https://github.com/bar-counter/slog/releases)

## for what

- this project used to golang log management

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/bridgewwater/gin-api-swagger-temple)](https://github.com/bridgewwater/gin-api-swagger-temple/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

## depends

in go mod project

```bash
# warning use privte git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "http://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q http://github.com/bar-counter/slog.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/bar-counter/slog
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/bar-counter/slog | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## evn

- golang sdk 1.17+

## Features

- [X] easy API to use, `slog.Debug("this is debug")`...
- [X] easy config new `slog.DefaultLagerDefinition()`
- [X] config load by `yaml file`
- [X] support stdout and file
- [X] color stdout support
- [X] rolling policy at file output
  - log_rotate_date: max 10 days, greater than will change to 1, rotate date, coordinate `log_rotate_date: daily`
  - log_rotate_size: max 64M, greater than will change to 10, rotate size，coordinate `rollingPolicy: size`
  - log_backup_count: max 100 files, greater than will change to 7, log system will compress the log file when log reaches rotate set, this set is max file count
- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

- use `slog.DefaultLagerDefinition()`

```go
package main
import (
  "fmt"
  "github.com/bar-counter/slog"
  "testing"
)

func TestMainLog(t *testing.T) {
  lagerDefinition := slog.DefaultLagerDefinition()
  err := slog.InitWithConfig(lagerDefinition)
  if err != nil {
    t.Fatal(err)
  }

  slog.Debug("this is debug")
  slog.Infof("this is info %v", "some info")
  slog.Warn("this is warn")
  slog.Error("this is error", fmt.Errorf("some error"))
}
```

- load with `*.yaml`

```yaml
writers: stdout # file,stdout。`file` will let `logger_file` to file，`stdout` will show at std, most of time use bose
logger_level: DEBUG # DEBUG INFO WARN ERROR FATAL
logger_file: logs/foo.log # "" is not writer log file, and this will cover by env: CHASSIS_HOME
log_format_text: true # format `false` will format json, `true` will show std
rolling_policy: size # rotate policy, can choose as: daily, size. `daily` store as daily，`size` will save as max
log_rotate_date: 1 # max 10 days, greater than will change to 1, rotate date, coordinate `log_rotate_date: daily`
log_rotate_size: 8 # max 64M, greater than will change to 10, rotate size，coordinate `rollingPolicy: size`
log_backup_count: 7 # max 100 files, greater than will change to 7, log system will compress the log file when log reaches rotate set, this set is max file count
```

- use `slog.InitWithFile("log.yaml")`

```go
package main

import (
    "fmt"
	"github.com/bar-counter/slog"
)

func main() {
  err := slog.InitWithFile("log.yaml")
  if err != nil {
    panic(err)
  }

  slog.Debug("this is debug")
  slog.Infof("this is info %v", "some info")
  slog.Warn("this is warn")
  slog.Error("this is error", fmt.Errorf("some error"))
}
```

# dev

```bash
make init dep
```

- test code

```bash
make test
```

add main.go file and run

```bash
# run at env dev
make dev

# run at env ordinary
make run
```

- ci to fast check

```bash
make ci
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```

## use

- use to replace
  `bar-counter/slog` to you code

