package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnvironment_UnmarshalText(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name     string
			input    string
			expected Environment
		}{
			{name: "local (通常)", input: "local", expected: EnvLocal},
			{name: "test (通常)", input: "test", expected: EnvTest},
			{name: "prd (通常)", input: "prd", expected: EnvPrd},
			{name: "大文字・小文字を区別しない", input: "LOCAL", expected: EnvLocal},
			{name: "スペースが混ざっていても変換できる", input: " local ", expected: EnvLocal},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				b := []byte(tt.input)

				var actual Environment

				err := actual.UnmarshalText(b)

				require.NoError(t, err)
				require.Equal(t, tt.expected, actual)
			})
		}
	})

	t.Run("異常系", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name  string
			input string
		}{
			{name: "空の文字列", input: ""},
			{name: "未知の環境", input: "unknown"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				b := []byte(tt.input)

				var actual Environment

				err := actual.UnmarshalText(b)

				require.Error(t, err)
			})
		}
	})
}

func TestEnvironment_Validate(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name  string
			input Environment
		}{
			{name: "EnvLocalが通過する", input: EnvLocal},
			{name: "EnvTestが通過する", input: EnvTest},
			{name: "EnvPrdが通過する", input: EnvPrd},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				err := tt.input.Validate()

				require.NoError(t, err)
			})
		}
	})

	t.Run("異常系", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name  string
			input Environment
		}{
			{name: "空のEnvironmentは失敗する", input: TestEnvEmpty},
			{name: "不明なEnvironmentは失敗する", input: TestEnvUnkown},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				err := tt.input.Validate()

				require.Error(t, err)
			})
		}
	})
}
