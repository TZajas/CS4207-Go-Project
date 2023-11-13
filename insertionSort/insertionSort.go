package bubbleSort

import (
	"sync"
)

// This concurrent insertion sort algorithm takes the number of goroutines you want to create
// Then splits the array of numbers into the given number of goroutines
// Then concurrently runs regular insertion sort on each of the subarrays
// Lastly it combines the seperated slices using the merge algorithm commonly used in mergeSort
func ConcurrentInsertionSort(nums []int, goroutines int) []int {
	// Create a WaitGroup to synchronize the goroutines.
	var wg sync.WaitGroup

	// Initialise return slice
	var sortedSlices []int

	// Split input slice into chunks(subarrays) depending on the goroutines parameter
	chunks := splitSlice(nums, goroutines)

	// Iterate through the subarrays
	for i := 0; i < len(chunks); i++ {
		wg.Add(1)

		// For each chunk(subarray) create a goroutine and sort the specific chunk
		// using the insertion sort algorithm
		go func(index int) {
			chunks[index] = insertionSort(chunks[index], &wg)
		}(i)
	}

	wg.Wait()

	// Combine and sort all the sorted chunks(subarrays) using the mergeSlices
	// algorithm commonly used in merge sort
	for _, sortedSlice := range chunks {
		sortedSlices = mergeSlices(sortedSlices, sortedSlice)
	}

	return sortedSlices

}

// The insertion sort algorithm works by the assumption that the first element in the array is already sorted.
// Then iterating through the remaining unsorted elements, one at a time.
// For each unsorted element, comparing it with the elements in the sorted portion and insert it into its correct position in the sorted part of the array.
// Repeating this process until the entire array is sorted.
func insertionSort(nums []int, wg *sync.WaitGroup) []int {
	var len = len(nums)
	for i := 1; i < len; i++ {
		j := i
		for j > 0 {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
			j = j - 1
		}
	}

	defer wg.Done()

	return nums
}

// Splits a slice into a given number of chunks and returns a slice of slices
func splitSlice(nums []int, numOfChunks int) [][]int {
	var chunks [][]int
	for i := 0; i < len(nums); i += numOfChunks {
		end := i + numOfChunks

		// Avoiding splitting beyond slice capacity
		if end > len(nums) {
			end = len(nums)
		}

		chunks = append(chunks, nums[i:end])
	}

	// return a slice of slices containing the subarrays
	return chunks
}

// Takes two sorted slices and merges them into a single sorted slice
func mergeSlices(left, right []int) []int {
	output := make([]int, len(left)+len(right))
	i, j := 0, 0

	// Merge elements from the left and right chunks while maintaining the sorted order.
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			output[i+j] = left[i]
			i++
		} else {
			output[i+j] = right[j]
			j++
		}
	}

	// Append any remaining elements from the left chunk if there are any left
	for i < len(left) {
		output[i+j] = left[i]
		i++
	}

	// Append any remaining elements from the right chunk if there are any left
	for j < len(right) {
		output[i+j] = right[j]
		j++
	}

	// Return the merged slice.
	return output
}
