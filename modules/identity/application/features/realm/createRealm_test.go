package realmFeature

import (
	"context"
	"github.com/mehdihadeli/go-mediatr"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCreateRealm_ShouldRegisterCreateRealmHandler(t *testing.T) {

	g := NewGomegaWithT(t)

	err := RegisterCreateRealmHandler()

	g.Expect(err).Should(BeNil())
}

func TestCreateRealmHandler_Hanndle_ShouldHandleRequest(t *testing.T) {
	g := NewGomegaWithT(t)
	_ = RegisterCreateRealmHandler()
	cmd := &CreateRealmCommand{
		Name: "stella",
	}

	resp, err := mediatr.Send[*CreateRealmCommand, *CreateRealmResponse](context.Background(), cmd)

	g.Expect(err).Should(BeNil())
	g.Expect(resp).ShouldNot(BeNil())
}
