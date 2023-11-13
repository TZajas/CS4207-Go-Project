package mergeSort

import (
	"sync"
)

// ConcurrentMergeSort performs a concurrent sort on an input slice of ints.
func ConcurrentMergeSort(nums []int) []int {
	// If the lenght of input is less than or equal to 1 return input
	if len(nums) <= 1 {
		return nums
	}

	// Calculate the middle index.
	mid := len(nums) / 2

	// Initialize slices for the left and right sub-arrays.
	var left, right []int

	// Create a WaitGroup to synchronize the goroutines.
	var wg sync.WaitGroup
	wg.Add(2)

	// Launch a goroutine to sort the left sub-array.
	go func() {
		left = ConcurrentMergeSort(nums[:mid])
		wg.Done()
	}()

	// Launch a goroutine to sort the right sub-array.
	go func() {
		right = ConcurrentMergeSort(nums[mid:])
		wg.Done()
	}()

	// Wait for both goroutines to complete.
	wg.Wait()

	// Merge the sorted left and right sub-arrays and return the result.
	return mergeSlices(left, right)

}

// Merges two sorted slices into a single sorted slice.
func mergeSlices(left, right []int) []int {
	output := make([]int, len(left)+len(right))
	i, j := 0, 0

	// Merge elements from the left and right sub-arrays while maintaining the sorted order.
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			output[i+j] = left[i]
			i++
		} else {
			output[i+j] = right[j]
			j++
		}
	}

	// Append any remaining elements from the left sub-array. (if any left)
	for i < len(left) {
		output[i+j] = left[i]
		i++
	}

	// Append any remaining elements from the right sub-array. (if any left)
	for j < len(right) {
		output[i+j] = right[j]
		j++
	}

	// Return the merged and sorted slice.
	return output
}
