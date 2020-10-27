package wallet

import "testing"

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertThrowsError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Expected an error, did not get one.")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, w, want)
	})

	t.Run("Withdraw when enough balance", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(10)}
		w.Withdraw(Bitcoin(10))
		want := Bitcoin(0)
		assertBalance(t, w, want)
	})

	t.Run("Withdraw when below balance", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		w := Wallet{balance: startingBalance}
		err := w.Withdraw(Bitcoin(40))
		assertBalance(t, w, startingBalance)
		assertThrowsError(t, err, ErrNotEnoughBalance)
	})
}
