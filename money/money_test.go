package money

import (
	"testing"
)

func TestMultiplicationSuccess(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	if product.Amount != 10 {
		t.Errorf("$5 * 2 != 10. result is %d", product.Amount)
	}
	product = five.times(3)
	if product.Amount != 15 {
		t.Errorf("$5 * 3 != 15 result is %d", product.Amount)
	}
}
