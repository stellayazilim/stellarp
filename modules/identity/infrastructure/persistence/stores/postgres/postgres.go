package postgres

import (
	"StellaRP/modules/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type IConnection interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

func UsePostgres(cfg *config.Config) (IConnection, error) {

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.GetPostgresUser(),
		cfg.GetPostgresPassword(),
		cfg.GetPostgresHost(),
		5432,
		cfg.GetPostgresDatabase(),
	)

	return pgx.Connect(context.Background(), connString)
}
