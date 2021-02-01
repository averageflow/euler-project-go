package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var result int

	for i := 0; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			result += i
		}
	}

	fmt.Printf(
		"\nExecution time: %v\nResult: %d\n\n",
		time.Since(start),
		result,
	)
}
