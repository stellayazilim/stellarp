package realm

import (
	"StellaRP/modules/identity/domain/realm/Entity"
	"StellaRP/modules/identity/domain/realm/ValueObject"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRealm_NewRealm_ShouldCreateRealm(t *testing.T) {

	g := NewGomegaWithT(t)

	r := NewRealm(
		ValueObject.NewRealmId(),
		"",
		[]Entity.IRealmUser{},
	)

	g.Expect(r).ShouldNot(BeNil())
}

func TestRealm_Realm_ShouldGetId(t *testing.T) {
	g := NewGomegaWithT(t)
	i := ValueObject.NewRealmId()
	r := NewRealm(
		i, ",", nil,
	)

	g.Expect(r.GetId()).To(Equal(i))
}

func TestRealm_Realm_ShouldReturnTrueIfEqual(t *testing.T) {
	g := NewGomegaWithT(t)
	r := NewRealm(ValueObject.NewRealmId(), "", nil)

	g.Expect(r.IsEqual(r)).Should(BeTrue())
}

func TestRealm_Realm_ShouldReturnFalseIfNotEqual(t *testing.T) {
	g := NewGomegaWithT(t)

	r := NewRealm(ValueObject.NewRealmId(), "", nil)
	r2 := NewRealm(ValueObject.NewRealmId(), "", nil)

	g.Expect(r2.IsEqual(r)).Should(BeFalse())
}

func TestRealm_Realm_ShouldGetName(t *testing.T) {
	g := NewGomegaWithT(t)
	n := "realm name"
	r := NewRealm(ValueObject.NewRealmId(), n, nil)

	g.Expect(r.GetName()).To(Equal(n))
}

func TestRealm_Realm_ShouldSetName(t *testing.T) {
	g := NewGomegaWithT(t)
	n := "realm name"
	r := NewRealm(ValueObject.NewRealmId(), n, nil)

	n2 := "new realm name"
	r.SetName(n2)

	g.Expect(r.GetName()).To(Equal(n2))
}

func TestRealm_Realm_ShouldGetUsers(t *testing.T) {
	g := NewGomegaWithT(t)

	u := []Entity.IRealmUser{
		Entity.NewRealmUser(ValueObject.NewRealmUserId(), "", "", "", ""),
	}

	r := NewRealm(ValueObject.NewRealmId(), "", u)

	g.Expect(r.GetUsers()).Should(Equal(u))
}

func TestRealm_Realm_ShouldSetUsers(t *testing.T) {
	g := NewGomegaWithT(t)
	u := []Entity.IRealmUser{}
	u2 := []Entity.IRealmUser{
		Entity.NewRealmUser(ValueObject.NewRealmUserId(), "", "", "", ""),
	}

	r := NewRealm(ValueObject.NewRealmId(), "", u)

	r.SetUsers(u2)

	g.Expect(r.GetUsers()).ShouldNot(Equal(u))
	g.Expect(r.GetUsers()).Should(Equal(u2))
}
