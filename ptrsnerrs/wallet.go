package wallet

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin // NOTE: This is private.
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// NOTE: This is a global variable
var ErrNotEnoughBalance = errors.New("ErrNotEnoughBalance")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrNotEnoughBalance
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	// NOTE: return (*w).balance is also okay. (*w) struct pointers are dereferenced automatically.
	return w.balance
}
