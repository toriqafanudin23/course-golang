package main

import "fmt"

func main() {
	var fruits = []string{"Apel", "Pisang", "Anggur", "Semangka"}
	fmt.Println(fruits)
	fmt.Println(fruits[2])
	fruits[2] = "Melon"
	fmt.Println(fruits)

	var scores [3]int
	scores[0] = 92
	scores[1] = 80
	scores[2] = 17
	fmt.Println(scores)

	var days = [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)
	fmt.Println(len(days)) // panjang array

	for i := 0; i < len(days); i++ {
		fmt.Printf("Elemen %d : %s\n", i, days[i])
	}

	for i, day := range days {
		fmt.Printf("Index %d : %s\n", i, day)
	}

	for _, day := range days { // hanya mengambil value
		fmt.Println(day)
	}

	for i := range days { // hanya index
		fmt.Println(i)
	}

	var numbers = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Sebelum:", numbers)
	for i := 0; i < len(numbers); i++ {
		numbers[i] *= 2
	}
	fmt.Println("Sesudah:", numbers)

	var matrix = [2][3]int{
		{1, 2, 3}, {4, 5, 6},
	}
	fmt.Println("Matrix:", matrix)
	fmt.Println(matrix[1][1])
}
