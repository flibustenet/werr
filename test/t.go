package main

import (
	"database/sql"
	"errors"
	"fmt"
	"runtime/debug"

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
	return fmt.Errorf("from b : %w", werr.Wrap(e))
}
func c() error {
	return errors.New("uuuuuuuuuuuuuuuuuuuu") //werr.New("iiiiiiiiiiiiiiiiiiiiiiiiiiiiii") //Errorf("Kikou : %s", "ici")
	e := d()
	return e
}
func d() error {
	//panic("ici")
	e := lib.Lili()
	//werr.Check(e)

	werr.MustWrapf(e, "panique à bord")
	if errors.Is(e, sql.ErrNoRows) {
		return werr.Wrapf(e, "SQLNOROWS in D")
	}
	return fmt.Errorf("from d : %w", werr.Wrap(e))
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			if err, ok := e.(error); ok {
				werr.Print(err)
			}
		} else {
			fmt.Printf("%s", debug.Stack())
		}
	}()
	e := a()
	fmt.Println("==================================================================================")
	werr.Print(e)
}
