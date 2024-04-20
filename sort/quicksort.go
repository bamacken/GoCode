package sort

func QuickSort(arr []int, low, high int) []int {
	if low < high {
		arr, p := partition(arr, low, high) // partition the array and get the pivot index
		arr = QuickSort(arr, low, p-1)      // sort the left side of the pivot
		arr = QuickSort(arr, p+1, high)     // sort the right side of the pivot
	}

	return arr
}

func partition(arr []int, low int, high int) ([]int, int) {
	// use the last element as the pivot
	pivot := arr[high]
	i := low

	// move all elements smaller than the pivot to the left
	for j := low; j < high; j++ {
		// if the current element is smaller than the pivot
		if arr[j] < pivot {
			// swap the current element with the element at index i
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	// we do this because we want to place the pivot at its correct position in the sorted array
	// so that all elements to the left of the pivot are smaller than the pivot and all elements to the right of the pivot are greater than the pivot
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
