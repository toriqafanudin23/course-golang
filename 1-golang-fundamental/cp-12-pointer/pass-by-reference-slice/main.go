package main

import "fmt"

func main() {
	var numbers = []int{4, 6, 7, 1}
	var anotherNumbers = numbers
	fmt.Println(numbers)
	fmt.Println(anotherNumbers)

	// Slice memiliki sifat pass by reference
	numbers[1] = 10
	fmt.Println(numbers)
	fmt.Println(anotherNumbers)

	multiplyBy10(numbers)
	fmt.Println(numbers)

}

func multiplyBy10(numbers []int) {
	for i := range numbers {
		numbers[i] = numbers[i] * 10
	}
	fmt.Println("numbers di function:", numbers)
}
