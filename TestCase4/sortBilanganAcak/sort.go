package sortbilanganacak

import "fmt"

// Buat input berupa integer
func mergeSort(arr []int) []int {
	// Hitung total input dan balikan kondisinya jika sesuai
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2           //Hitung indeks array input dengan dibagi menjadi 2
	left := mergeSort(arr[:mid])  // Panggil function untuk mengurutkan dari kiri ke bagian titik dengan indeks
	right := mergeSort(arr[mid:]) // Panggil function untuk mengurutkan dari kanan ke bagian titik dengan indeks

	// Kembalikan nilainya yang sudah dihitung dan dibagi dengan dibagi 2
	return merge(left, right)
}

func merge(left, right []int) []int {
	// Buat slice kosong dengan panjang yang sesuai dengan input
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0 //Definisikan 2 variable indeks untuk left dan right

	// Jika variable l dan r sesuai dengan panjang indeks, masukan datanya kedalam slice kosong
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	// Tambahkan hasil kedalam slice kosong
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

func Run() {
	// Input bilangan acak
	numbers := []int{5, 2, 9, 1, 5, 6}

	// Menampilkan bilangan acak sebelum diurutkan
	fmt.Println("Sebelum diurutkan:", numbers)

	// Mengurutkan bilangan acak menggunakan Merge Sort
	sortedNumbers := mergeSort(numbers)

	// Menampilkan bilangan acak setelah diurutkan
	fmt.Println("Setelah diurutkan:", sortedNumbers)
}
