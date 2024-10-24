package postgres

import (
	"StellaRP/modules/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func UsePostgres(cfg *config.Config) (*pgx.Conn, error) {

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.GetPostgresUser(),
		cfg.GetPostgresPassword(),
		cfg.GetPostgresHost(),
		5432,
		cfg.GetPostgresDatabase(),
	)

	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	return conn, nil
}
