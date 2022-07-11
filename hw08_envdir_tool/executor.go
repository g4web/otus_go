package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

var exitError *exec.ExitError

func RunCmd(cmd []string, env Environment) (returnCode int) {
	executor := createExecutor(cmd)
	setEnvToExecutor(executor, env)
	return runExecutor(executor)
}

func createExecutor(cmd []string) *exec.Cmd {
	command := cmd[0]
	args := cmd[1:]
	cmdExec := exec.Command(command, args...)
	cmdExec.Env = os.Environ()
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr

	return cmdExec
}

func setEnvToExecutor(executor *exec.Cmd, env Environment) {
	for envVar, val := range env {
		executor.Env = append(executor.Env, envVar+"="+val)
	}
}

func runExecutor(executor *exec.Cmd) int {
	exitCode := 0

	if err := executor.Start(); err != nil {
		log.Fatal(err)

		return exitCode
	}

	if err := executor.Wait(); err != nil {
		if errors.As(err, &exitError) {
			exitCode = exitError.ExitCode()
		} else {
			log.Fatal(err)
		}
	}

	return exitCode
}
