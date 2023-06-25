package main

import "errors"

func fail() error {
	return errors.New("fail")
}
