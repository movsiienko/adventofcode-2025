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
			if remaning+len(stack) > 12 {
				replaced := false
				for stackIndex, value := range stack {
					if ((stackIndex + remaning + 1) > 12) && (batteryInt > value) {
						fmt.Printf("Replaced index %d to value %d. New stack: %v \n", stackIndex, batteryInt, stack)
						stack = append(stack[:stackIndex], batteryInt)
						replaced = true
						break
					}
				}
				if !replaced && len(stack) < 12 {
					stack = append(stack, batteryInt)
				}
			} else if len(stack) < 12 {
				fmt.Printf("Appending %d. Current stack: %v \n", batteryInt, stack)
				stack = append(stack, batteryInt)
			}
		}
		fmt.Printf("Bank: %s. Stack %v\n", bank, stack)
		if len(stack) != 12 {
			errStr := fmt.Sprintf("Stack len should be 12. Current: %d. Stack: %v", len(stack), stack)
			panic(errStr)
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
