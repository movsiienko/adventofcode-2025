package util

import (
	"os"
	"strings"
)

// ReadInput reads the input file and returns its contents as a string
func ReadInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadLines reads the input file and returns its contents as a slice of lines
func ReadLines(filename string) ([]string, error) {
	data, err := ReadInput(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(data), "\n")
	return lines, nil
}
