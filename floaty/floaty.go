package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		fmt.Println("not enough money in the acct")
	} else {
		fmt.Printf("%d withdraw : Balance : %d\n", v, a.balance)
		a.balance -= v
	}
}

func main() {
	var acct Account

	acct.balance = 100
	fmt.Println("Balance :", acct.GetBalance())

	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)

}
