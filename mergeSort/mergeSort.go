package mergeSort

import (
	"sync"
)

func ConcurrentMergeSort(nums []uint64) []uint64 {
	// If the lenght of input is less than or equal to 1 return input
	if len(nums) <= 1 {
		return nums
	}

	if len(nums) <= 2000 {
		return sequentialMergeSort(nums)
	} else {
		mid := len(nums) / 2

		var left, right []uint64

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			left = ConcurrentMergeSort(nums[:mid])
			wg.Done()
		}()

		go func() {
			right = ConcurrentMergeSort(nums[mid:])
			wg.Done()
		}()

		wg.Wait()

		return mergeSlices(left, right)
	}
}

func sequentialMergeSort(nums []uint64) []uint64 {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := sequentialMergeSort(nums[:mid])
	right := sequentialMergeSort(nums[mid:])

	return mergeSlices(left, right)
}

func mergeSlices(left, right []uint64) []uint64 {
	output := make([]uint64, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			output[i+j] = left[i]
			i++
		} else {
			output[i+j] = right[j]
			j++
		}
	}

	for i < len(left) {
		output[i+j] = left[i]
		i++
	}

	for j < len(right) {
		output[i+j] = right[j]
		j++
	}

	return output
}
