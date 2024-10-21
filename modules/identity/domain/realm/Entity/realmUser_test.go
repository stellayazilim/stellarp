package Entity

import (
	"StellaRP/modules/identity/domain/realm/ValueObject"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRealmUser_ShouldCreateNewRealmUser(t *testing.T) {

	g := NewGomegaWithT(t)

	r := NewRealmUser(ValueObject.NewRealmUserId(), "", "", "", "")

	g.Expect(r).ToNot(BeNil())
}

func TestRealmUser_ShouldCompareEqualityBetweenTwoRealmUser(t *testing.T) {
	g := NewGomegaWithT(t)

	r := NewRealmUser(ValueObject.NewRealmUserId(), "", "", "", "")

	r2 := NewRealmUser(ValueObject.NewRealmUserId(), "", "", "", "")

	g.Expect(r.IsEqual(r)).To(BeTrue())
	g.Expect(r.IsEqual(r2)).NotTo(BeTrue())
}

func TestRealmUser_ShouldGetId(t *testing.T) {
	g := NewGomegaWithT(t)
	i := ValueObject.NewRealmUserId()

	r := NewRealmUser(i, "", "", "", "")

	g.Expect(r.GetId()).To(Equal(i))
}

func TestRealmUser_ShouldGetEmail(t *testing.T) {
	g := NewGomegaWithT(t)
	e := "test@test.com"
	r := NewRealmUser(ValueObject.NewRealmUserId(), e, "", "", "")

	g.Expect(r.GetEmail()).To(Equal(e))
}

func TestRealmUser_ShouldSetEmail(t *testing.T) {
	g := NewGomegaWithT(t)
	e := "test@test.com"
	e2 := "test2@test.com"
	r := NewRealmUser(ValueObject.NewRealmUserId(), e, "", "", "")

	r.SetEmail(e2)

	g.Expect(r.GetEmail()).To(Equal(e2))
}

func TestRealmUser_ShouldGetFirstName(t *testing.T) {
	g := NewGomegaWithT(t)
	n := "first name"

	r := NewRealmUser(ValueObject.NewRealmUserId(), "", n, "", "")
	g.Expect(r.GetFirstName()).To(Equal(n))
}

func TestRealmUser_ShouldSetFirstName(t *testing.T) {
	g := NewGomegaWithT(t)
	n := "first name"
	n2 := "second name"

	r := NewRealmUser(ValueObject.NewRealmUserId(), "", n, "", "")

	r.SetFirstName(n2)
	g.Expect(r.GetFirstName()).To(Equal(n2))
}

func TestRealmUser_ShouldGetLastName(t *testing.T) {
	g := NewGomegaWithT(t)
	n := "last name"
	r := NewRealmUser(ValueObject.NewRealmUserId(), "", "", n, "")

	g.Expect(r.GetLastName()).To(Equal(n))
}

func TestRealmUser_ShouldSetLastName(t *testing.T) {
	g := NewGomegaWithT(t)
	n := "last name"
	n2 := "second name"
	r := NewRealmUser(ValueObject.NewRealmUserId(), "", "", n, "")

	r.SetLastName(n2)
	g.Expect(r.GetLastName()).To(Equal(n2))
}

func TestRealmUser_ShouldGetPassword(t *testing.T) {
	g := NewGomegaWithT(t)
	p := "password"

	r := NewRealmUser(ValueObject.NewRealmUserId(), "", "", "", p)

	g.Expect(r.GetPassword()).To(Equal(p))
}

func TestRealmUser_ShouldSetPassword(t *testing.T) {
	g := NewGomegaWithT(t)
	p := "password"
	p2 := "password2"
	r := NewRealmUser(ValueObject.NewRealmUserId(), "", "", "", p)

	r.SetPassword(p2)
	g.Expect(r.GetPassword()).To(Equal(p2))
}
