package main

import "fmt"

func main() {
	user := map[string]string{
		"username": "faisa",
		"email":    "faisa@gmail.com",
	}
	fmt.Println(user)

	var scores = make(map[string]int)
	fmt.Println(scores)
	scores["java"] = 85
	scores["react"] = 87
	scores["kotlin"] = 90
	fmt.Println(scores)
	fmt.Println(scores["java"])
	fmt.Println(scores["golang"])
	fmt.Println(scores)
	scores["java"] = 100
	fmt.Println(scores)
	delete(scores, "kotlin")
	fmt.Println(scores)

	var names = make(map[int]string)
	names[1] = "superman"
	names[2] = "batman"
	names[3] = "naruto"
	names[4] = "zoro"
	for key, value := range names {
		fmt.Println(key, value)
	}
	for _, value := range names {
		fmt.Println(value)
	}
}
