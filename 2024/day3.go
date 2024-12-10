package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

func Day3() {
	GetInputFile(3)
	operations := make([]string, 0)
	mul_pattern := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)`)

	PasreInputFile(3, func(line string) {
		operations = slices.Concat(operations, mul_pattern.FindAllString(line, -1))
	})

	sum := 0
	number_pattern := regexp.MustCompile(`\d{1,3}`)
	fmt.Printf("operations: %v\n", operations)
	for _, operation := range operations {
		operands := number_pattern.FindAllString(operation, -1)
		fmt.Printf("operands: %v\n", operands)
		left, _ := strconv.Atoi(operands[0])
		right, _ := strconv.Atoi(operands[1])

		sum += (left * right)
	}
	fmt.Printf("sum: %v\n", sum)
}

// part 2

