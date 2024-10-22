package contracts

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	. "github.com/onsi/gomega"
	"testing"
)

func TestREalm_CreateRealmRequest_MustUnMarshalJson(t *testing.T) {

	g := NewWithT(t)
	jsonS := `{
    		"name": ""	
	}`

	jData := &RealmCreateRequest{}
	err := json.Unmarshal([]byte(jsonS), jData)

	g.Expect(err).To(BeNil())

	g.Expect(jData.Name).To(Equal("jdoe"))
}

func TestRealm_CreateRealmRequest_MustHaveNameFieldPresented(t *testing.T) {
	g := NewWithT(t)
	validate := validator.New(validator.WithRequiredStructEnabled())
	jsonS := `{
    		"name": ""	
	}`

	jData := &RealmCreateRequest{}
	if err := json.Unmarshal([]byte(jsonS), jData); err != nil {
		Expect(err).To(BeNil())
	}

	err := validate.Struct(jData)

	g.Expect(err).NotTo(BeNil())

	g.Expect(jData.Name).To(Equal(""))
}
