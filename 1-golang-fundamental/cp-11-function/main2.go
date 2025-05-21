package main

import "fmt"

func main() {
	fmt.Println(tambah(7, 11))
	result := tambah(50, 20)
	fmt.Println(result)
	result2 := kalikan(7, 8)
	fmt.Println(result2)

	kalkulasi := tambah(5, kalikan(3, 21))
	fmt.Println(kalkulasi)
	total := sum(1, 2, 3, 4, 5, 6)
	fmt.Println("total:", total)

	num := []int{4, 3, 1, 7, 5, 9, 8}
	min, max := minMax(num)
	fmt.Println("min:", min)
	fmt.Println("max:", max)

	angka := []int{20, 10, 40, 60, 90, 5, 15}
	kecil, besar := kecilBesar(angka)
	fmt.Println("kecil:", kecil)
	fmt.Println("besar:", besar)
}

func tambah(a int, b int) int {
	result := a + b
	return result
}

func kalikan(a, b int) int {
	result := a * b
	return result
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func minMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}

	return min, max
}

func kecilBesar(numbers []int) (min int, max int) {
	min = numbers[0]
	max = numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}

	return
}
