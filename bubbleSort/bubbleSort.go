package bubbleSort

import "sync"

func BubbleSortDriver(nums []uint64) []uint64 {
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]

	var wg sync.WaitGroup
	ch := make(chan bool)

	wg.Add(2)

	go ConcurrentBubbleSort(left, &wg, ch)
	go ConcurrentBubbleSort(right, &wg, ch)

	wg.Wait()

	<-ch // Wait for the first half to be sorted
	<-ch // Wait for the second half to be sorted
	close(ch)

	return nums
}

//Creating a parallel version of the bubble sort algorithm involves
//dividing the array into subarrays and sorting each subarray concurrently.
func ConcurrentBubbleSort(nums []uint64, wg *sync.WaitGroup, ch chan bool) {
	len := len(nums)

	//Perform Bubble Sort on inputted sub array
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	defer wg.Done()

	ch <- true
}
