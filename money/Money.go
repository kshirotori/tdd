package money

// Money ...
type Money interface {
	Equals(Money) bool
	Times(int) Money
	Amount() int
	Currency() int
}

// const ...
const (
	USD = iota
	CHF
)

// Impl ...
type Impl struct {
	amount   int
	currency int
}

// Equals ...
func (m *Impl) Equals(a Money) bool {
	return m.amount == a.Amount() &&
		m.currency == a.Currency()
}

// Amount ...
func (m *Impl) Amount() int {
	return m.amount
}

// Currency ...
func (m *Impl) Currency() int {
	return m.currency
}

// NewDollar ...
func (m *Impl) NewDollar(amount int) Money {
	return &Dollar{
		&Impl{
			amount:   amount,
			currency: USD,
		},
	}
}

// NewFranc ...
func (m *Impl) NewFranc(amount int) Money {
	return &Franc{
		&Impl{
			amount:   amount,
			currency: CHF,
		},
	}
}
