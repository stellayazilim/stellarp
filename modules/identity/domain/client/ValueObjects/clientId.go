package ValueObjects

import (
	"github.com/lucsky/cuid"
	"reflect"
)

type ClientId struct {
	value string
}

func (c ClientId) IsEqual(other ClientId) bool {
	return reflect.DeepEqual(c, other)
}

func (c ClientId) GetValue() string {
	return c.value
}

func ParseClientId(value string) (ClientId, error) {

	if err := cuid.IsCuid(value); err != nil {
		return ClientId{}, err
	}

	return ClientId{
		value: value,
	}, nil
}
func NewClientId() ClientId {
	return ClientId{
		value: cuid.New(),
	}
}
