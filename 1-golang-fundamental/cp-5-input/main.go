package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var firstName string
	var lastName string
	var age int

	fmt.Print("Enter your name : ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Print("Enter your age : ")
	fmt.Scanln(&age)
	birthYear := (2025 - age)

	fmt.Println("Your name is", firstName, lastName)
	fmt.Println("Your are born in", birthYear)

	fmt.Println("================")
	var fullName string
	var myAge int

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukan nama anda : ")
	scanner.Scan()
	fullName = scanner.Text()

	fmt.Print("Masukan umur anda : ")
	scanner.Scan()
	myAge, _ = strconv.Atoi(scanner.Text())

	fmt.Printf("%s | %d", fullName, myAge)
}
