package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	genap := filter(numbers, isEven)
	fmt.Println("Genap:", genap)
	ganjil := filter(numbers, isOdd)
	fmt.Println("Ganjil", ganjil)

}

func filter(numbers []int, f func(int) bool) []int {
	var result []int
	for _, number := range numbers {
		if f(number) {
			result = append(result, number)
		}
	}

	return result
}

func isEven(number int) bool {
	return number%2 == 0
}

func isOdd(number int) bool {
	return number%2 != 0
}
