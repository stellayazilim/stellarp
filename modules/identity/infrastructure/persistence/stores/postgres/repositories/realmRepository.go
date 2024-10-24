package repositories

import (
	"StellaRP/modules/identity/domain/realm"
	"StellaRP/modules/identity/infrastructure/persistence/interfaces"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type (
	realmStore struct {
		conn *pgx.Conn
	}
)

func (r *realmStore) Find(ctx context.Context) []realm.IRealm {
	//TODO implement me
	return nil
}

func (r *realmStore) FindOne(ctx context.Context, entity realm.IRealm) realm.IRealm {
	//TODO implement me
	return nil
}

func (r *realmStore) Update(ctx context.Context, entity realm.IRealm) error {
	//TODO implement me
	return nil
}

func (r *realmStore) Delete(ctx context.Context, entity realm.IRealm) error {
	//TODO implement me
	return nil
}

func (r *realmStore) Create(ctx context.Context, realm realm.IRealm) error {
	tx, err := r.conn.Begin(ctx)

	if err != nil {
		return err
	}
	fmt.Println("tx begin")
	if _, err := tx.Exec(ctx /*sql*/, `INSERT INTO realms VALUES (@id, @name)`, pgx.NamedArgs{
		"id":   realm.GetId().GetValue(),
		"name": realm.GetName(),
	}); err != nil {
		return tx.Rollback(ctx)
	}

	return tx.Commit(ctx)
}

func UsePgRealmStore(conn *pgx.Conn) interfaces.IRealmStore {
	return &realmStore{
		conn: conn,
	}
}
