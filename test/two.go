package main

import (
	"go.flibuste.net/werr"
)

func two() error {
	return werr.Wrapf(one(), "two")
}
func one() error {
	return werr.Wrapf(fail(), "one")
}
