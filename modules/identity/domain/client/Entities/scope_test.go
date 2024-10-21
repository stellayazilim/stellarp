package Entities

import (
	"StellaRP/modules/identity/domain/client/ValueObjects"
	"github.com/lucsky/cuid"
	. "github.com/onsi/gomega"
	"testing"
)

func TestScope_ShouldCreateNewScopeWithGeneratedId(t *testing.T) {
	g := NewGomegaWithT(t)

	n := "Application"
	d := "Application scope"
	i := ValueObjects.NewScopeId()
	s := NewScope(i, n, d)

	e := cuid.IsCuid(s.GetId().GetValue())

	g.Expect(s.GetId().GetValue()).NotTo(BeEmpty())
	g.Expect(e).To(BeNil())

}

func TestScope_ShouldCreatedWithRequiredFields(t *testing.T) {
	g := NewGomegaWithT(t)
	i := ValueObjects.NewScopeId()
	n := "Application"
	d := "Application scope"

	s := NewScope(i, n, d)

	g.Expect(s.GetId()).To(BeAssignableToTypeOf(ValueObjects.ScopeId{}))
	g.Expect(s.GetId()).To(Equal(i))
	g.Expect(s.GetId())
	g.Expect(s.GetName()).To(Equal(n))
	g.Expect(s.GetDescription()).To(Equal(d))
}

func TestScope_GetIdShouldReturnScopeId(t *testing.T) {
	g := NewGomegaWithT(t)
	i := ValueObjects.NewScopeId()
	s := NewScope(i, "", "")

	g.Expect(s.GetId()).To(Equal(i))
	g.Expect(s.GetId().GetValue()).To(Equal(i.GetValue()))
	g.Expect(s.GetId().IsEqual(i)).To(BeTrue())
}

func TestScope_ShouldNameCanBeUpdated(t *testing.T) {

	g := NewGomegaWithT(t)
	i := ValueObjects.NewScopeId()
	n := "Application"
	d := "Application scope"

	s := NewScope(i, n, d)

	s.SetName("app")

	g.Expect(s.GetName()).To(Equal("app"))
}

func TestScope_ShouldDescriptionCanBeUpdatedWithNewName(t *testing.T) {
	g := NewGomegaWithT(t)
	i := ValueObjects.NewScopeId()
	n := "Application"
	d := "Application scope"
	s := NewScope(i, n, d)
	s.SetDescription("app description")
	g.Expect(s.GetDescription()).To(Equal("app description"))
}
