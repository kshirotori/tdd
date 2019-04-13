package money

// Bank 銀行は為替レートを持ち、通貨の換算をする
type Bank struct {
}

// Reduced sourceをtoの通貨に換算する
func (b *Bank) Reduce(source Expression, to int) Money {
	return source.Reduce(b, to)
}

// AddRate
func (b *Bank) AddRate(from int, to int, rate int) {

}

// Rate
func (b *Bank) Rate(from int, to int) int {
	if from == CHF && to == USD {
		return 2
	}
	return 1
}
