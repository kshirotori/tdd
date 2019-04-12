package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	var five Money
	five = Dollar(5)
	assert.Equal(t, Dollar(10), five.Times(2))
	assert.Equal(t, Dollar(15), five.Times(3))
}

func TestEquarity(t *testing.T) {
	assert.True(t, Dollar(5).Equals(Dollar(5)))
	assert.False(t, Dollar(5).Equals(Dollar(6)))
	assert.False(t, Franc(5).Equals(Dollar(5)))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, USD, Dollar(1).Currency())
	assert.Equal(t, CHF, Franc(1).Currency())
}

func TestSimpleAddition(t *testing.T) {
	var five, reduced Money
	var sum Expression
	var bank Bank

	five = Dollar(5)
	sum = five.Plus(five)
	bank = Bank{}
	reduced = bank.Reduced(sum, "USD")
	assert.Equal(t, Dollar(10), reduced)
}
