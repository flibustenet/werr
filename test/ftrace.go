package main

import (
	"go.flibuste.net/werr"
)

func ftrace() error {
	return werr.Trace(fail())
}
