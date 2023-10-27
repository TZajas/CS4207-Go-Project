package mergeSort

import (
	"sync"
)

// ConcurrentMergeSort performs a concurrent sort on an input slice of uint64.
func ConcurrentMergeSort(nums []uint64) []uint64 {
	// If the lenght of input is less than or equal to 1 return input
	if len(nums) <= 1 {
		return nums
	}

	// If the length of the input is small (<= 2000), use the sequential merge sort.
	if len(nums) <= 2000 {
		return sequentialMergeSort(nums)
	} else {
		// Calculate the middle index.
		mid := len(nums) / 2

		// Initialize slices for the left and right sub-arrays.
		var left, right []uint64

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
}

func sequentialMergeSort(nums []uint64) []uint64 {
	if len(nums) <= 1 {
		return nums
	}

	// Recursively sort the left and right sub-arrays.
	mid := len(nums) / 2
	left := sequentialMergeSort(nums[:mid])
	right := sequentialMergeSort(nums[mid:])

	return mergeSlices(left, right)
}

// Merges two sorted slices into a single sorted slice.
func mergeSlices(left, right []uint64) []uint64 {
	output := make([]uint64, len(left)+len(right))
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
