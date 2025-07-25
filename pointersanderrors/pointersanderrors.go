package pointersanderrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin // private outside the package
}

type Bank interface {
	Deposit() float64
	Balance() float64
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit %p\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// (*w).balance -> valid go code but Go creaters thought it was cumbersome
// so the language permits w.balance.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("oh no")
	}
	w.balance -= amount
	return nil
}
