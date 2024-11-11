package main

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(balance int) {
	w.balance = balance
}

func (w *Wallet) Balance() int {
	return w.balance
}