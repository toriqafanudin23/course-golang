package main

import "fmt"

func main() {
	/*
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Terjadi panic :", r)
			}
		}()
	*/

	var input string
	fmt.Print("Masukan nama : ")
	fmt.Scanln(&input)

	fmt.Println("Start")
	if !isEmpty(input) {
		fmt.Println(input)
	}
	fmt.Println("Done")

}

func isEmpty(input string) (empty bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi panic :", r)
			empty = true
		}
	}()
	if input == "" {
		panic("inputan tidak boleh kosong")
	}
	empty = false
	return
}
