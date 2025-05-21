package main

import "fmt"

func main() {
	// Anonymus function
	func() {
		fmt.Println("Hello World!")
	}()

	// Anonymus function dalam variabel
	halo := func() {
		fmt.Println("Halo Duniaa..")
	}
	halo()

	// Passing argument
	func(word string) {
		fmt.Println(word)
	}("Sebuah Kata-kata Indah")

	// Passing argument dalam Anonymus function
	hello := func(word string) {
		fmt.Println("Hello", word)
	}
	hello("Faisa")
}
