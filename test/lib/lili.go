package lib

import (
	"database/sql"
	"fmt"

	"go.flib.fr/werr"
)

func Lili() error {
	err := Lolo()
	//return werr.Wrapf(err, "in my lib %s", "ici")
	return fmt.Errorf("in my lib : %w", werr.Stack(err))
}

func Lolo() error {
	return sql.ErrNoRows
}
