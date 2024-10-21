package repositories

import (
	"StellaRP/modules/identity/domain/realm"
	"context"

	"github.com/jackc/pgx/v5"
)

type (
	IRealmRepository interface{}

	RealmRepository struct {
		conn *pgx.Conn
	}
)

func (r RealmRepository) Create(ctx context.Context, realm realm.Realm) error {

	_, _ = r.conn.Exec(ctx, `
	INSERT INTO realms (id, name) VALUES ($1, $1);`,
		realm.GetId(),
		realm.GetName(),
	)

	return nil
}

func NewRealmRepository(conn *pgx.Conn) IRealmRepository {
	return &RealmRepository{
		conn: conn,
	}
}
