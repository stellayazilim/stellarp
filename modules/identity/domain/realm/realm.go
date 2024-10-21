package realm

import (
	"StellaRP/modules/identity/domain/realm/Entity"
	"StellaRP/modules/identity/domain/realm/ValueObject"
)

type IRealm interface {
	GetId() ValueObject.IRealmId
	IsEqual(other IRealm) bool
	GetName() string
	SetName(name string)
	GetUsers() []Entity.IRealmUser
	SetUsers(users []Entity.IRealmUser)
}

type realm struct {
	id    ValueObject.IRealmId
	name  string
	users []Entity.IRealmUser
}

func (r *realm) GetId() ValueObject.IRealmId {
	return r.id
}

func (r *realm) IsEqual(other IRealm) bool {

	id := r.GetId()

	return id.IsEqual(other.GetId())
}

func (r *realm) GetName() string {
	return r.name
}

func (r *realm) SetName(value string) {
	r.name = value
}

func (r *realm) GetUsers() []Entity.IRealmUser {
	return r.users
}

func (r *realm) SetUsers(users []Entity.IRealmUser) {
	r.users = users
}

func NewRealm(
	id ValueObject.IRealmId,
	name string,
	users []Entity.IRealmUser,
) IRealm {
	return &realm{
		id:    id,
		name:  name,
		users: users,
	}
}
