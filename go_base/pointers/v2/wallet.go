package main

import "fmt"

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

func (w *Wallet) WithDraw(num Bitcoin) {
	w.balance -= num
}