package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()

		t.Run("正常な Config の場合、バリデーションを通過する", func(t *testing.T) {
			t.Parallel()

			config := Config{
				App: AppConfig{Env: EnvTest},
			}

			err := config.Validate()

			require.NoError(t, err)
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Parallel()

		t.Run("不正な AppConfig の場合、バリデーションに失敗する", func(t *testing.T) {
			t.Parallel()

			config := Config{
				App: AppConfig{Env: TestEnvUnkown},
			}

			err := config.Validate()

			require.Error(t, err)
		})
	})
}

func TestConfig_IsEmpty(t *testing.T) {
	t.Parallel()

	t.Run("全ての必須項目に値が設定されている場合、falseを返す", func(t *testing.T) {
		t.Parallel()

		config := Config{
			App: AppConfig{Env: EnvTest},
		}

		actual := config.IsEmpty()

		require.False(t, actual)
	})

	t.Run("AppConfig の必須項目が全て設定されていない場合、trueを返す", func(t *testing.T) {
		t.Parallel()

		config := Config{
			App: AppConfig{Env: TestEnvEmpty},
		}

		actual := config.IsEmpty()

		require.True(t, actual)
	})
}

func TestAppConfig_Validate(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()

		t.Run("正常な AppConfig の場合、バリデーションを通過する", func(t *testing.T) {
			t.Parallel()

			config := AppConfig{
				Env: EnvTest,
			}

			err := config.Validate()

			require.NoError(t, err)
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Parallel()

		t.Run("Env が空の場合、バリデーションが失敗する", func(t *testing.T) {
			t.Parallel()

			config := AppConfig{
				Env: TestEnvEmpty,
			}

			err := config.Validate()

			require.Error(t, err)
		})

		t.Run("Env が存在しない値だった場合、バリデーションが失敗する", func(t *testing.T) {
			t.Parallel()

			config := AppConfig{
				Env: TestEnvUnkown,
			}

			err := config.Validate()

			require.Error(t, err)
		})
	})
}

func TestAppConfig_IsEmpty(t *testing.T) {
	t.Parallel()

	t.Run("全ての必須項目に値が設定されている場合、falseを返す", func(t *testing.T) {
		t.Parallel()

		config := AppConfig{
			Env: EnvTest,
		}

		actual := config.IsEmpty()

		require.False(t, actual)
	})

	t.Run("Env が空の場合、trueを返す", func(t *testing.T) {
		t.Parallel()

		config := AppConfig{
			Env: TestEnvEmpty,
		}

		actual := config.IsEmpty()

		require.True(t, actual)
	})
}
