package main

import (
	"fmt"
	"ovsiienko.xyz/advent-of-code/internal/util"
)

const (
	ActiveCell   = '@'
	InactiveCell = 'x'
	MinNeigbors  = 4
)

type positions = struct {
	startPosition int
	endPosition   int
}

func calculatePositions(index int, limit int) *positions {
	return &positions{startPosition: max(0, index-1), endPosition: min(limit-1, index+1)}
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
	for {
		iterationResult := 0
		for lineNumber, line := range lines {
			if line == "" {
				continue
			}
			for columnNumber, value := range line {
				if value != ActiveCell {
					continue
				}
				adjPosCount := 0
				xPositions := calculatePositions(columnNumber, len(line))
				yPositions := calculatePositions(lineNumber, len(lines))

				for x := xPositions.startPosition; x <= xPositions.endPosition; x++ {
					for y := yPositions.startPosition; y <= yPositions.endPosition; y++ {
						if x == columnNumber && y == lineNumber {
							continue
						}
						adjValue := lines[y][x]
						if adjValue == ActiveCell {
							adjPosCount++
						}
					}
				}

				if adjPosCount < MinNeigbors {
					iterationResult++
					lineRune := []rune(iterationOutput[lineNumber])
					lineRune[columnNumber] = InactiveCell
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
