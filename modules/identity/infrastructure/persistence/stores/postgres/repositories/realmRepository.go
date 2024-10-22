package repositories

import (
	"StellaRP/modules/identity/domain/realm"
	"StellaRP/modules/identity/infrastructure/persistence/interfaces"
	"context"

	"github.com/jackc/pgx/v5"
)

type (
	realmStore struct {
		conn *pgx.Conn
	}
)

func (r *realmStore) FindRealms() []realm.IRealm {
	//TODO implement me
	panic("implement me")
}

func (r *realmStore) FindOneRealm(entity realm.IRealm) realm.IRealm {
	//TODO implement me
	panic("implement me")
}

func (r *realmStore) CreateRealm(entity realm.IRealm) error {
	//TODO implement me
	panic("implement me")
}

func (r *realmStore) UpdateRealm(entity realm.IRealm) error {
	//TODO implement me
	panic("implement me")
}

func (r *realmStore) DeleteRealm(entity realm.IRealm) error {
	//TODO implement me
	panic("implement me")
}

func (r *realmStore) Create(ctx context.Context, realm realm.IRealm) error {

	return nil
}

func UsePgRealmStore(conn *pgx.Conn) interfaces.IRealmStore {
	return &realmStore{
		conn: conn,
	}
}
