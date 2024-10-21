package ValueObject

import (
	"reflect"

	"github.com/lucsky/cuid"
)

type IRealmId interface {
	GetValue() string
	IsEqual(other IRealmId) bool
}

type RealmId struct {
	value string
}

func (r *RealmId) GetValue() string {
	return r.value
}

func (r *RealmId) IsEqual(other IRealmId) bool {
	return reflect.DeepEqual(r, other)
}

func ParseRealmId(value string) (IRealmId, error) {

	if err := cuid.IsCuid(value); err != nil {
		return nil, err
	}

	return &RealmId{value: value}, nil
}

func NewRealmId() IRealmId {
	return &RealmId{
		value: cuid.New(),
	}
}
