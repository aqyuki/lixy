package config

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/cockroachdb/errors"
	koanfYAML "github.com/knadh/koanf/parsers/yaml"
	koanfEnv "github.com/knadh/koanf/providers/env/v2"
	koanfFS "github.com/knadh/koanf/providers/fs"
	"github.com/knadh/koanf/v2"
)

const LixyEnvPrefix = "LIXY_"

var ErrUndefinedAppEnv = errors.New("LIXY_APP_ENV is not defined")

//go:embed env/*.yaml
var envFiles embed.FS

func Load() (*Config, error) {
	v, ok := os.LookupEnv("LIXY_APP_ENV")
	if !ok {
		return nil, ErrUndefinedAppEnv
	}

	var env Environment
	if err := env.UnmarshalText([]byte(v)); err != nil {
		return nil, errors.Wrap(err, "failed to spacify APP_ENV")
	}

	return LoadFrom(env)
}

func LoadFrom(env Environment) (*Config, error) {
	k := koanf.New(".")

	fsProvider := koanfFS.Provider(envFiles, fmt.Sprintf("env/%s.yaml", env.String()))
	envProvider := koanfEnv.Provider(".", koanfEnv.Opt{
		Prefix: LixyEnvPrefix,
		TransformFunc: func(key, value string) (string, any) {
			key = strings.TrimPrefix(key, LixyEnvPrefix)
			key = strings.ToLower(key)
			key = strings.ReplaceAll(key, "_", ".")

			return key, value
		},
	})

	parser := koanfYAML.Parser()

	if err := k.Load(fsProvider, parser); err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}

	if err := k.Load(envProvider, nil); err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}

	var config Config
	if err := k.Unmarshal("", &config); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}

	return &config, nil
}
