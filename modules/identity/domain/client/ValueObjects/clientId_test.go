package ValueObjects

import (
	"errors"
	. "github.com/onsi/gomega"
	"testing"
)

func TestClientId_ShouldCreateNewId(t *testing.T) {
	g := NewGomegaWithT(t)

	c := NewClientId()

	g.Expect(c).NotTo(BeNil())

}

func TestClientId_ShouldHaveValue(t *testing.T) {
	g := NewGomegaWithT(t)

	c := NewClientId()
	g.Expect(c.GetValue()).To(Equal(c.value))
}

func TestClientId_MustBeTypeOfClientId(t *testing.T) {
	g := NewGomegaWithT(t)
	c := NewClientId()
	g.Expect(c).To(BeAssignableToTypeOf(ClientId{}))
}

func TestClientId_MustBeParsedFromValidString(t *testing.T) {
	g := NewGomegaWithT(t)
	c := NewClientId()

	p, e := ParseClientId(c.GetValue())

	g.Expect(e).To(BeNil())
	g.Expect(p).To(Equal(c))
}

func TestClientId_MustBeReturnErrorParsedFromInvalidString(t *testing.T) {
	g := NewGomegaWithT(t)
	s := "incorrectid"
	p, e := ParseClientId(s)
	g.Expect(p.value).To(BeEmpty())
	g.Expect(e).NotTo(BeNil())
	g.Expect(e).Should(MatchError(errors.New("Incorrect format")))
}

func TestClientId_IsEqual_ShouldReturnTrueIfIdsAreEqual(t *testing.T) {
	g := NewGomegaWithT(t)
	s := NewClientId()
	s2, _ := ParseClientId(s.value)
	r := s.IsEqual(s2)
	g.Expect(r).To(BeTrue())
}

func TestClientId_IsEqual_ShouldReturnFalseIfIdsAreNotEqual(t *testing.T) {
	g := NewGomegaWithT(t)
	s := NewClientId()
	s2 := NewClientId()

	r := s.IsEqual(s2)

	g.Expect(r).To(BeFalse())
}
