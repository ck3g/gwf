package gwf

import (
	"database/sql"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func (g *GWF) OpenDB(dbType, dsn string) (*sql.DB, error) {
	if dbType == "postgres" || dbType == "postresql" {
		dbType = "pgx"
	}

	db, err := sql.Open(dbType, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
