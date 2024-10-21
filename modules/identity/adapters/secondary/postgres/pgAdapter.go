package postgres

import (
	"StellaRP/modules/identity/infrastructure/persistence/postgres/repositories"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type IPgAdapter interface{}
type PgAdapter struct {
	RealmRepository repositories.IRealmRepository
}

func NewPgAdapter() IPgAdapter {

	pgUser := os.Getenv("postgres_USER")
	pgPassword := os.Getenv("postgres_PASSWORD")
	pgHost := os.Getenv("postgres_HOST")
	pgDb := os.Getenv("postgres_DB")

	fmt.Println("db env: ", pgUser, pgPassword, pgHost, pgDb)
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:5432/%s", pgUser, pgPassword, pgHost, pgDb))

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)

		os.Exit(1)
	}

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			os.Exit(1)
		}
	}(conn, context.Background())

	return PgAdapter{
		RealmRepository: repositories.NewRealmRepository(conn),
	}
}
