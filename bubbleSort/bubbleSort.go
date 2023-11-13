package bubbleSort

import (
	"sync"
)

func ConcurrentBubbleSort(nums []int) []int {
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		left = bubbleSort(left, &wg)
	}()
	go func() {
		right = bubbleSort(right, &wg)
	}()

	wg.Wait()

	return combineSlices(left, right)
}

// Creating a parallel version of the bubble sort algorithm involves
// dividing the array into subarrays and sorting each subarray concurrently.
func bubbleSort(nums []int, wg *sync.WaitGroup) []int {
	len := len(nums)

	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	defer wg.Done()
	return nums
}

func combineSlices(left, right []int) []int {
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
