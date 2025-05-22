package main

import "fmt"

func main() {
	fmt.Println("Operator Aritmatika")
	x := 9
	y := 4
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)
	fmt.Println(3 + 4*8 - 11)
	fmt.Println("\nOperator Assignment")
	var a int = 10
	var b int = 3
	a = a - 3
	fmt.Println(a)
	a += b
	fmt.Println(a)
	a -= b
	fmt.Println(a)
	a *= b
	fmt.Println(a)
	a /= b
	fmt.Println(a)
	a %= b
	fmt.Println(a)

	fmt.Println("\nOperator Perbandingan")
	var c int = 11
	var d int = 10

	fmt.Println(c, "==", d, "maka", c == d)
	fmt.Println(c, "!=", d, "maka", c != d)
	fmt.Println(c, "<=", d, "maka", c <= d)
	fmt.Println(c, ">=", d, "maka", c >= d)
	fmt.Println(c, "<", d, "maka", c < d)
	fmt.Println(c, ">", d, "maka", c > d)

	fmt.Println("\nOperator Logika")
	f := true
	g := false
	fmt.Println(f, "&&", g, "maka", f && g)
	fmt.Println(f, "||", g, "maka", f || g)
	fmt.Println("!true maka", !f)
	fmt.Println("false || true && false maka", g || f && g)
	// && dijalankan terlebih dahulu

	j := 100
	j++
	fmt.Println(j)

}
