package Entities

import (
	"StellaRP/modules/identity/domain/client/ValueObjects"
)

type IScope interface {
	GetId() ValueObjects.ScopeId
	GetName() string
	SetName(name string)
	GetDescription() string
	SetDescription(description string)
}
type Scope struct {
	id          ValueObjects.ScopeId
	name        string
	description string
}

func (s *Scope) GetId() ValueObjects.ScopeId {
	return s.id
}

func (s *Scope) GetName() string {
	return s.name
}

func (s *Scope) SetName(name string) {
	s.name = name
}

func (s *Scope) GetDescription() string {
	return s.description
}

func (s *Scope) SetDescription(description string) {
	s.description = description
}

func NewScope(
	id ValueObjects.ScopeId,
	name string,
	description string,
) IScope {
	return &Scope{
		id:          id,
		name:        name,
		description: description,
	}
}
