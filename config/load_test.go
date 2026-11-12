package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.yaml.in/yaml/v4"
)

func TestLoad(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		t.Setenv("LIXY_APP_ENV", "test")

		config, err := Load()

		require.NoError(t, err)
		require.NotNil(t, config)
		require.Equal(t, EnvTest, config.App.Env)
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("LIXY_APP_ENVの値が不正な場合、エラーが返される", func(t *testing.T) {
			t.Setenv("LIXY_APP_ENV", TestEnvUnkown.String())

			config, err := Load()

			require.Error(t, err)
			require.Nil(t, config)
		})
	})
}

func TestLoadFrom(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		t.Setenv("LIXY_APP_ENV", EnvTest.String())

		config, err := LoadFrom(EnvTest)

		require.NoError(t, err)
		require.NotNil(t, config)
		require.Equal(t, EnvTest, config.App.Env)
	})
}

// Test_EnvFile は設定ファイルが以下の規約に則っているかを検証するテストである。
//   - env/<env>.yaml というファイル名であること
//   - YAMLとしてパース可能であること
//   - env/<env>.yaml の app.env が <env> と一致すること
//
//nolint:testifylint // 検査に失敗したら次のファイルを対象に検査を行うために assert を使っている
func Test_EnvFile(t *testing.T) {
	t.Parallel()

	// 実行順序に依存した結果が生じないようにランダムに読み込んで実行する
	targets := map[Environment]struct{}{
		EnvLocal: {}, EnvTest: {}, EnvPrd: {},
	}

	for env := range targets {
		path := fmt.Sprintf("env/%s.yaml", env.String())

		f, err := envFiles.Open(path)
		if !assert.NoError(t, err, "env/%s.yaml の読み込みに失敗", env.String()) {
			continue
		}
		defer f.Close()

		var data struct {
			App struct {
				Env string `yaml:"env"`
			} `yaml:"app"`
		}

		err = yaml.NewDecoder(f).Decode(&data)
		if !assert.NoError(t, err, "env/%s.yaml のパースに失敗", env.String()) {
			continue
		}

		if !assert.Equal(
			t,
			data.App.Env,
			env.String(),
			"env/%s.yaml の app.env が %s と一致しない",
			env.String(),
			env.String(),
		) {
			continue
		}

		t.Logf("env/%s.yaml の読み込みに成功", env.String())
	}
}
