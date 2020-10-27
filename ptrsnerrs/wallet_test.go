package wallet

import "testing"

func TestWallet(t *testing.T) {
	w := Wallet{}
	w.Deposit(10)
	got := w.Balance()
	want := 10
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
