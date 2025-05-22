package main

import (
	"access-modifier/library"
	"fmt"
)

func main() {
	fmt.Println("Hour in a day:", library.HourInADay)
	fmt.Println("Name:", library.Name)
	day := 7
	fmt.Println(day, "hari adalah", library.DaysToMinutes(7), "menit")
}
