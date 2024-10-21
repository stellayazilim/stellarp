package postgres

import (
	"StellaRP/modules/identity/infrastructure/persistence/postgres/repositories"

	"github.com/jackc/pgx/v5"
)

type Store struct {
	RealmRepository repositories.IRealmRepository
}

func UsePostgres(conn *pgx.Conn) *Store {
	return &Store{
		RealmRepository: repositories.NewRealmRepository(conn),
	}
}
