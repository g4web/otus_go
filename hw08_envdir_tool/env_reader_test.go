package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("test env", func(t *testing.T) {
		envs, _ := ReadDir("testdata/env/")
		require.Equal(t, "bar", envs["BAR"])
		require.Equal(t, ``, envs["EMPTY"])
		require.Equal(t, `   foo
with new line`, envs["FOO"])
		require.Equal(t, `"hello"`, envs["HELLO"])
		require.Equal(t, "", envs["UNSET"])
	})

	t.Run("non-existent directory", func(t *testing.T) {
		_, err := ReadDir("non_existent_directory/")
		require.NotEqual(t, nil, err)
	})

	t.Run("file", func(t *testing.T) {
		_, err := ReadDir("testdata/echo.sh")
		require.NotEqual(t, nil, err)
	})
}
