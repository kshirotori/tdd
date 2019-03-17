package money

// Dollar ...
type Dollar struct {
	Amount int
}

// NewDollar ...
func NewDollar(amount int) *Dollar {
	Dollar := Dollar{Amount: amount}
	return &Dollar
}

func (d *Dollar) times(multiplier int) *Dollar {
	return NewDollar(d.Amount * multiplier)
}
