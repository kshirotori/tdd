package money

// Sum ...
type Sum struct {
	augend, addend Money
}

// Reduce ...
func (s *Sum) Reduce(bank *Bank, to int) Money {
	amount := s.augend.Amount() + s.addend.Amount()
	return &Impl{amount: amount, currency: to}
}
