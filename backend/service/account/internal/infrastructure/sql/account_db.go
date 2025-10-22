package account_db

import (
	"database/sql"

	db "github.com/ilkerciblak/buldum-app/shared/core/infrastructure/sql"
)

func New(db db.DBTX) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db db.DBTX
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}
