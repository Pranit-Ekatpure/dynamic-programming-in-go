// Memoization technique example with Fibonacci Sequence problem

package main

import (
	"log"
	"time"
)

// Function for measuring the execution time
func trackTime(start time.Time, result int, name string) {
	elapsed := time.Since(start)
	log.Printf("--> %s solution | result: %v | took: %s", name, result, elapsed)
}

// Recursion solution
func expensiveFib(n int) int {
	if n < 2 {
		return n
	}

	result := expensiveFib(n-1) + expensiveFib(n-2)

	return result
}

// Using loop
func fibByLoop(n int) int {
	fibBox := []int{0, 1}

	for i := int(0); i < n; i++ {
		v := fibBox[i] + fibBox[i+1]
		fibBox = append(fibBox, v)
	}

	result := int(fibBox[n])
	defer trackTime(time.Now(), result, "Loop")

	return result
}

// Memoization solution
// Using memoization
func refinedExpensiveFib(n int, cache map[int]int) int {
	if n < 2 {
		cache[n] = n // n is either 0 or 1
		return n
	}

	// Check cache before calling the function recursively
	if _, ok := cache[n-1]; !ok {
		// We haven't come across n-1 before
		cache[n-1] = refinedExpensiveFib(n-1, cache)
	}
	// At this point, n-1 is in the cache

	if _, ok := cache[n-2]; !ok {
		// We haven't come across n-2 before
		cache[n-2] = refinedExpensiveFib(n-2, cache)
	}
	// At this point, n-2 is in the cache

	// Returns the summed up cache
	return cache[n-1] + cache[n-2]
}

func memoize(n int) int {
	cache := make(map[int]int)
	bucket := make([]int, n)

	for i := 1; i <= n; i++ {
		bucket[i-1] = refinedExpensiveFib(i, cache)
	}

	result := bucket[n-1]
	defer trackTime(time.Now(), result, "Memoization")
	return result
}

func main() {
	input := 50

	// Recursion
	start := time.Now()
	result := expensiveFib(input)
	defer trackTime(start, result, "Recursion")

	// Loop
	fibByLoop(input)

	// Memoization
	memoize(input)
}
