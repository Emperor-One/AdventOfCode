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
	pattern := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)|do\(\)|don't\(\)`)

	PasreInputFile(3, func(line string) {
		operations = slices.Concat(operations, pattern.FindAllString(line, -1))
	})

	const (
		Do   = "do()"
		Dont = "don't()"
	)

	NOP := false
	sum := 0
	for _, operation := range operations {
		switch operation{
		case Do:
			NOP = false
			continue
		case Dont:
			NOP = true
			continue
		}

		if !NOP {
			sum += Multiply(operation)
		}
	}

	fmt.Printf("sum: %v\n", sum)
}

func Multiply(operation string) int {
	number_pattern := regexp.MustCompile(`\d{1,3}`)
	operands := number_pattern.FindAllString(operation, -1)

	left, _ := strconv.Atoi(operands[0])
	right, _ := strconv.Atoi(operands[1])

	return left * right
}
