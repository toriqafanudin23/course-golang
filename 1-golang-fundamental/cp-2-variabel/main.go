package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	firstName = "Faisa"
	lastName = "Nirbita"
	fmt.Println(firstName, lastName)
	fmt.Printf("Halo, selamat datang Faisa\n")
	fmt.Printf("Halo, %s %s \n", firstName, lastName)

	var (
		age     int
		address string
	)

	age = 25
	address = "Jl. Ronggolawe no 9"

	fmt.Printf("Halo, nama saya %s, saya berumur %d tahun. Saya sekarang tinggal di %s \n", firstName, age, address)
	var (
		bootcampName, bootcampAddress = "Enigma Camp", "Jakarta Selatan"
	)
	fmt.Println(bootcampName, bootcampAddress)

	day := "Monday"
	date := "20"
	month := "Desember"
	fmt.Println(day, date, month+" 2022")

	const phi = 3.14
	bootcamp, _ := "Enigma", "Aktif7"
	fmt.Println(bootcamp, phi)
}
