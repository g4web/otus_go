package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Sets envs", func(t *testing.T) {
		envs := make(Environment)
		envs["HELLO"] = "greeting"
		envs["BAR"] = "awesome"

		cmd := []string{"/bin/bash", "testdata/echo.sh"}

		r, w, _ := os.Pipe()
		os.Stdout = w
		_ = RunCmd(cmd, envs)
		w.Close()
		out, _ := ioutil.ReadAll(r)
		require.Equal(t, "HELLO is (greeting)\nBAR is (awesome)", string(out[0:36]))
	})

	t.Run("non-existing command", func(t *testing.T) {
		cmd := []string{"/bin/bash", "non-existing.sh"}
		envs := make(Environment)
		ret := RunCmd(cmd, envs)
		require.Equal(t, 127, ret)
	})
}
