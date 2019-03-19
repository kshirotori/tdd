package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), five.Times(2))
	assert.Equal(t, NewDollar(15), five.Times(3))
}

func TestEquarity(t *testing.T) {
	assert.True(t, Equals(NewDollar(5), NewDollar(5)))
	assert.False(t, Equals(NewDollar(5), NewDollar(6)))
	assert.True(t, Equals(NewFranc(5), NewFranc(5)))
	assert.False(t, Equals(NewFranc(5), NewFranc(6)))
	assert.False(t, Equals(NewFranc(5), NewDollar(5)))
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, NewFranc(10), five.Times(2))
	assert.Equal(t, NewFranc(15), five.Times(3))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(1).Currency())
	assert.Equal(t, "CHF", NewFranc(1).Currency())
}
