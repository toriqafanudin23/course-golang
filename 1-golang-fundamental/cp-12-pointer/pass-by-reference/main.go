package main

import "fmt"

func main() {
	var x int = 4
	var y *int = &x

	fmt.Println("x:", x)
	fmt.Println("&x:", &x)
	fmt.Println("y:", y)
	fmt.Println("&y:", &y)
	fmt.Println("Nilai reference dari pointer y:", *y)
	*y = *y + 1
	fmt.Println("x:", x)
	var a = 3
	increase(&a)
	fmt.Println("a:", a)
}

func increase(n *int) {
	*n = *n + 1
	fmt.Println("n di function increase:", *n)
}
