package money

// Franc ...
type Franc struct {
	*Impl
}

// Times ...
func (d *Franc) Times(multiplier int) Money {
	money := Impl{}
	return money.NewFranc(d.amount * multiplier)
}
