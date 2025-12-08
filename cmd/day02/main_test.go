package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestIsValidIdRanges(t *testing.T) {
	count := 0
	var tests = []struct {
		startRange int
		endRange   int
		invalid    []int
	}{
		{11, 22, []int{11, 22}},
		{95, 115, []int{99, 111}},
		{998, 1012, []int{999, 1010}},
		{1188511880, 1188511890, []int{1188511885}},
		{222220, 222224, []int{222222}},
		{1698522, 1698528, []int{}},
		{446443, 446449, []int{446446}},
		{38593856, 38593862, []int{38593859}},
		{565653, 565659, []int{565656}},
		{824824821, 824824827, []int{824824824}},
		{2121212118, 2121212124, []int{2121212121}},
	}
	for _, tt := range tests {
		for i := tt.startRange; i <= tt.endRange; i++ {
			testName := fmt.Sprintf("%d", i)
			t.Run(testName, func(t *testing.T) {
				isValid := isValidId(i)
				count++
				expected := !slices.Contains(tt.invalid, i)
				if isValid != expected {
					t.Errorf("Got %t, want %t", isValid, expected)
				}
			})
		}
	}
}
