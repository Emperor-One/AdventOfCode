package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day1() (TotalDistance, SimilarityScore int){
	GetInputFile(1)

	var (
		List1         []int
		List2         []int
	)

	ParseInputFile(1, func(line string) {
		lineSlice := strings.Split(line, "   ")
		left, _ := strconv.Atoi(lineSlice[0])
		right, _ := strconv.Atoi(lineSlice[1])

		List1 = append(List1, left)
		List2 = append(List2, right)
	})

	slices.Sort(List1)
	slices.Sort(List2)

	for i, num := range List1 {
		TotalDistance += int(math.Abs(float64(List2[i] - num)))
	}

	// Part 2
	for _, num1 := range List1 {
		numCount := 0
		for _, num2 := range List2 {
			if num1 == num2 {
				numCount++
			}
		}
		SimilarityScore += num1 * numCount
	}


	return
}
