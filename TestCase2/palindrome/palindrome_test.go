package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    string
		expected bool
	}{
		{"", true}, // null string in palindrome
		{"Able was I ere I saw Elba", true},
		{"hello", false},
		{"racecar", true},
		{"12321", true},
		{"not a palindrome", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := isPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
