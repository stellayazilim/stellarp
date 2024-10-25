package config

import (
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"os"
	"reflect"
	"testing"
)

func TestLoadEnvShouldLoadEnvFile(t *testing.T) {
	g := NewGomegaWithT(t)

	err := loadEnv("../../.test.env")
	g.Expect(err).Should(BeNil())
	g.Expect(os.Getenv("POSTGRES_DATABASE")).To(Equal("postgres"))
}

func TestShouldCreateConfigFromLoadedEnv(t *testing.T) {
	g := NewGomegaWithT(t)
	_ = loadEnv("../../.test.env")

	c := newConfig()

	tests := []struct {
		description string
		expected    string
		method      string
	}{
		{

			description: "should have postgresPassword set",
			expected:    "postgres",
			method:      "GetPostgresPassword",
		},
		{
			description: "should have postgresUsername set",
			expected:    "postgres",
			method:      "GetPostgresUser",
		},
		{
			description: "should have postgresPort set",
			expected:    "5432",
			method:      "GetPostgresPort",
		},
		{
			description: "should have postgresHost set",
			expected:    os.Getenv("POSTGRES_HOST"),
			method:      "GetPostgresHost",
		},
		{
			description: "should have postgresDBName set",
			expected:    "postgres",
			method:      "GetPostgresDatabase",
		},
		{
			description: "should have identityModulePort set",
			expected:    "8091",
			method:      "GetIdentityModulePort",
		},
	}

	for _, test := range tests {
		g.Expect(reflect.ValueOf(c).MethodByName(test.method).Call(nil)[0].String()).To(Equal(test.expected))

	}
}

func TestInitConfigModule(t *testing.T) {
	g := NewGomegaWithT(t)

	configModule := UseConfigModule("../../.test.env")

	testModule := fxtest.New(t, configModule, fx.Invoke(func(c *Config) {

		g.Expect(c).NotTo(BeNil())
	}))

	defer testModule.RequireStop()

	testModule.RequireStart()

}
