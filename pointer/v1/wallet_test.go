package pointer

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, got Bitcoin, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got, want error) {
		t.Helper()
		if got == nil && got == want {
			return
		}

		if got == nil || want == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got.Error() != want.Error() {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)

		assertError(t, err, nil)
		assertBalance(t, got, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, startingBalance, Bitcoin(20))
	})
}
