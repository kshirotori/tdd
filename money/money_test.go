package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplicationSuccess(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	assert.Equal(t, NewDollar(10), product)
	product = five.Times(3)
	assert.Equal(t, NewDollar(15), product)
}

func TestEquarity(t *testing.T) {
	if NewDollar(5).Equals(NewDollar(5)) != true {
		t.Errorf("Equals is not true")
	}

	if NewDollar(5).Equals(NewDollar(6)) != false {
		t.Errorf("Equals must have faild")
	}
}
