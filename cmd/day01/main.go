package main

import (
	"fmt"
	"os"
	"strconv"

	"ovsiienko.xyz/advent-of-code/internal/util"
)

const MAX_ITEMS = 100

func main() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Printf("Failed to read the file: %s\n", err)
		os.Exit(1)
	}
	currentValue := 50
	code := 0
	for i, v := range input {
		if v == "" {
			continue
		}
		direction := string(v[0])
		valueStr := v[1:]
		value, _ := strconv.Atoi(valueStr)
		iterations := value / MAX_ITEMS
		value = value % MAX_ITEMS
		switch direction {
		case "L":
			if currentValue == 0 {
				currentValue = 100
			}
			currentValue -= value
		case "R":
			currentValue += value
		}
		if currentValue < 0 {
			iterations += (currentValue / MAX_ITEMS) + 1
			currentValue += MAX_ITEMS
		} else if currentValue > MAX_ITEMS {
			iterations += currentValue / MAX_ITEMS
			currentValue -= MAX_ITEMS
		} else if currentValue == MAX_ITEMS {
			currentValue = 0
			iterations += 1
		} else if currentValue == 0 {
			iterations += 1
		}
		code += iterations
		fmt.Printf("Index: %d, Input was: %s, currentValue: %d, code: %d, value: %d, iterations: %d \n", i, v, currentValue, code, value, iterations)

	}
	fmt.Println(code)
}
