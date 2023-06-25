package main

import (
	"go.flibuste.net/werr"
)

func wrap() error {
	return werr.Wrap(fail())
}
