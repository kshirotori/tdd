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
	reduced = bank.Reduce(sum, USD)
	assert.Equal(t, Dollar(10), reduced)
}

func TestPlusReturnSum(t *testing.T) {
	var five Money
	var result Expression
	var sum *Sum

	five = Dollar(5)
	result = five.Plus(five)
	sum = result.(*Sum)
	assert.Equal(t, five, sum.augend)
	assert.Equal(t, five, sum.addend)
}

func TestReduceSum(t *testing.T) {
	var bank Bank
	var sum Expression
	var result Money

	sum = &Sum{Dollar(3), Dollar(4)}
	bank = Bank{}
	result = bank.Reduce(sum, USD)
	assert.Equal(t, Dollar(7), result)
}

func TestReduceMoney(t *testing.T) {
	bank := Bank{}
	result := bank.Reduce(Dollar(1), USD)
	assert.Equal(t, Dollar(1), result)
}

func TestReduceFlancToDollar(t *testing.T) {
	bank := Bank{}
	bank.AddRate(CHF, USD, 2)
	dollar := bank.Reduce(Franc(2), USD)
	assert.Equal(t, Dollar(1), dollar)
}
