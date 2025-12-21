package main

import (
	"fmt"
	"ovsiienko.xyz/advent-of-code/internal/util"
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
	for true {
		iterationResult := 0
		for lineNumber, line := range lines {
			if line == "" {
				continue
			}
			for columnNumber, value := range line {
				if value != '@' {
					continue
				}
				adjPosCount := 0
				xPositions := calcualtePositions(columnNumber, len(line))
				yPositions := calcualtePositions(lineNumber, len(lines))

				for x := xPositions.startPosition; x <= xPositions.endPosition; x++ {
					for y := yPositions.startPosition; y <= yPositions.endPosition; y++ {
						if x == columnNumber && y == lineNumber {
							continue
						}
						adjValue := lines[y][x]
						if adjValue == '@' {
							adjPosCount++
						}
					}
				}

				if adjPosCount < 4 {
					iterationResult++
					lineRune := []rune(iterationOutput[lineNumber])
					lineRune[columnNumber] = 'x'
					iterationOutput[lineNumber] = string(lineRune)
				}
			}
		}
		result += iterationResult
		if iterationResult == 0 {
			break
		}
		lines = iterationOutput
	}
	fmt.Println(result)
}
