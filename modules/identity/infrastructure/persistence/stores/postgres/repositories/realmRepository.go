package repositories

import (
	"StellaRP/modules/identity/application/interfaces"
	"StellaRP/modules/identity/domain/realm"
	"StellaRP/modules/identity/infrastructure/persistence/stores/postgres"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type (
	realmStore struct {
		conn postgres.IConnection
	}
)

func (r *realmStore) Find(ctx context.Context) []realm.IRealm {
	//TODO implement me
	panic("implement me")
}

func (r *realmStore) FindOne(ctx context.Context, entity realm.IRealm) realm.IRealm {
	//TODO implement me
	panic("implement me")
}

func (r *realmStore) Update(ctx context.Context, entity realm.IRealm) error {
	//TODO implement me
	panic("implement me")
}

func (r *realmStore) Delete(ctx context.Context, entity realm.IRealm) error {
	//TODO implement me
	panic("implement me")
}

func (r *realmStore) Create(ctx context.Context, realm realm.IRealm) error {
	// Begin a transaction
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}

	// Attempt to insert the realm data into the database using named parameters
	_, err = tx.Exec(ctx, `INSERT INTO realms (id, name) VALUES (@id, @name)`, pgx.NamedArgs{
		"id":   realm.GetId().GetValue(),
		"name": realm.GetName(),
	})

	// If there's an error, rollback the transaction and return the error
	if err != nil {
		rollbackErr := tx.Rollback(ctx)
		if rollbackErr != nil {
			return fmt.Errorf("rollback failed: %v, original error: %v", rollbackErr, err)
		}
		return err
	}

	// Commit the transaction if everything is successful
	return tx.Commit(ctx)
}

func UsePgRealmStore(conn postgres.IConnection) interfaces.IRealmStore {
	return &realmStore{
		conn: conn,
	}
}
