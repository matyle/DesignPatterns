package ch2

import (
	"fmt"
	"math"
)

type CashStrategyer[T float64 | int64] interface {
	CalculateCash() T
}

// CashSuper
type CashSuper[T float64 | int64] struct {
	amount T
	price  T
}

// CashNormal
type CashNormal[T float64 | int64] struct {
	CashSuper[T]
}

func (c *CashNormal[T]) New(amount, price T) {
	c.price = price
	c.amount = amount
}

func (c *CashNormal[T]) CalculateCash() T {
	return c.amount * c.price
}

// CashRebate
type CashRebate[T float64 | int64] struct {
	CashSuper[T]
	rebate float64 //rebate
}

func (c *CashRebate[T]) New(amount, price T, reba float64) {
	c.rebate = reba
	c.price = price
	c.amount = amount
}

func (c *CashRebate[T]) CalculateCash() T {
	return c.amount * c.price * T(c.rebate)
}

// CashReturn
type CashReturn[T float64 | int64] struct {
	CashSuper[T]
	moneyCondition T
	moneyReturn    T
}

func (c *CashReturn[T]) New(moneyCondition, moneyReturn, amount, price T) {
	c.moneyCondition = moneyCondition
	c.moneyReturn = moneyReturn
	c.price = price
	c.amount = amount
}

func (c *CashReturn[T]) CalculateCash() T {
	money := c.amount * c.price
	var result T
	if c.moneyCondition <= money {
		result = money - T(math.Floor(float64(money/c.moneyCondition)))*c.moneyReturn
	}
	return result
}

// context
type CashContext[T float64 | int64] struct {
	strategy CashStrategyer[T]
}

func NewCashContext[T float64 | int64](strtegy CashStrategyer[T]) *CashContext[T] {
	context := new(CashContext[T])
	context.strategy = strtegy
	return context
}

func (this *CashContext[T]) CashContextFactory(conType string) {
	switch conType {
	case "Normal":
		cn := &CashNormal[T]{}
		cn.New(10, 3)
		this.strategy = cn
	case "Rebate":
		cn := &CashRebate[T]{}
		cn.New(10, 3, 0.8)
		this.strategy = cn
	}
}

func (this *CashContext[T]) GetCashResult() T {
	return this.strategy.CalculateCash()
}

func CashRun() {
	cn := &CashNormal[float64]{}
	cn.New(10, 3)
	context := NewCashContext[float64](cn)
	res := context.strategy.CalculateCash()
	fmt.Println("Normal:", res)

	cr := &CashRebate[float64]{}
	cr.New(10, 3, 0.8)
	context = NewCashContext[float64](cr)
	res = context.strategy.CalculateCash()
	fmt.Println("Rebate:", res)

	//strategy+simplyfactory
	const (
		NORMAL = "Normal"
	)
	contextNormal := new(CashContext[float64])
	contextNormal.CashContextFactory(NORMAL)
	res = contextNormal.GetCashResult()
	fmt.Println("Normal 2:", res)
}
