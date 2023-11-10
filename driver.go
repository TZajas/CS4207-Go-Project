package main

import (
	bubbleSort "algorithms/bubbleSort"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func main() {
	const numRuns = 3
	rand.Seed(420) // adds determinism to Slice generation
	fmt.Println("Sorting 300 million numbers...")
	times := make([]time.Duration, numRuns)
	for i := 0; i < numRuns; i++ {
		slice := generateSlice(200) //takes 5-10 seconds to sort for optimised algorithms
		startTime := time.Now()

		//Insert Algorithm Here
		sortedSlice := bubbleSort.BubbleSortDriver(slice)

		time := time.Since(startTime)

		if sort.SliceIsSorted(sortedSlice, func(i, j int) bool { return sortedSlice[i] <= sortedSlice[j] }) {
			fmt.Println("Sorted, Algorithm functional")
			times[i] = time
			fmt.Println("Run", i+1, "time:", times[i])
		} else {
			fmt.Println("Not sorted, Algorithm not functional")
			os.Exit(1) // terminate test if algo fails
		}
	}

	totalTime := time.Duration(0)
	for _, v := range times {
		totalTime += v
	}
	fmt.Println("Average time:", totalTime/numRuns)
}

// Generates a slice of size, size filled with random positive 64bit numbers
func generateSlice(size int) []uint64 {
	slice := make([]uint64, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Uint64()
	}
	return slice
}
