package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := Dollar(5)
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

func TestAddition(t *testing.T) {
	sum := Dollar(5).Plus(Dollar(5))
	assert.Equal(t, Dollar(10), sum)
}
