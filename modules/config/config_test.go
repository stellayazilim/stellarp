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

	err := loadEnv()
	g.Expect(err).Should(BeNil())
	g.Expect(os.Getenv("POSTGRES_DATABASE")).To(Equal("postgres"))
}

func TestShouldCreateConfigFromLoadedEnv(t *testing.T) {
	g := NewGomegaWithT(t)
	_ = loadEnv()

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
			expected:    "",
			method:      "GetPostgresPort",
		},
		{
			description: "should have postgresHost set",
			expected:    "localhost",
			method:      "GetPostgresHost",
		},
		{
			description: "should have postgresDBName set",
			expected:    "postgres",
			method:      "GetPostgresDatabase",
		},
	}

	for _, test := range tests {
		g.Expect(reflect.ValueOf(c).MethodByName(test.method).Call(nil)[0].String()).To(Equal(test.expected))

	}
}

func TestInitConfigModule(t *testing.T) {
	g := NewGomegaWithT(t)

	configModule := UseConfigModule()

	testModule := fxtest.New(t, configModule, fx.Invoke(func(c *Config) {

		g.Expect(c).NotTo(BeNil())
	}))

	defer testModule.RequireStop()

	testModule.RequireStart()

}
