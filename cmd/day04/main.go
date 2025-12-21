package main

import (
	"fmt"
	"ovsiienko.xyz/advent-of-code/internal/util"
	"strings"
)

type postions = struct {
	startPosition int
	endPosition   int
}

func calcualtePositions(index int, limit int) *postions {
	limit--
	var startPosition, endPosition int
	if index == 0 {
		startPosition = 0
	} else {
		startPosition = index - 1
	}
	if index == limit {
		endPosition = index
	} else {
		endPosition = index + 1
	}
	return &postions{startPosition: startPosition, endPosition: endPosition}
}

func main() {
	lines, err := util.ReadLines("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	iterationOutput := make([]string, len(lines))
	copy(iterationOutput, lines)
	result := 0
	for lineNumber, line := range lines {
		if line == "" {
			continue
		}
		for columnNumber, value := range line {
			if string(value) == "." {
				continue
			}
			adjPosCount := 0
			xPositions := calcualtePositions(columnNumber, len(line))
			yPositions := calcualtePositions(lineNumber, len(lines))
			// if lineNumber == 0 {
			// 	fmt.Printf("For column %d positions are: \n Start positions: %v \n End positions: %v \n", columnNumber, xPositions, yPositions)
			// }

			for x := xPositions.startPosition; x <= xPositions.endPosition; x++ {
				for y := yPositions.startPosition; y <= yPositions.endPosition; y++ {
					if x == columnNumber && y == lineNumber {
						continue
					}
					adjValue := lines[y][x]
					// if lineNumber == 0 {
					// 	fmt.Printf("  For x pos: %d and y pos: %d found value is %s \n", x, y, string(adjValue))
					// }
					if adjValue == '@' {
						adjPosCount++
					}
				}
			}

			if adjPosCount < 4 {
				result++
				lineRune := []rune(iterationOutput[lineNumber])
				lineRune[columnNumber] = 'x'
				iterationOutput[lineNumber] = string(lineRune)
			}

		}
	}
	fmt.Println(result)
	fmt.Println(strings.Join(iterationOutput, "\n"))
}
