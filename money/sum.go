package money

type Sum struct {
	augend, addend Money
}

func (s *Sum) Reduce(to int) Money {
	amount := s.augend.Amount() + s.addend.Amount()
	return &Impl{amount: amount, currency: to}
}
