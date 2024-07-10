package polycarp

import (
	"testing"
)

func Test_IsMultipleOfThree(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{1, false},
		{3, true},
		{6, true},
		{10, false},
		{15, true},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := isMultipleOfThree(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v for input %v, but got %v", tc.expected, tc.input, result)
			}
		})
	}
}

func Test_Polycarp(t *testing.T) {
	// Test case for polycarp function
	start := 1
	end := 10
	expectedOutput := []int{1, 2, 4, 5, 7, 8, 10}

	result := polycarp(start, end)

	// Check if the output matches the expected output
	if len(result) != len(expectedOutput) {
		t.Errorf("Expected output length %d, but got %d", len(expectedOutput), len(result))
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expectedOutput[i] {
			t.Errorf("Expected %d at index %d, but got %d", expectedOutput[i], i, result[i])
		}
	}
}
