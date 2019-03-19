package money

// Franc ...
type Franc struct {
	*Impl
}

// NewFranc ...
func NewFranc(amount int) *Franc {
	return &Franc{
		&Impl{
			amount:   amount,
			currency: "CHF",
		},
	}
}

// Times ...
func (d *Franc) Times(multiplier int) Money {
	return NewFranc(d.amount * multiplier)
}
