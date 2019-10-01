package pointeranderrors

// Wallet type
type Wallet struct {
	balance int
}

// Balance will return the balance
func (w *Wallet) Balance() int {
	return w.balance
}

// Deposit will return the amount of deposit
func (w *Wallet) Deposit(amount int) int {
	w.balance += amount
	return w.balance
}
