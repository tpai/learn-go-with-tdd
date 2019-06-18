package main

import "testing"

func TestCountryCode(t *testing.T) {
	countries := Countries{"TW": "Taiwan"}

	t.Run("known country", func(t *testing.T) {
		got, _ := countries.Search("TW")
		expected := "Taiwan"

		assertStrings(t, got, expected)
	})

	t.Run("unknown country", func(t *testing.T) {
		_, err := countries.Search("US")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new country", func(t *testing.T) {
		countries := make(Countries)
		code := "US"
		country := "United States"
		err := countries.Add(code, country)

		assertNoError(t, err)
		assertDefinition(t, countries, code, country)
	})

	t.Run("existing country", func(t *testing.T) {
		code := "US"
		country := "United States"
		countries := Countries{code: country}
		err := countries.Add("US", "Uh Stop")

		assertError(t, err, ErrCountryExist)
		assertDefinition(t, countries, code, country)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("country exist", func(t *testing.T) {
		code := "TW"
		country := "Chinese Taipei"
		countries := Countries{code: country}
		correctCountry := "Taiwan"
		err := countries.Update(code, correctCountry)

		assertNoError(t, err)
		assertDefinition(t, countries, code, correctCountry)
	})

	t.Run("has no country", func(t *testing.T) {
		code := "TW"
		country := "Taiwan"
		countries := Countries{}
		err := countries.Update(code, country)

		assertError(t, err, ErrCountryDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	code := "TW"
	countries := Countries{code: "Chinese Taipei"}
	countries.Delete(code)
	_, err := countries.Search(code)
	if err != ErrNotFound {
		t.Errorf("expected '%s' to be deleted", code)
	}
}

func assertDefinition(t *testing.T, countries Countries, code, expected string) {
	t.Helper()
	country, err := countries.Search(code)
	if err != nil {
		t.Fatal("should find added country:", err)
	}
	assertStrings(t, country, expected)
}

func assertStrings(t *testing.T, got string, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected '%s' but got '%s'", expected, got)
	}
}

func assertError(t *testing.T, err, expected error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if err != expected {
		t.Errorf("expected error '%s' but got '%s'", expected, err)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
