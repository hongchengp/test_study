package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Fatalf("want %s, got %s", want, got)
		}
	})

	t.Run("WithDraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(20)
		wallet.WithDraw(10)
		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Fatalf("want %s, got %s", want, got)
		}
	})
}

