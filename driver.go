package main

import (
	bubbleSort "algorithms/bubbleSort"
	// mergeSort "algorithms/mergeSort"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	numbers := generateNumbers()
	startTime := time.Now()

	sortedNumbers := bubbleSort.ConcurrentBubbleSort(numbers)

	endTime := time.Since(startTime)

	isSorted := sort.SliceIsSorted(sortedNumbers, func(i, j int) bool {
		return sortedNumbers[i] < sortedNumbers[j]
	})

	if isSorted {
		fmt.Printf("Successfully Sorted Array in %v\n", endTime)
	} else {
		fmt.Println("Failed to sort Array.")
	}
}

func generateNumbers() []int {
	// Seed the random number generator
	rand.Seed(420)

	// Specify the range for random numbers
	min := -9999
	max := 9999

	// Generate a random array of 10,000 integers
	randomSlice := make([]int, 100000)
	for i := 0; i < 10000; i++ {
		randomSlice[i] = rand.Intn(max-min+1) + min
	}

	return randomSlice
}
