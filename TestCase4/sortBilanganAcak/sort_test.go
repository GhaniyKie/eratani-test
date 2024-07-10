package sortbilanganacak

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	// Test case dengan array kosong
	input := []int{}
	expected := []int{}
	result := mergeSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test case dengan array satu elemen
	input = []int{5}
	expected = []int{5}
	result = mergeSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test case dengan array beberapa elemen
	input = []int{5, 2, 9, 1, 5, 6}
	expected = []int{1, 2, 5, 5, 6, 9}
	result = mergeSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
