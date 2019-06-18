package main

type Countries map[string]string
type CountriesErr string

func (err CountriesErr) Error() string {
	return string(err)
}

const (
	ErrNotFound            = CountriesErr("could not find any country you were looking for")
	ErrCountryExist        = CountriesErr("cannot add country because it exist")
	ErrCountryDoesNotExist = CountriesErr("cannot update country because it does not exist")
)

func (c Countries) Search(code string) (string, error) {
	country, ok := c[code]
	if !ok {
		return "", ErrNotFound
	}
	return country, nil
}

func (c Countries) Add(code string, country string) error {
	_, err := c.Search(code)
	switch err {
	case ErrNotFound:
		c[code] = country
	case nil:
		return ErrCountryExist
	default:
		return err
	}
	return nil
}

func (c Countries) Update(code string, country string) error {
	_, err := c.Search(code)
	switch err {
	case ErrNotFound:
		return ErrCountryDoesNotExist
	case nil:
		c[code] = country
	default:
		return err
	}
	return nil
}

func (c Countries) Delete(code string) {
	delete(c, code)
}
