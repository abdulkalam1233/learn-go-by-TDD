package pointers_and_errors

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoins(10))
		assertBalance(t, wallet, Bitcoins(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		startingBalance := Bitcoins(20)
		wallet := Wallet{startingBalance}
		err := wallet.withdraw(Bitcoins(10))
		assertNoErrors(t, err)
		assertBalance(t, wallet, Bitcoins(10))
	})

	t.Run("Withdraw more than balance", func(t *testing.T) {
		startingBalance := Bitcoins(10)
		wallet := Wallet{startingBalance}
		err := wallet.withdraw(Bitcoins(20))
		assertBalance(t, wallet, Bitcoins(10))
		assertErrors(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoins) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertErrors(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoErrors(
	t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got an error but didn't want one")
	}
}
