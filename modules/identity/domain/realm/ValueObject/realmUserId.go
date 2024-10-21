package ValueObject

import (
	"github.com/lucsky/cuid"
	"reflect"
)

type IRealmUserId interface {
	IsEqual(other IRealmUserId) bool
	GetValue() string
}

type realmUserId struct {
	value string
}

func (r *realmUserId) IsEqual(other IRealmUserId) bool {

	return reflect.DeepEqual(r, other)
}

func (r *realmUserId) GetValue() string {
	return r.value
}

func ParseRealmUserId(value string) (IRealmUserId, error) {

	if err := cuid.IsCuid(value); err != nil {
		return nil, err
	}

	return &realmUserId{
		value: value,
	}, nil
}

func NewRealmUserId() IRealmUserId {
	return &realmUserId{
		value: cuid.New(),
	}
}
