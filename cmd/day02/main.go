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
			id := strconv.Itoa(idNumber)
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

