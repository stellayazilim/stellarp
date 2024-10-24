package config

import (
	"fmt"
	"go.uber.org/fx"
	"os"
	"path/filepath"
)

func UseConfigModule(fileNames ...string) fx.Option {
	cwd, _ := os.Getwd()

	for i := range fileNames {
		fileNames[i] = filepath.Join(cwd, fileNames[i])
	}
	module := fx.Module("ConfigModule", fx.Provide(
		func() (*Config, error) {
			err := loadEnv(fileNames...)

			fmt.Println(err)
			return newConfig(), err
		},
	))

	return module
}
