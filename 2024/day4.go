package main

import (
	"fmt"
	"strings"
)

var (
	directions = [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
		{-1, -1},
	}

	diagonalDirections = [][]int{
		{1, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}
)

const (
	XMAS = "XMAS"
	X    = "X"
	M    = "M"
	A    = "A"
	S    = "S"
)

func Day4() {
	GetInputFile(4)

	grid := [][]string{}

	PasreInputFile(4, func(line string) {
		grid = append(grid, strings.Split(line, ""))
	})
	count := 0
	for row := range len(grid) {
		for col := range len(grid[0]) {
			if grid[row][col] == A {
				count += ExploreDirections(grid, row, col)
			}
		}
	}

	fmt.Printf("\ncount: %v\n", count)
}

// Part 2
func ExploreDirections(grid [][]string, row, col int) (result int) {
	buffer := []string{}
	for _, direction := range diagonalDirections {
		dx, dy := direction[0], direction[1]
		nr, nc := row+dx, col+dy

		if IsInbound(grid, nr, nc) {
			buffer = append(buffer, grid[nr][nc])
		}
	}

	if len(buffer) != 4 {
		return
	}
	
	count := 0
	for i := 0; i < len(buffer); i += 2 {
		// Check each half of the buffer for validity
		if (buffer[i] != buffer[i+1]) && (buffer[i] != X) && (buffer[i+1] != X) && (buffer[i] != A) && (buffer[i+1] != A) {
			count++
		}
	}

	// If both halves are valid, then valid
	if count == 2 {
		result++
	}

	return result
}

// func ExploreDirections(grid [][]string, row, col int) (result int){
// 	for _, direction := range directions {
// 		r,c := direction[0], direction[1]
// 		tr, tc := row, col
// 		xmas_ptr := 1
// 		for i := 0; i < len(XMAS) - 1; i++ {
// 			nr, nc := r + tr, c + tc

// 			if !IsInbound(grid, nr, nc) ||  grid[nr][nc] != string(XMAS[xmas_ptr]){
// 				break
// 			}

// 			xmas_ptr++
// 			tr, tc = nr, nc
// 		}

// 		if xmas_ptr >= len(XMAS) {
// 			result++
// 		}

// 	}

// 	return result
// }

func IsInbound(grid [][]string, row, col int) bool {
	return 0 <= row && row < len(grid) && 0 <= col && col < len(grid[0])
}
