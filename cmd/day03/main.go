package main

import (
	"fmt"
	"math"
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
				fmt.Printf("Switching head: %d. Current stack: %v \n", batteryInt, stack)
				stack = []int{batteryInt}
				head = batteryInt
				continue
			}
			latest := stack[len(stack)-1]
			// TODO: Simillar to how we check head we need to check other numbers since the begining depending on how much we have left in the bank
			if remaning+len(stack) > 12 && batteryInt > latest {
				fmt.Printf("Switching to latest: %d. Current stack: %v \n", batteryInt, stack)
				stack[len(stack)-1] = batteryInt
			} else if remaning+len(stack) > 12 {

			} else if len(stack) < 12 {
				fmt.Printf("Appending %d. Current stack: %v \n", batteryInt, stack)
				stack = append(stack, batteryInt)
			}
		}
		fmt.Printf("Bank: %s. Stack %v\n", bank, stack)
		if len(stack) != 12 {
			panic("Stack len should be 12.")
		}
		for i, num := range stack {
			result += int(math.Pow10(len(stack)-i-1)) * num
		}
	}
	fmt.Println(result)

}

func runeToInt(r rune) (int, error) {
	return strconv.Atoi(string(r))
}
