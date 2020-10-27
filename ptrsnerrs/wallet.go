package wallet

type Wallet struct {
	balance int // NOTE: This is private.
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	// NOTE: return (*w).balance is also okay. (*w) struct pointers are dereferenced automatically.
	return w.balance
}
