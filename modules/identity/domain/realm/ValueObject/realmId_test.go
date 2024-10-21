package ValueObject

import (
	"github.com/lucsky/cuid"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRealmId_NewRealmId_ShouldCreateNewRealmId(t *testing.T) {

	g := NewWithT(t)
	i := NewRealmId()
	g.Expect(i).NotTo(BeNil())
}

func TestRealmId_ParseRealmId_ShouldParsedWithValidId(t *testing.T) {
	g := NewWithT(t)

	c := cuid.New()

	i, err := ParseRealmId(c)

	g.Expect(err).To(BeNil())

	g.Expect(i).NotTo(BeNil())
	g.Expect(i.GetValue()).To(Equal(c))
}

func TestRealmId_ParseRealmId_ShouldReturnErrorWithInvalidId(t *testing.T) {
	g := NewWithT(t)
	c := "invalid id"
	i, err := ParseRealmId(c)

	g.Expect(i).To(BeNil())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err).To(MatchError("Incorrect format"))
}

func TestRealmId_IsEqual_ShouldReturnTrueWithEqualRealmId(t *testing.T) {
	g := NewWithT(t)

	i := NewRealmId()

	g.Expect(i.IsEqual(i)).To(BeTrue())
}

func TestRealmId_IsEqual_ShouldReturnFalseWithNotEqualRealmId(t *testing.T) {
	g := NewWithT(t)
	i := NewRealmId()
	i2 := NewRealmId()

	g.Expect(i.IsEqual(i2)).To(BeFalse())
}
