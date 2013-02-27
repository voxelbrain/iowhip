package main

import (
	"syscall"
)

const (
	O_DIRECT = syscall.O_DIRECT
	O_DSYNC  = syscall.O_DSYNC
)
