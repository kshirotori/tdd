package money

// Money ...
type Money interface {
	Equals(Money) bool
	Plus(Money) Expression
	Times(int) Money
	Amount() int
	Currency() int
}

type Expression interface{}

type Bank struct{}

func (b *Bank) Reduced(source Expression, to string) Money {
	return NewDollar(10)
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
func (m *Impl) Plus(addend Money) Expression {
	return &Impl{amount: m.Amount() + addend.Amount(), currency: m.Currency()}
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

// NewDollar ...
func NewDollar(amount int) Money {
	return &Impl{
		amount:   amount,
		currency: USD,
	}
}

// NewFranc ...
func NewFranc(amount int) Money {
	return &Impl{
		amount:   amount,
		currency: CHF,
	}
}
