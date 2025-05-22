package main

import "fmt"

func main() {
	name := "Faisa"
	fmt.Println(name)
	fmt.Println(&name)

	age := 24
	fmt.Println(age)
	fmt.Println(&age)

	// Pass by value
	var x = 4
	var y = x //Menyalin nilai, alamat memorinya berbeda
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("&x:", &x)
	fmt.Println("&y:", &y)
	y = y + 1
	fmt.Println("y:", y)   // Valuenya berubah
	fmt.Println("&y:", &y) // Alamatnya tetap sama

	var a = 4
	increase(a) // Function juga memiliki sifat pass by value
	fmt.Println("a:", a)
	fmt.Println("&a:", &a)

	var numbers = [4]int{4, 7, 3, 1}
	var anotherNumbers = numbers
	fmt.Println(numbers)
	fmt.Println(anotherNumbers)
	numbers[1] = 0
	fmt.Println(numbers)
	fmt.Println(anotherNumbers)

	var scores = [4]int{3, 8, 2, 5}
	multiplyBy10(scores)
	fmt.Println(scores)
	fmt.Println(&scores[0], &scores[1], &scores[2], &scores[3])
}

func multiplyBy10(n [4]int) {
	for i := range n {
		n[i] = n[i] * 10
	}
	fmt.Println(n)
}

func increase(n int) {
	n = n + 1
	fmt.Println("n:", n)
	fmt.Println("&n:", &n)
}
