package ValueObject

import (
	"github.com/lucsky/cuid"
	. "github.com/onsi/gomega"
	"testing"
)

func TestUserid_NewUserId_ShouldCreateNewUserId(t *testing.T) {
	g := NewWithT(t)
	u := NewUserId()

	g.Expect(u).ToNot(BeNil())

}

func TestUserid_ParseUserId_ShouldCreateUserIdFromValidString(t *testing.T) {
	g := NewWithT(t)
	s := cuid.New()

	u, e := ParseUserId(s)

	g.Expect(e).To(BeNil())
	g.Expect(u).ToNot(BeNil())
}

func TestUserId_ParseUserId_ShouldFailWithInvalidString(t *testing.T) {
	g := NewWithT(t)
	s := "invalid id"

	u, e := ParseUserId(s)

	g.Expect(u).To(BeNil())

	g.Expect(e).ToNot(BeNil())
	g.Expect(e).Should(MatchError("Incorrect format"))
}

func TestUserId_GetValue_ShouldReturnValue(t *testing.T) {
	g := NewWithT(t)
	s := cuid.New()
	u, _ := ParseUserId(s)

	g.Expect(u.GetValue()).NotTo(BeEmpty())
	g.Expect(u.GetValue()).To(Equal(s))
}
