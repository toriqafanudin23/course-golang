package main

import (
	"fmt"
)

func main() {
	var positiveNumber uint8 = 90
	var negatifveNumber int = -90
	var decimalNumber float32 = 3.9345

	fmt.Println("Positif with uint8 :", positiveNumber)
	fmt.Println("Negatif with int :", negatifveNumber)
	fmt.Println("Decimal with float32 :", decimalNumber)

	fmt.Printf("Dua angka dibelakang koma : %.2f \n", decimalNumber)

	var exist = true
	fmt.Printf("exist? %t \n", exist)
	salah := false
	fmt.Println(salah)

	var message = "Halooo"
	fmt.Printf("Faisa, %s", message)
}
