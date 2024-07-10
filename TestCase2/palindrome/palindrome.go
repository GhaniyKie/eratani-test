package palindrome

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func Run() {
	fmt.Println("Masukkan sebuah string:")
	var inputString string
	fmt.Scanln(&inputString)

	if isPalindrome(inputString) {
		fmt.Printf("%s adalah palindrome. \n", inputString)
	} else {
		fmt.Printf("%s bukan palindrome. \n", inputString)
	}
}
