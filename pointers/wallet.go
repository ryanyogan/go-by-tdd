package pointers

import (
	"errors"
	"fmt"
)

// Stringer is an override for fmt package casting our values to strings
type Stringer interface {
	String() string
}

// Bitcoin aliases the int type.
type Bitcoin int

// String overrides %s to return [value] BTC
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet holds the balance of Bitcoin as state.
type Wallet struct {
	balance Bitcoin
}

// Deposit takes an amount in Bitcoin and adds it to a Wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds ..
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw takes an amount in Bitcoin and subtracts it from a Wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance returns the balance of money in a Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
