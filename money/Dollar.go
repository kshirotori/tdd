package money

// Dollar ...
type Dollar struct {
	amount int
}

// NewDollar ...
func NewDollar(amount int) *Dollar {
	Dollar := Dollar{amount: amount}
	return &Dollar
}

// Times ...
func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

// Equals ...
func (d *Dollar) Equals(dollar *Dollar) bool {
	return d.amount == dollar.amount
}
