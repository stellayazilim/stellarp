package client

import (
	"StellaRP/modules/identity/domain/client/Entities"
	"StellaRP/modules/identity/domain/client/ValueObjects"
	. "github.com/onsi/gomega"
	"testing"
)

func TestClient_NewClientShouldCreateClient(t *testing.T) {

	g := NewGomegaWithT(t)

	i := ValueObjects.NewClientId()

	s := []Entities.IScope{
		Entities.NewScope(ValueObjects.NewScopeId(), "", ""),
	}

	c := NewClient(i, "", "", s)

	g.Expect(c).NotTo(BeNil())
	g.Expect(c).To(BeAssignableToTypeOf(&Client{}))
}

func TestClient_IsEqualShouldBeTrueIfClientsAreEqual(t *testing.T) {
	g := NewGomegaWithT(t)

	c1 := NewClient(ValueObjects.NewClientId(), "", "", []Entities.IScope{})
	c2 := NewClient(ValueObjects.NewClientId(), "", "", []Entities.IScope{})

	b := c1.IsEqual(c2)
	b2 := c1.IsEqual(c1)

	g.Expect(b).NotTo(BeTrue())
	g.Expect(b2).To(BeTrue())

}

func TestClient_ShouldGetName(t *testing.T) {
	g := NewGomegaWithT(t)

	i := ValueObjects.NewClientId()

	n := "test client"

	c := NewClient(i, n, "", []Entities.IScope{})

	g.Expect(c.GetName()).To(Equal(n))
}

func TestClient_ShouldSetName(t *testing.T) {
	g := NewGomegaWithT(t)
	i := ValueObjects.NewClientId()
	n := "test client"
	en := "new test client"

	c := NewClient(i, n, "", []Entities.IScope{})

	c.SetName(en)

	g.Expect(c.GetName()).To(Equal(en))

}

func TestClient_ShouldGetDescription(t *testing.T) {
	g := NewGomegaWithT(t)
	d := "client description"

	c := NewClient(ValueObjects.NewClientId(), "", d, []Entities.IScope{})

	g.Expect(c.GetDescription()).To(Equal(d))
}

func TestClient_ShouldSetDescription(t *testing.T) {
	g := NewGomegaWithT(t)
	d := "client description"
	d2 := "new client description"

	c := NewClient(ValueObjects.NewClientId(), "", d, []Entities.IScope{})

	c.SetDescription(d2)

	g.Expect(c.GetDescription()).To(Equal(d2))
}

func TestClient_ShouldGetValue(t *testing.T) {
	g := NewGomegaWithT(t)

	scopes := []Entities.IScope{}

	c := NewClient(ValueObjects.NewClientId(), "", "", scopes)

	g.Expect(c.GetScopes()).To(Equal(scopes))
}
