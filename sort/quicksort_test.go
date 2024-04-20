package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{5, 2, 9, 1, 7}
	expected := []int{1, 2, 5, 7, 9}

	result := QuickSort(arr, 0, len(arr)-1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
