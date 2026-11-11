package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_greet(t *testing.T) {
	t.Parallel()

	actual := greet("Gopher")

	require.Equal(t, "Hello Gopher!\n", actual)
}
