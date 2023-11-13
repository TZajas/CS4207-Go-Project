package main

import (
	insertionSort "algorithms/insertionSort"
	mergeSort "algorithms/mergeSort"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	numbers := generateNumbers(10000, -10000)

	// Start Timer
	startTime := time.Now()

	// Insert either "insertion" or "merge" into the string parameter to use either
	// the concurrent insertion sort or the concurrent merge sort
	sortedNumbers := concurrentSort("insertion", numbers)

	// End Timer
	endTime := time.Since(startTime)

	// Make sure array is sorted properly
	isSorted := sort.SliceIsSorted(sortedNumbers, func(i, j int) bool {
		return sortedNumbers[i] < sortedNumbers[j]
	})

	if isSorted {
		fmt.Printf("Successfully Sorted Array in %v\n", endTime)
	} else {
		fmt.Println("Failed to sort Array.")
	}
}

// Generate a random array of integers between max and min
func generateNumbers(max, min int) []int {
	// Seed the random number generator
	rand.Seed(420)

	// Generate a random array of 10,000 integers
	randomSlice := make([]int, 100000)
	for i := 0; i < 10000; i++ {
		randomSlice[i] = rand.Intn(max-min+1) + min
	}

	return randomSlice
}

// Decides which algorithm to use
func concurrentSort(alg string, nums []int) []int {
	if alg == "insertion" {
		// Through trial and error I found using 1800 goroutines for this algorithm yields the quickest results.
		// I believe this is because the regular insertion sort works best on smaller arrays
		// But if make the chunks any smaller the algorithm gets slowed down by the mergeSlices function
		// Which has too many slices to merge
		// I found the sweet spot of the number of goroutines to be approx 18% of total amount of numbers in the array
		return insertionSort.ConcurrentInsertionSort(nums, 1800)
	} else if alg == "merge" {
		return mergeSort.ConcurrentMergeSort(nums)
	} else {
		return nums
	}
}
