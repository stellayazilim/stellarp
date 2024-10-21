package ValueObject

import (
	"github.com/lucsky/cuid"
)

type IUserId interface {
	GetValue() string
}

type userId struct {
	value string
}

func (u *userId) GetValue() string {
	return u.value
}

func ParseUserId(value string) (IUserId, error) {
	if err := cuid.IsCuid(value); err != nil {
		return nil, err
	}

	return &userId{value}, nil
}

func NewUserId() IUserId {
	return &userId{
		value: cuid.New(),
	}
}
