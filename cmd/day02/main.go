package main

import (
	"fmt"
	"ovsiienko.xyz/advent-of-code/internal/util"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	input := lines[0]
	products := strings.Split(input, ",")
	total := 0
	for _, v := range products {
		rangeIds := strings.Split(v, "-")
		if len(rangeIds) != 2 {
			fmt.Printf("Invalid range: %s", v)
			continue
		}
		lower, err := strconv.Atoi(rangeIds[0])
		if err != nil {
			fmt.Printf("Failed to convert lower %s", rangeIds[0])
			panic(err)
		}
		upper, err := strconv.Atoi(rangeIds[1])
		if err != nil {
			fmt.Printf("Faile to convert upper: %s", rangeIds[1])
			panic(err)
		}
		for idNumber := lower; idNumber <= upper; idNumber++ {
			if !isValidId(idNumber) {
				total += idNumber
			}
		}
	}
	fmt.Println(total)
}

func isValidId(num int) bool {
	digits := strconv.Itoa(num)
	start := rune(digits[0])
	seq := []rune{start}
	for _, currentDigit := range digits[1:] {
		if currentDigit == start {
			// fmt.Printf("Seq: %s \n", string(seq))
			seqLen := len(seq)
			remaining := len(digits) - seqLen
			if remaining%seqLen != 0 {
				seq = append(seq, currentDigit)
				continue
			}
			shouldContinue := false
			for y := 1; y <= remaining/seqLen; y++ {
				seqCandidate := digits[y*seqLen : (y+1)*seqLen]
				// fmt.Printf("Seq candidate: %v \n", seqCandidate)
				if !slices.Equal(seq, []rune(seqCandidate)) {
					seq = append(seq, currentDigit)
					shouldContinue = true
					break
				}
			}
			if shouldContinue {
				continue
			}
			return false
		}
		seq = append(seq, currentDigit)
	}
	return true
}
