package lib

import (
	"database/sql"

	"go.flib.fr/werr"
)

func Lili() error {
	err := Lolo()
	return werr.Wrapf(err, "in my lib")
}

func Lolo() error {
	return sql.ErrNoRows
}
