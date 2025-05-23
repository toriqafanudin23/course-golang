package main

import "fmt"

func main() {
	/*num := 4
	defer fmt.Println("Result Defer:", num)

	num += 10
	num *= 2
	fmt.Println("Result Main:", num)

	num1 := 4
	num2 := 0
	defer fmt.Println("Done")
	fmt.Println("Start")
	fmt.Println(num1 / num2) */
	defer add(5, multiply(2, 3))
	add(1, 3)

}

// jika ada lebih dari 1 defer: Last In First Out

func add(num1, num2 int) {
	result := num1 + num2
	fmt.Println(result)
}

func multiply(num1 int, num2 int) int {
	result := num1 * num2
	fmt.Println(result)
	return result
}
