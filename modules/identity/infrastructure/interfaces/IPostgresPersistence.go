package interfaces

import "github.com/jackc/pgx/v5"

type PostgresPersistence interface {
	GetConnection() *pgx.Conn
}
