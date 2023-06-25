package main

import "go.flibuste.net/werr"

func fnew() error {
	return werr.New("new")
}
