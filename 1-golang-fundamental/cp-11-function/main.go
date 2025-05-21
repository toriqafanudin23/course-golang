package main

import "fmt"

var age = 23

func main() {
	hello()
	fmt.Println("Afan, umur:", age)
	greet("Mita", 25)
	greet("Naruto", 18)
	greet("Iron Man", 40)
	add(11, 13)
}

func hello() {
	var age = 25 // Variabel global age akan diabaikan
	var name = "Faisa"
	fmt.Println("Hello World!")
	fmt.Println("Hello", name, "umur", age)
}

func greet(name string, age int) {
	fmt.Println("Selamat pagi", name, age)
}

func add(a int, b int) {
	fmt.Println(a, "+", b, "=", a+b)
}
