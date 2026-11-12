package config

import (
	"encoding"
	"strings"

	"github.com/cockroachdb/errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	EnvLocal Environment = "local"
	EnvTest  Environment = "test"
	EnvPrd   Environment = "prd"
)

var ErrUnknownEnv = errors.New("specified unknown env")

var _ encoding.TextUnmarshaler = (*Environment)(nil)

type Environment string

func (e Environment) String() string {
	return strings.ToLower(strings.TrimSpace(string(e)))
}

func (e *Environment) UnmarshalText(b []byte) error {
	env := Environment(strings.ToLower(strings.TrimSpace(string(b))))

	if err := env.Validate(); err != nil {
		return ErrUnknownEnv
	}

	*e = env

	return nil
}

func (e Environment) Validate() error {
	return validation.Validate(
		e.String(),
		validation.Required,
		validation.In(
			EnvLocal.String(),
			EnvTest.String(),
			EnvPrd.String(),
		),
	)
}

func (e Environment) IsEmpty() bool {
	return e.String() == ""
}
