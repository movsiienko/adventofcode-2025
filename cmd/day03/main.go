package main

import (
	"fmt"
	"ovsiienko.xyz/advent-of-code/internal/util"
	"strconv"
)

func main() {
	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	result := 0
	for _, bank := range lines {
		if bank == "" {
			continue
		}
		stack := []int{}
		head := 0
		for i, battery := range bank {
			batteryInt, err := runeToInt(battery)
			if err != nil {
				fmt.Println("Failed to convert battery")
				panic(1)
			}
			if len(stack) == 0 {
				stack = append(stack, batteryInt)
				head = batteryInt
				continue
			}
			remaning := len(bank) - i
			if remaning >= 12 && batteryInt > head {
				stack = []int{batteryInt}
				continue
			}
			latest := stack[len(stack)-1]
			if remaning+len(stack) >= 12 && batteryInt > latest {
				stack[len(stack)-1] = batteryInt
			} else if len(stack) < 12 {
				stack = append(stack, batteryInt)
			}
		}
		if len(stack) != 12 {
			panic("Stack len should be 12.")
		}
		for i, num := range stack {
			result += (len(stack) - i) * 10 * num
		}
		fmt.Printf("Bank: %s. Stack %v\n", bank, stack)
	}
	fmt.Println(result)

}

func runeToInt(r rune) (int, error) {
	return strconv.Atoi(string(r))
}
