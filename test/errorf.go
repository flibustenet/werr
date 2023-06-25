package main

import (
	"go.flibuste.net/werr"
)

func errorf() error {
	return werr.Errorf("errorf: %v", fail())
}
