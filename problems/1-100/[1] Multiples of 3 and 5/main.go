package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var multiplesOf3 []int
	var multiplesOf5 []int

	for i := 1; i < 1000; i++ {
		item := i * 3
		if item >= 1000 {
			break
		}

		multiplesOf3 = append(multiplesOf3, item)
	}

	for i := 1; i < 1000; i++ {
		item := i * 5
		if item >= 1000 {
			break
		}

		var existsIn3 bool
		for j := range multiplesOf3 {
			if item == multiplesOf3[j] {
				existsIn3 = true
			}
		}

		if !existsIn3 {
			multiplesOf5 = append(multiplesOf5, item)
		}
	}

	var sumOf3 int

	for i := range multiplesOf3 {
		sumOf3 += multiplesOf3[i]
	}

	var sumOf5 int

	for i := range multiplesOf5 {
		sumOf5 += multiplesOf5[i]
	}

	fmt.Printf(
		"\nExecution time: %v\nResult: %d\n\n",
		time.Since(start),
		sumOf3+sumOf5,
	)
}
