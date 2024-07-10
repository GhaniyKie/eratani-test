package polycarp

import "fmt"

func isMultipleOfThree(num int) bool {
	return num%3 == 0
}

func polycarp(start, end int) []int {
	var result []int
	for i := start; i <= end; i++ {
		if !isMultipleOfThree(i) {
			result = append(result, i)
		}
	}
	return result
}

func Run() {
	start := 1
	end := 1000
	output := polycarp(start, end)

	// Cetak elemen-elemen dari slice sebagai output integer
	for _, num := range output {
		fmt.Println(num)
	}
}
