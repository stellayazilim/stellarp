package realmFeature

import (
	"StellaRP/modules/identity/application/interfaces"
	"context"
	"fmt"
	"github.com/mehdihadeli/go-mediatr"
	_ "github.com/mehdihadeli/go-mediatr"
)

type (
	CreateRealmCommand struct {
		Name string
	}
	CreateRealmResponse struct{}
	createRealmHandler  struct {
		realmStore interfaces.IRealmStore
	}
)

func (c *createRealmHandler) Handle(ctx context.Context, command *CreateRealmCommand) (*CreateRealmResponse, error) {

	fmt.Println(command.Name)
	return &CreateRealmResponse{}, nil
}

func RegisterCreateRealmHandler() error {
	handler := &createRealmHandler{}
	return mediatr.RegisterRequestHandler[*CreateRealmCommand, *CreateRealmResponse](handler)

}
