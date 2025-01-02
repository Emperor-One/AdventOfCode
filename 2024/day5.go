package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day5() {
	GetInputFile(5)

	isUpdateSection := false
	order := make(map[int][]int)
	updates := [][]string{}
	ParseInputFile(5, func(line string) {
		if line == "" {
			isUpdateSection = true
			return
		}

		if !isUpdateSection {
			input := strings.Split(line, "|")

			left, err := strconv.Atoi(input[0])
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
			right, err := strconv.Atoi(input[1])
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}

			order[left] = append(order[left], right)

		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	})

	sum := 0
	for _, update := range updates {
		sum += FindBadUpdate(order, update)
	}
	fmt.Printf("order: %v\n", order)
	fmt.Printf("sum: %v\n", sum)
}

func FindBadUpdate(order map[int][]int, update []string) int {
	isValid := false
	for i := 0; i < len(update) - 1; i++ {
		updateInt, _ := strconv.Atoi(update[i])
		nextInt, _ := strconv.Atoi(update[i+1])
		for _, num := range order[updateInt] {
			if nextInt == num {
				isValid = true
				break
			}
		}
		if !isValid {
			return FixBadUpdate(order, update)
		}
		isValid = false
	}
	return 0
}

func FixBadUpdate(order map[int][]int, update []string) int {
	sort.Slice(update, func(i, j int) bool {
		current, _ := strconv.Atoi(update[i])
		next, _ := strconv.Atoi(update[j])

		for _, num := range order[current] {
			if num == next {
				return true
			}
		}
		return false
	})


	middleIndex := len(update) / 2
	middlePage, _ := strconv.Atoi(update[middleIndex])
	return middlePage
}