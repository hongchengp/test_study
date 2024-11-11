package main

import "testing"

func TestWallet(t *testing.T) {

	assertBlance := func(t *testing.T, w *Wallet, want Bitcoin) {
		got := w.Balance()
		if want != got {
			t.Fatalf("want %s, got %s", want, got)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBlance(t, &wallet, want)
	})

	t.Run("WithDraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(20)
		wallet.WithDraw(10)
		want := Bitcoin(10)
		assertBlance(t, &wallet, want)
	})
}

