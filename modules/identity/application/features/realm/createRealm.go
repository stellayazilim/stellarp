package realmFeature

import (
	realm "StellaRP/modules/identity/domain/realm"
	"StellaRP/modules/identity/domain/realm/ValueObject"
	"StellaRP/modules/identity/infrastructure/persistence/interfaces"
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

	fmt.Println("handler")
	newRealm := realm.NewRealm(
		ValueObject.NewRealmId(),
		command.Name,
		nil,
	)

	if err := c.realmStore.Create(ctx, newRealm); err != nil {
		fmt.Println("handler", err)
		return nil, err
	}

	return &CreateRealmResponse{}, nil
}

func RegisterCreateRealmHandler(store interfaces.IRealmStore) error {
	handler := &createRealmHandler{
		realmStore: store,
	}
	return mediatr.RegisterRequestHandler[*CreateRealmCommand, *CreateRealmResponse](handler)

}
