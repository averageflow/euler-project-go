package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var resultingItems []int64

	resultingItems = append(resultingItems, 0)
	resultingItems = append(resultingItems, 1)

	i := 2

	for {
		sum := resultingItems[i-1] + resultingItems[i-2]

		if sum < 4000000 {
			resultingItems = append(resultingItems, sum)
			i++
		} else {
			break
		}
	}

	var result int64

	for j := range resultingItems {
		if resultingItems[j]%2 == 0 {
			result += resultingItems[j]
		}
	}

	fmt.Printf(
		"\nExecution time: %v\nResult: %d\n\n",
		time.Since(start),
		result,
	)
}
