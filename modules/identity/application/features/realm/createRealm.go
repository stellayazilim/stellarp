package realmFeature

import (
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
	createRealmHandler  struct{}
)

func (c *createRealmHandler) Handle(ctx context.Context, command *CreateRealmCommand) (*CreateRealmResponse, error) {

	fmt.Println(command.Name)
	return &CreateRealmResponse{}, nil
}

func RegisterCreateRealmHandler() error {
	handler := &createRealmHandler{}
	return mediatr.RegisterRequestHandler[*CreateRealmCommand, *CreateRealmResponse](handler)

}
