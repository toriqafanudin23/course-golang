package main

import "fmt"

func main() {
	if false {
		fmt.Println("Kode dijalankan")
	}

	fmt.Println("Done!")

	if result := "lulus"; result == "lulus" {
		fmt.Println("Selamat, anda " + result)
	} else {
		fmt.Println("Maaf anda " + result)
	}

	x := 20

	if x > 8 && x < 25 {
		fmt.Println("Nilai X masih dalam rentang")
	} else {
		fmt.Println("Nilai X keluar dari rentang")
	}

	GPA := 3.2
	var graduationStatus string

	if GPA > 3.6 {
		graduationStatus = "Cumlaude"
	} else if GPA > 3 {
		graduationStatus = "Memuaskan"
	} else {
		graduationStatus = "Kurang memuaskan"
	}

	fmt.Printf("Selamat kamu lulus dengan predikat %s dengan IPK %.2f", graduationStatus, GPA)

	a := 3
	b := -1

	if a > 0 {
		if b > 0 {
			fmt.Println("\nnilai a dan b lebih dari nol")
		} else {
			fmt.Println("\nnilai a lebih dari nol dan b kurang dari nol")
		}
	}
	fmt.Print("Program selesai, I love pemrograman!\n")
	fmt.Print("Done!")
}
