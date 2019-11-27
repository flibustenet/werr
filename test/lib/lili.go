package lib

import (
	"database/sql"
	"fmt"

	"go.flib.fr/werr"
)

func Lili() error {
	err := Lolo()
	//return fmt.Errorf("in my lib : %v", err)
	//return fmt.Errorf("in my lib : %w", err) //werr.Stack(err)) //werr.Wrapf(err, "in my lib %s", "ici")
	return fmt.Errorf("in my lib : %w", werr.Stack(err)) //werr.Wrapf(err, "in my lib %s", "ici")
}

func Lolo() error {
	return sql.ErrNoRows
}
