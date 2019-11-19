package werr

import "database/sql"

func Op() {
	Wrapf(sql.ErrNoRows, "ok", "ça marche")
}
