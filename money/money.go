package money

// Money ...
type Money interface {
	Equals(Money) bool
	Times(int) Money
	Plus(Money) Money
	Amount() int
	Currency() int
}

// Expression ...
type Expression interface {
	Plus(Money) Money
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

// Plus ...
func (m *Impl) Plus(addend Money) Money {
	return &Impl{amount: m.amount + addend.Amount(), currency: m.currency}
}

// Times ...
func (m *Impl) Times(multiplier int) Money {
	//return Dollar(m.amount * multiplier)
	return &Impl{amount: m.amount * multiplier, currency: m.currency}
}

// Amount ...
func (m *Impl) Amount() int {
	return m.amount
}

// Currency ...
func (m *Impl) Currency() int {
	return m.currency
}

// Dollar ...
func Dollar(amount int) *Impl {
	return &Impl{
		amount:   amount,
		currency: USD,
	}
}

// Franc ...
func Franc(amount int) *Impl {
	return &Impl{
		amount:   amount,
		currency: CHF,
	}
}
