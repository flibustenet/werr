package lib

import (
	"database/sql"
	"fmt"

	"go.flibuste.net/werr"
)

func Lili() error {
	err := Lolo()
	//return werr.Wrapf(err, "in my lib %s", "ici")
	return fmt.Errorf("in my lib : %w", werr.Wrap(err))
}

func Lolo() error {
	return sql.ErrNoRows
}
