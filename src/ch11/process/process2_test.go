package process

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Account struct {
	money int
	lock sync.Mutex
}

func (a *Account) SetMoney(m int) {
	a.money += m
}

func (a *Account) Spend(m int) {
	a.money -= m
}

func (a *Account) GetMoney() int {
	return a.money
}

func (a *Account) Check() {
	time.Sleep(1 * time.Second)
}

func TestBuy(t *testing.T) {
	var account Account
	account.SetMoney(10)

	go func(bm int) {
		//account.lock.Lock()
		if account.money > bm {
			account.Check()
			account.Spend(bm)
		}
		//account.lock.Unlock()
	}(5)

	go func(bm int) {
		//account.lock.Lock()
		if account.money > bm {
			account.Check()
			account.Spend(bm)
		}
		//account.lock.Unlock()
	}(6)
	time.Sleep(2 * time.Second)
	fmt.Println(account.GetMoney())
}