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
			mode = checkSafety
		} else {
			mode = checkSafety
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
func checkSafety(list []int) int {
	safety_inc, safety_dec := 0, 0

	for i := 0; i < len(list); i++ {
		safety_inc = increasing(slices.Concat(list[:i], list[i+1:]))
		safety_dec = decreasing(slices.Concat(list[:i], list[i+1:]))

		if safety_inc == 1 || safety_dec == 1 {
			fmt.Println(slices.Concat(list[:i], list[i+1:]))
			return 1
		}
	}
	return 0
}
