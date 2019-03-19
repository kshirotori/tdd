package money

// Dollar ...
type Dollar struct {
	*Impl
}

// NewDollar ...
func NewDollar(amount int) *Dollar {
	return &Dollar{
		&Impl{
			amount:   amount,
			currency: "USD",
		},
	}
}

// Times ...
func (d *Dollar) Times(multiplier int) Money {
	return NewDollar(d.amount * multiplier)
}
