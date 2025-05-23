package main

import "fmt"

type Item struct {
	id    int
	name  string
	price int
}

type ItemNotFoundError struct {
	Id int
}

func (i *ItemNotFoundError) Error() string {
	return fmt.Sprintf("Item dengan id %d tidak ada", i.Id)
}

var items = []Item{
	{id: 1, name: "Laptop", price: 2000000},
	{id: 2, name: "Hp", price: 1000000},
	{id: 3, name: "Jam", price: 150000},
}

func getItemById(id int) (Item, error) {
	for _, item := range items {
		if item.id == id {
			return item, nil
		}
	}
	return Item{}, &ItemNotFoundError{Id: id}
}

func main() {
	result, err := getItemById(4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
