package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	var five Money
	five = NewDollar(5)
	assert.Equal(t, NewDollar(10), five.Times(2))
	assert.Equal(t, NewDollar(15), five.Times(3))
}

func TestEquarity(t *testing.T) {
	assert.True(t, NewDollar(5).Equals(NewDollar(5)))
	assert.False(t, NewDollar(5).Equals(NewDollar(6)))
	assert.False(t, NewFranc(5).Equals(NewDollar(5)))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, USD, NewDollar(1).Currency())
	assert.Equal(t, CHF, NewFranc(1).Currency())
}

func TestSimpleAddition(t *testing.T) {
	var five, reduced Money
	var sum Expression
	var bank Bank

	five = NewDollar(5)
	sum = five.Plus(five)
	bank = Bank{}
	reduced = bank.Reduced(sum, "USD")
	assert.Equal(t, NewDollar(10), reduced)
}
