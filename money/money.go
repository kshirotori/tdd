package money

// Money ...
type Money interface {
	Equals(Money) bool
	Plus(Money) Expression
	Times(int) Money
	Amount() int
	Currency() int
}

// Expression ...
type Expression interface {
	Reduce(bank *Bank, to int) Money
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
	//return &Impl{amount: m.Amount() + addend.Amount(), currency: m.Currency()}
	return &Sum{augend: m, addend: addend}
}

// Times ...
func (m *Impl) Times(multiplier int) Money {
	//return Dollar(m.amount * multiplier)
	return &Impl{amount: m.amount * multiplier, currency: m.currency}
}

// Reduce ...
func (m *Impl) Reduce(bank *Bank, to int) Money {
	rate := bank.Rate(m.currency, to)
	return &Impl{amount: m.amount / rate, currency: to}
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
