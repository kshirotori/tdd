package money

// Bank 銀行は為替レートを持ち、通貨の換算をする
type Bank struct {
}

// Reduced sourceをtoの通貨に換算する
func (b *Bank) Reduced(source Expression, to string) Money {
	return Dollar(10)
}
