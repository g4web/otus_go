package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	environments := make(Environment)

	for _, f := range files {
		filePath := path.Join(dir, f.Name())
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}

		firstLine := getFirstLineFromFile(file)
		environments[f.Name()] = clearLine(firstLine)

		file.Close()
	}

	return environments, nil
}

func getFirstLineFromFile(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var firstLine string
	if scanner.Scan() {
		return scanner.Text()
	}

	return firstLine
}

func clearLine(firstLine string) string {
	firstLine = strings.TrimRight(firstLine, " \t\n")
	firstLine = string(bytes.ReplaceAll([]byte(firstLine), []byte("\x00"), []byte("\n")))

	return firstLine
}
