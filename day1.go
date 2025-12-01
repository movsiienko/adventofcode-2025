package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_ITEMS = 100

func main() {
	data, err := os.ReadFile("input.txt")
	input := strings.Split(string(data), "\n")
	if err != nil {
		fmt.Errorf("Failed to read the file: %s", err)
		panic(1)
	}
	currentValue := 50
	code := 0
	for _, v := range input {
		if v == "" {
			continue
		}
		direction := string(v[0])
		valueStr := v[1:]
		value, _ := strconv.Atoi(valueStr)
		value = value % 100
		switch direction {
		case "L":
			currentValue -= value
		case "R":
			currentValue += value

		}
		// fmt.Printf("Interminidate value: %d . ", currentValue)
		if currentValue < 0 {
			currentValue += MAX_ITEMS
		} else if currentValue > MAX_ITEMS {
			currentValue -= MAX_ITEMS
		} else if currentValue == MAX_ITEMS {
			currentValue = 0
		}
		// fmt.Printf("Index: %d, Input was: %s, currentValue: %d \n", i, v, currentValue)
		if currentValue == 0 {
			code += 1
		}

	}
	fmt.Println(code)
}
