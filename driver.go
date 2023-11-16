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
	fmt.Println("Sorting Array of 10,000 Random Integers...")
	fmt.Println("")

	numbers := generateNumbers(10000, -10000)

	// Start Timer
	startTime := time.Now()

	// Insert either "insertion" or "merge" into the string parameter to use either
	// the concurrent insertion sort or the concurrent merge sort
	sortedNumbersInsertion := concurrentSort("insertion", numbers)

	// End Timer
	endTimeInsertion := time.Since(startTime)

	// Make sure array is sorted properly
	isSortedInsertion := sort.SliceIsSorted(sortedNumbersInsertion, func(i, j int) bool {
		return sortedNumbersInsertion[i] < sortedNumbersInsertion[j]
	})

	// Printing completion time of insertion sort
	if isSortedInsertion {
		fmt.Printf("Successfully Sorted Array Using a Concurrent Insertion Sort in %v\n", endTimeInsertion)
		fmt.Println("")
	} else {
		fmt.Println("Failed to Sort Array.")
	}

	// Start Timer again
	startTime = time.Now()

	// Insert either "insertion" or "merge" into the string parameter to use either
	// the concurrent insertion sort or the concurrent merge sort
	sortedNumbersMerge := concurrentSort("merge", numbers)

	// End Timer again
	endTimeMerge := time.Since(startTime)

	isSortedMerge := sort.SliceIsSorted(sortedNumbersMerge, func(i, j int) bool {
		return sortedNumbersMerge[i] < sortedNumbersMerge[j]
	})

	// Printing completion time of merge sort
	if isSortedMerge {
		fmt.Printf("Successfully Sorted Array Using a Concurrent Merge Sort in %v\n", endTimeMerge)
		fmt.Println("")
	} else {
		fmt.Println("Failed to Sort Array.")
	}

	// Printing the difference between the two
	if endTimeInsertion < endTimeMerge {
		fmt.Printf("The Insertion Sort was %v faster than the Merge Sort", endTimeMerge-endTimeInsertion)
	} else {
		fmt.Printf("The Merge Sort was %v faster than the Insertion Sort", endTimeInsertion-endTimeMerge)
	}
}

// Generate a random array of integers between max and min
func generateNumbers(max, min int) []int {
	// Seed the random number generator
	rand.Seed(420)

	// Generate a random array of 10,000 integers
	randomSlice := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		randomSlice[i] = rand.Intn(max-min+1) + min
	}

	return randomSlice
}

// Decides which algorithm to use
func concurrentSort(alg string, nums []int) []int {
	if alg == "insertion" {
		// Through trial and error I found using 150 goroutines for this algorithm on 10,000 integers yields the quickest results.
		// I believe this is because the regular insertion sort works best on smaller arrays
		// But if I make the chunks any smaller the algorithm gets slowed down by the mergeSlices function
		// Which has too many slices to merge
		// Also creating and managing goroutines has associated overhead.
		// If you create too many goroutines the overhead might dominate the actual computation leading to slower performance.
		return insertionSort.ConcurrentInsertionSort(nums, 150)
	} else if alg == "merge" {
		return mergeSort.ConcurrentMergeSort(nums)
	} else {
		return nums
	}
}
