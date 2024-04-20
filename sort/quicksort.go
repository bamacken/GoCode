package sort

func QuickSort(arr []int, low, high int) []int {
	if low < high {
		arr, p := partition(arr, low, high) // partition the array and get the pivot index
		QuickSort(arr, low, p-1)            // sort the left side of the pivot
		QuickSort(arr, p+1, high)           // sort the right side of the pivot
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
	// place the pivot so that elements to the left are smaller and right are larger
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
