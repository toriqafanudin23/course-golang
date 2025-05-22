package main

import "fmt"

type mesin struct {
	tipe string
	cc   int
}

type kendaraan struct {
	merek string
	tahun int
	model string
	harga int
	mesin
}

func main() {
	/*
		var a kendaraan
		a.merek = "Toyota"
		a.tahun = 1999
		a.model = "Inova"
		a.harga = 700000000
		fmt.Println(a)
		fmt.Println(a.merek)

		var b = kendaraan{}
		b.merek = "Mitsubishi"
		b.tahun = 2007
		b.model = "Kijang"
		b.harga = 25000000
		fmt.Println("b:", b)

		var c = kendaraan{"Suzuki", 2005, "Maspro", 100000000}
		fmt.Println("c:", c)

		var d = kendaraan{tahun: 2004, merek: "Daihatsu", model: "Xenia", harga: 275000000}
		fmt.Println("d:", d)

		var x = kendaraan{"Honda", 2025, "HRV", 350000000}
		var y kendaraan = x
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Printf("alamat x : %p\n", &x)
		fmt.Printf("alamat y : %p\n", &y)

		y.model = "CPR"
		fmt.Println("x:", x)
		fmt.Println("y:", y)

		var z = kendaraan{merek: "HONDA", model: "JADUL", harga: 400000000, tahun: 1945}
		updateKendaraan(z)
		fmt.Println("old kendaraan:", z)

		var z2 *kendaraan = &z
		fmt.Printf("Alamat z: %p \n", &z)
		fmt.Printf("Alamat z2: %p \n", z2)
		z2.model = "CRV" // Kenapa sebuah alamat bisa dirubah nilainya??
		fmt.Println("z:", z)
		fmt.Println("z2:", z2)
		updateKendaraanPointer(&z)
		fmt.Println("Old kendaraan:", z)

		var k = new(kendaraan)
		fmt.Printf("alamat k: %p \n", k)
	*/
	var a = kendaraan{
		merek: "Toyota",
		tahun: 2010,
		model: "Camry",
		harga: 200000000,
		mesin: mesin{
			tipe: "Premium",
			cc:   2000,
		},
	}
	fmt.Println(a)
	fmt.Println(a.mesin.tipe)
	b := kendaraan{
		merek: "Apaya",
	}
	fmt.Println(b)
}

// Pass by value
func updateKendaraan(newKendaraan kendaraan) {
	newKendaraan.merek = "TOYOTA"
	newKendaraan.model = "Modern"
	newKendaraan.harga = 650000000
	newKendaraan.tahun = 2025
	fmt.Println("new kendaraan:", newKendaraan)
}

// Pass by reference
func updateKendaraanPointer(newKendaraan *kendaraan) {
	newKendaraan.merek = "TOYOTA"
	newKendaraan.model = "Modern"
	newKendaraan.harga = 650000000
	newKendaraan.tahun = 2025
	fmt.Println("new kendaraan:", newKendaraan)
}
