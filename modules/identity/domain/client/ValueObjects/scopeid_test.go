package ValueObjects

import (
	"errors"
	. "github.com/onsi/gomega"
	"testing"
)

func TestScopeId_ShouldCreateNewId(t *testing.T) {
	g := NewGomegaWithT(t)

	c := NewScopeId()

	g.Expect(c).NotTo(BeNil())

}

func TestScopeId_ShouldHaveValue(t *testing.T) {
	g := NewGomegaWithT(t)

	c := NewScopeId()
	g.Expect(c.GetValue()).To(Equal(c.value))
}

func TestScopeId_MustBeTypeOfScopeId(t *testing.T) {
	g := NewGomegaWithT(t)
	c := NewScopeId()
	g.Expect(c).To(BeAssignableToTypeOf(ScopeId{}))
}

func TestScopeId_MustBeParsedFromValidString(t *testing.T) {
	g := NewGomegaWithT(t)
	c := NewScopeId()

	p, e := ParseScopeId(c.GetValue())

	g.Expect(e).To(BeNil())
	g.Expect(p).To(Equal(c))
}

func TestScopeId_MustBeReturnErrorParsedFromInvalidString(t *testing.T) {
	g := NewGomegaWithT(t)
	s := "incorrectid"
	p, e := ParseScopeId(s)
	g.Expect(p.value).To(BeEmpty())
	g.Expect(e).NotTo(BeNil())
	g.Expect(e).Should(MatchError(errors.New("Incorrect format")))
}

func TestScopeId_IsEqual_ShouldReturnTrueIfIdsAreEqual(t *testing.T) {
	g := NewGomegaWithT(t)
	s := NewScopeId()
	s2, _ := ParseScopeId(s.value)
	r := s.IsEqual(s2)
	g.Expect(r).To(BeTrue())
}

func TestScopeId_IsEqual_ShouldReturnFalseIfIdsAreNotEqual(t *testing.T) {
	g := NewGomegaWithT(t)
	s := NewScopeId()
	s2 := NewScopeId()

	r := s.IsEqual(s2)

	g.Expect(r).To(BeFalse())
}
