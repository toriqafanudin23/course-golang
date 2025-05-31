package main

import (
	"fmt"
)

func main() {
	animalChannel := make(chan string)
	fruitsChannel := make(chan string)

	go func() {
		animalChannel <- "Monkey"
		animalChannel <- "Cat"
		animalChannel <- "Fish"
		animalChannel <- "Elephant"
		animalChannel <- "Bird"
		close(animalChannel)
	}()

	go func() {
		fruitsChannel <- "Banana"
		fruitsChannel <- "Apple"
		fruitsChannel <- "Grapes"
		fruitsChannel <- "Watermelon"
		fruitsChannel <- "Orange"
		close(fruitsChannel)
	}()

	fmt.Println("Proses")

	var animalStatus bool
	var fruitsStatus bool

	for {
		select {
		case data, status := <-animalChannel:
			animalStatus = status
			if animalStatus {
				fmt.Println("Animal : " + data)

			}
		case data, status := <-fruitsChannel:
			fruitsStatus = status
			if fruitsStatus {
				fmt.Println("Fruit  : " + data)

			}
		}
		if !animalStatus && !fruitsStatus {
			break
		}
	}

	fmt.Println("Done")
}
