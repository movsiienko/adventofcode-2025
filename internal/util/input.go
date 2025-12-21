package util

import (
	"os"
	"strings"
)

func ReadInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadLines(filename string) ([]string, error) {
	data, err := ReadInput(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(data), "\n")
	return lines, nil
}
