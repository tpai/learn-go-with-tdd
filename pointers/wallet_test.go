package main

import "fmt"
import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	}

	assertError := func(t *testing.T, err error, expected error) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if err != expected {
			t.Errorf("got '%s', expected '%s'", err, expected)
		}
	}

	assertNoError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		fmt.Printf("balance in test is %v\n", &wallet.balance)
		expected := Bitcoin(10)

		assertBalance(t, wallet, expected)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))
		expected := Bitcoin(10)

		assertBalance(t, wallet, expected)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(30))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
