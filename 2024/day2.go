package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day2() {
	GetInputFile(2)
	safeLists := 0
	PasreInputFile(2, func(line string) {
		lineSlice := strings.Split(line, " ")

		var inputList []int
		for _, str := range lineSlice {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("error while converting to int", err)
			}
			inputList = append(inputList, num)
		}

		var mode func([]int) int
		if inputList[0] > inputList[1] {
			mode = decreasingPart2
		} else {
			mode = increasingPart2
		}
		safeLists += mode(inputList)
	})

	fmt.Printf("safeLists: %v\n", safeLists)
}

func decreasing(list []int) int {
	for i := 1; i < len(list); i++ {
		if list[i] >= list[i-1] {
			return 0
		} else if (list[i-1] - list[i]) > 3 {
			return 0
		}
	}
	return 1
}

func increasing(list []int) int {
	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			return 0
		} else if (list[i] - list[i-1]) > 3 {
			return 0
		}
	}
	return 1
}

// part 2
const (
	INCREASING = "increasing"
	DECREASING = "decreasing"
)

func checkSafety(list []int, order string) int {
	for i := 0; i < len(list)-1; i++ {
		list = slices.Concat(list[:i], list[i+2:]) 

		fmt.Println(list)
	}
	return 1
}

func decreasingPart2(list []int) int {

	for i := 1; i < len(list); i++ {
		if list[i] >= list[i-1] {
			return checkSafety(list, DECREASING)
		} else if (list[i-1] - list[i]) > 3 {
			return checkSafety(list, DECREASING)
		}
	}
	return 1
}

func increasingPart2(list []int) int {

	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			return checkSafety(list, INCREASING)
		} else if (list[i] - list[i-1]) > 3 {
			return checkSafety(list, INCREASING)
		}
	}
	return 1
}
