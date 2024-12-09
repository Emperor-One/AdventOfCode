package main

import "fmt"

func Day3() {
	GetInputFile(3)
	count := 0
	PasreInputFile(3, func(line string) {
		count++
	})
	fmt.Printf("count: %v\n", count)
}