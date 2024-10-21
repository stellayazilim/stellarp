package client

import (
	"StellaRP/modules/identity/domain/client/Entities"
	"StellaRP/modules/identity/domain/client/ValueObjects"
)

type IClient interface {
	IsEqual(other IClient) bool
	GetId() ValueObjects.ClientId
	GetName() string
	SetName(n string)
	GetDescription() string
	SetDescription(d string)
	GetScopes() []Entities.IScope
}

type Client struct {
	Id          ValueObjects.ClientId
	Name        string
	Description string
	Scopes      []Entities.IScope
}

func (c *Client) IsEqual(other IClient) bool {
	return c.Id.IsEqual(other.GetId())
}

func (c *Client) GetId() ValueObjects.ClientId {
	return c.Id
}

func (c *Client) GetName() string {
	return c.Name
}

func (c *Client) SetName(n string) {
	c.Name = n
}

func (c *Client) GetDescription() string {
	return c.Description
}

func (c *Client) GetScopes() []Entities.IScope {
	return c.Scopes
}

func (c *Client) SetDescription(d string) {
	c.Description = d
}

func NewClient(
	id ValueObjects.ClientId,
	name string,
	description string,
	scopes []Entities.IScope,
) IClient {
	return &Client{
		Id:          id,
		Name:        name,
		Description: description,
		Scopes:      scopes,
	}
}
