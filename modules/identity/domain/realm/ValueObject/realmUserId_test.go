package ValueObject

import (
	"github.com/lucsky/cuid"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRealmUserId_NewRealmUserId_ShouldCreateNewRealmUserId(t *testing.T) {

	g := NewWithT(t)
	i := NewRealmUserId()
	g.Expect(i).NotTo(BeNil())
}

func TestRealmUserId_ParseRealmUserId_ShouldParsedWithValidId(t *testing.T) {
	g := NewWithT(t)

	c := cuid.New()

	i, err := ParseRealmUserId(c)

	g.Expect(err).To(BeNil())

	g.Expect(i).NotTo(BeNil())
	g.Expect(i.GetValue()).To(Equal(c))
}

func TestRealmUserId_ParseRealmUserId_ShouldReturnErrorWithInvalidId(t *testing.T) {
	g := NewWithT(t)
	c := "invalid id"
	i, err := ParseRealmUserId(c)

	g.Expect(i).To(BeNil())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err).To(MatchError("Incorrect format"))
}

func TestRealmUserId_IsEqual_ShouldReturnTrueWithEqualRealmUserId(t *testing.T) {
	g := NewWithT(t)

	i := NewRealmUserId()

	g.Expect(i.IsEqual(i)).To(BeTrue())
}

func TestRealmUserId_IsEqual_ShouldReturnFalseWithNotEqualRealmUserId(t *testing.T) {
	g := NewWithT(t)
	i := NewRealmUserId()
	i2 := NewRealmUserId()

	g.Expect(i.IsEqual(i2)).To(BeFalse())
}
