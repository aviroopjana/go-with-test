package main

import "fmt"

//Bitcoin represent a number of bitcoins
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Add Bitcoin to wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns a number of Bitcoin a wallet has
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
