package pointeranderrors

import (
	"errors"
	"fmt"
)

// Bitcoin type
type Bitcoin int

// Wallet type
type Wallet struct {
	balance Bitcoin
}

// Stringer only return string
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit will return the amount of deposit
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw will deduct amount
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}
	w.balance -= amount
	return nil
}

// Balance will return the balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
