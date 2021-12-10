package point

import (
	"errors"
	"fmt"
	"testing"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(i Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += i
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(b Bitcoin) error {
	if w.balance < b {
		return errors.New("oh no")
	}
	w.balance -= b
	return nil
}

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	want := Bitcoin(101)

	if got != want {
		t.Errorf("got %d want %s", got, want)
	}

}
