package config

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Config struct {
	App AppConfig `koanf:"app"`
}

func (c *Config) Normalize() {
	// TODO: 正規化処理が必要になったらここに書いていく
}

func (c Config) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.App, validation.Required),
	)
}

func (c Config) IsEmpty() bool {
	return c.App.IsEmpty()
}

type AppConfig struct {
	Env Environment `koanf:"env"`
}

func (c AppConfig) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Env, validation.Required),
	)
}

//nolint:staticcheck // S1008: 今後の修正を楽にするために今はあえて冗長に書く
func (c AppConfig) IsEmpty() bool {
	if c.Env != "" {
		return false
	}

	return true
}
