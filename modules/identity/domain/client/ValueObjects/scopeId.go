package ValueObjects

import (
	"github.com/lucsky/cuid"
	"reflect"
)

type ScopeId struct {
	value string
}

func (s ScopeId) IsEqual(other ScopeId) bool {
	return reflect.DeepEqual(s, other)
}

func (s ScopeId) GetValue() string {
	return s.value
}

func ParseScopeId(value string) (ScopeId, error) {

	if err := cuid.IsCuid(value); err != nil {
		return ScopeId{}, err
	}
	return ScopeId{
		value: value,
	}, nil
}
func NewScopeId() ScopeId {
	return ScopeId{
		value: cuid.New(),
	}
}
