package main

import (
	"database/sql"
	"errors"
	"fmt"

	"go.flib.fr/werr"
	"go.flib.fr/werr/test/lib"
)

func a() error {
	e := b()
	return fmt.Errorf("from a : %w", e)
}
func b() error {
	e := c()
	if errors.Is(e, sql.ErrNoRows) {
		return werr.Wrapf(e, "SQLNOROWS in B")
	}
	return fmt.Errorf("from b : %w", werr.Stack(e))
}
func c() error {
	e := d()
	return e
}
func d() error {
	e := lib.Lili()
	if errors.Is(e, sql.ErrNoRows) {
		return werr.Wrapf(e, "SQLNOROWS in D")
	}
	return fmt.Errorf("from d : %w", werr.Stack(e))
}

func main() {
	e := a()
	fmt.Println(werr.PrintSkip(e, ""))
}
