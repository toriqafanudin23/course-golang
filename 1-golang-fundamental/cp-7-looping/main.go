package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	data := []string{"apel", "jeruk", "mangga"}
	for index, value := range data {
		fmt.Println(index, value)
	}

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // loop berhenti saat i = 5
		}
		fmt.Println(i)
	}

	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // lompat iterasi saat i = 3
		}
		fmt.Println(i)
	}

}
