package pointers_and_errors

import (
	"errors"
	"fmt"
)

type Bitcoins int

func (b Bitcoins) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoins
}

func (w *Wallet) Deposit(amount Bitcoins) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoins {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) withdraw(amount Bitcoins) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
