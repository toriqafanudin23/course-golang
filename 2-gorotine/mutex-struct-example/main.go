package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	account := BankAccount{Balance: 0}
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go account.topUp(10)
	}
	wg.Wait()
	fmt.Println(account.Balance)
}

type BankAccount struct {
	Balance int
	mutex   sync.Mutex
}

func (account *BankAccount) topUp(amount int) {
	defer wg.Done()
	defer account.mutex.Unlock()

	account.mutex.Lock()
	account.Balance += amount
}
