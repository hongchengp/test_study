package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, w *Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if want != got {
			t.Fatalf("want %s, got %s", want, got)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, &wallet, want)
	})

	t.Run("WithDraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(20)
		wallet.WithDraw(10)
		want := Bitcoin(10)
		assertBalance(t, &wallet, want)
	})

	
	t.Run("WithDraw overload fund", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(20)
		err := wallet.WithDraw(Bitcoin(100))
		assertBalance(t, &wallet, Bitcoin(20))
		if err == nil {
			t.Error("want an err, but err is")
		}
	})
}

