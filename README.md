[![golang-ci](https://github.com/bar-counter/slog/actions/workflows/golang-ci.yml/badge.svg)](https://github.com/bar-counter/slog/actions/workflows/golang-ci.yml)
[![go mod version](https://img.shields.io/github/go-mod/go-version/bar-counter/slog?label=go.mod)](https://github.com/bar-counter/slog)
[![GoDoc](https://godoc.org/github.com/bar-counter/slog?status.png)](https://godoc.org/github.com/bar-counter/slog/)
[![GoReportCard](https://goreportcard.com/badge/github.com/bar-counter/slog)](https://goreportcard.com/report/github.com/bar-counter/slog)
[![codecov](https://codecov.io/gh/bar-counter/slog/branch/main/graph/badge.svg)](https://codecov.io/gh/bar-counter/slog)
[![github release](https://img.shields.io/github/v/release/bar-counter/slog?style=social)](https://github.com/bar-counter/slog/releases)

## for what

- this project used to github golang lib project

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

