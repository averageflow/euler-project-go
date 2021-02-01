package main

import (
	"fmt"
	"time"
)

func IsPrime(n int64) bool {
	for i := int64(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()

	wantedNumber := int64(600851475143)

	max := int64(0)
	for i := int64(2); i <= wantedNumber; i++ {
		if wantedNumber%i == 0 {
			if IsPrime(i) {
				max = i
			}
			wantedNumber /= i
			i = 2
		}
	}

	fmt.Printf(
		"\nExecution time: %v\nResult: %d\n\n",
		time.Since(start),
		max,
	)
}
