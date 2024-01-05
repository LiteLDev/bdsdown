package main

import (
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("bds-down")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05} %{level} > %{color:reset}%{message}`,
)

func init() {
	stdout := logging.NewLogBackend(os.Stdout, "", 0)
	stdoutFormatter := logging.NewBackendFormatter(stdout, format)
	logging.SetBackend(stdoutFormatter)
}
