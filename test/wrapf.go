package main

import (
	"go.flibuste.net/werr"
)

func wrapf() error {
	return werr.Wrapf(fail(), "wrapf")
}
