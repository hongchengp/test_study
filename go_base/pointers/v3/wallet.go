package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(balance Bitcoin) {
	w.balance = balance
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) WithDraw(num Bitcoin) error {
	if num > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}
	w.balance -= num
	return nil
}