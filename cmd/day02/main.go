package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"ovsiienko.xyz/advent-of-code/internal/util"
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
		v = strings.ReplaceAll(v, "0", "")
		for id := range strings.SplitSeq(v, "-") {
			idNumber, err := strconv.Atoi(id)
			if err != nil {
				fmt.Printf("Failed to cast to number: %s.", id)
				panic(err)
			}
			idLength := int(math.Pow10(len(id) / 2))
			firstHalf := idNumber / idLength
			secondHalf := idNumber % idLength
			if firstHalf == secondHalf {
				total += idNumber
			}
			fmt.Printf("Input Id: %s, First half: %d, Sencod half: %d, Total: %d \n", id, firstHalf, secondHalf, total)
		}
	}
	fmt.Println(total)
}

