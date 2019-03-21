package money

// Dollar ...
type Dollar struct {
	*Impl
}

// Times ...
func (d *Dollar) Times(multiplier int) Money {
	money := Impl{}
	return money.NewDollar(d.amount * multiplier)
}
