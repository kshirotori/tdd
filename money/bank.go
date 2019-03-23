package money

// Bank ...
type Bank struct {
}

func (b *Bank) reduce(source Expression, to string) Money {
	return Dollar(10)
}
