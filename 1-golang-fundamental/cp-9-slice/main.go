package main

import "fmt"

func main() {
	numbersSlice := []int{2, 7, 9, 19, -4}
	fmt.Println(numbersSlice)
	fmt.Println("Panjang slice :", len(numbersSlice))
	fmt.Println("Kapasitas slice :", cap(numbersSlice))

	scores := make([]int, 4, 7)
	fmt.Println(scores)
	scores[0] = 70
	scores[1] = 80
	scores[2] = 90
	scores[3] = 100
	fmt.Println(scores)
	fmt.Println(len(scores))
	fmt.Println(cap(scores))

	scores2 := make([]int, 4) // tidak memberikan parameter ke 3
	fmt.Println(scores2)
	scores2[0] = 70
	scores2[1] = 80
	scores2[2] = 90
	scores2[3] = 100
	fmt.Println(scores2)
	fmt.Println(len(scores2))
	fmt.Println(cap(scores2))

	villain := make([]string, 3, 5)
	villain[0] = "Thanos"
	villain[1] = "Joker"
	villain[2] = "Deidara"
	fmt.Println("Villain:", villain)
	fmt.Println("Panjang:", len(villain))
	fmt.Println("Kapasitas:", cap(villain))

	villain = append(villain, "Itachi")
	fmt.Println("Villain:", villain)
	fmt.Println("Panjang:", len(villain))
	fmt.Println("Kapasitas:", cap(villain))

	villain = append(villain, "Orochimaru", "Pain") // menambahkan sekaligus 2 value
	fmt.Println("Villain:", villain)
	fmt.Println("Panjang:", len(villain))
	fmt.Println("Kapasitas:", cap(villain))

	// Perbedaan array dan Slice
	// Array pass by value
	var numbers = [4]int{90, 1, 12, 19}
	var anotherNumbers = numbers
	fmt.Println("Numbers:", numbers)
	fmt.Println("Another Numbers:", anotherNumbers)
	anotherNumbers[1] = 100
	fmt.Println("Numbers:", numbers)
	fmt.Println("Another Numbers:", anotherNumbers)

	// Slice pass by reference
	var prices = []int{10000, 11000, 12000}
	var anotherPrices = prices
	fmt.Println(prices)
	fmt.Println(anotherPrices)
	anotherPrices[2] = 23000
	fmt.Println(prices)
	fmt.Println(anotherPrices)

	// Membuat slice dari array
	var ages = [4]int{23, 40, 25, 17}
	var sliceAges = ages[1:3]
	fmt.Println(ages)
	fmt.Println(sliceAges)
	fmt.Println(cap(sliceAges))

	// Akan menjadi pass by reference
	sliceAges[0] = 100
	fmt.Println("Ages:", ages)
	fmt.Println("Slice Ages:", sliceAges)

	// Ages akan ikut berubah
	sliceAges = append(sliceAges, 132)
	fmt.Println("Ages:", ages)
	fmt.Println("Slice Ages:", sliceAges)
}
