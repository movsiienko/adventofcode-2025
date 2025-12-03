package main

import (
	"fmt"
	"testing"
)

func TestIsValidId(t *testing.T) {
	var tests = []struct {
		number int
		want   bool
	}{
		// {222222, false},
		// {565656, false},
		// {1010, false},
		// {11, false},
		// {22, false},
		// {99, false},
		// {111, false},
		// {999, false},
		// {1010, false},
		{1188511885, false},
		// {222222, false},
		{446446, false},
		// {38593859, false},
		// {565656, false},
		// {824824824, false},
		// {2121212121, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.number)
		t.Run(testname, func(t *testing.T) {
			isValid := isValidId(tt.number)
			if isValid != tt.want {
				t.Errorf("got %t, want %t", isValid, tt.want)
			}
		})
	}
}
