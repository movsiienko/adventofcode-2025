package main

import (
	"fmt"
	"strconv"

	"ovsiienko.xyz/advent-of-code/internal/util"
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
		mainBattery := 0
		secondBattery := 0
		for i, battery := range bank {
			batteryInt, err := runeToInt(battery)
			if err != nil {
				fmt.Println("Failed to convert battery")
				panic(1)
			}
			if batteryInt > mainBattery && i != len(bank)-1 {
				mainBattery = batteryInt
				secondBattery = 0
			} else if batteryInt > secondBattery {
				secondBattery = batteryInt
			}
		}
		fmt.Printf("Bank: %s. Main battery: %d. Second battery: %d. \n", bank, mainBattery, secondBattery)
		result += mainBattery*10 + secondBattery
	}
	fmt.Println(result)

}

func runeToInt(r rune) (int, error) {
	return strconv.Atoi(string(r))
}
