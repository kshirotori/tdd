package money

import (
	"reflect"
)

// Money ...
type Money interface {
	Equals(Money, Money) bool
	Times(int) Money
	Currency() string
	Amount() int
}

// Equals ...
func Equals(a Money, b Money) bool {
	return a.Amount() == b.Amount() &&
		reflect.TypeOf(a) == reflect.TypeOf(b)
}

// Impl ...
type Impl struct {
	Money
	amount   int
	currency string
}

// Amount ...
func (m *Impl) Amount() int {
	return m.amount
}

// Currency ...
func (m *Impl) Currency() string {
	return m.currency
}
