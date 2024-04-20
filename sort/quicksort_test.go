package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("Sort an array of integers", func(t *testing.T) {
		arr := []int{5, 2, 9, 3, 7, 6, 1, 8, 4}
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		result := QuickSort(arr, 0, len(arr)-1)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Empty array", func(t *testing.T) {
		arr := []int{}
		expected := []int{}

		result := QuickSort(arr, 0, len(arr)-1)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Array with one element", func(t *testing.T) {
		arr := []int{5}
		expected := []int{5}

		result := QuickSort(arr, 0, len(arr)-1)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Array with two elements", func(t *testing.T) {
		arr := []int{5, 2}
		expected := []int{2, 5}

		result := QuickSort(arr, 0, len(arr)-1)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
