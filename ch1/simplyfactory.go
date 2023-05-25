package ch1

import "fmt"

type OperationSFactory[T float64 | int64 | float32 | int] interface {
	CreateOperation(operator string) Operation[T]
}

type Operation[T float64 | int64 | float32 | int] interface {
	SetNumber1(number T)
	SetNumber2(number T)
	GetResult() T
}

type OperationBase[T float64 | int64 | float32 | int] struct {
	Number1 T
	Number2 T
}

type simpleFactory[T float64 | int64 | float32 | int] struct {
}

// add operator still need modify this
func (s *simpleFactory[T]) CreateOperation(operator string) Operation[T] {
	switch operator {
	case "+":
		return &OperationAdd[T]{}
	case "-":
		return &OperationSub[T]{}
	case "*":
		return &OperationMul[T]{}
	case "/":
		return &OperationDiv[T]{}
	default:
		return nil
	}
}

func (o *OperationBase[T]) SetNumber1(number T) {
	o.Number1 = number
}

func (o *OperationBase[T]) SetNumber2(number T) {
	o.Number2 = number
}

// add
type OperationAdd[T float64 | int64 | float32 | int] struct {
	OperationBase[T]
}

func (o *OperationAdd[T]) GetResult() T {
	return o.Number1 + o.Number2
}

// sub
type OperationSub[T float64 | int64 | float32 | int] struct {
	OperationBase[T]
}

func (o *OperationSub[T]) GetResult() T {
	return o.Number1 - o.Number2
}

// mul
type OperationMul[T float64 | int64 | float32 | int] struct {
	OperationBase[T]
}

func (o *OperationMul[T]) GetResult() T {
	return o.Number1 * o.Number2
}

// div
type OperationDiv[T float64 | int64 | float32 | int] struct {
	OperationBase[T]
}

func (o *OperationDiv[T]) GetResult() T {
	if o.Number2 == 0 {
		return 0
	}
	return o.Number1 / o.Number2
}

func Run[T float64 | int64 | float32 | int](num1, num2 T, operator string) T {
	simpleFactory := &simpleFactory[T]{}
	oper := simpleFactory.CreateOperation(operator)
	if oper == nil {
		fmt.Println("operator is not supported")
		return 0
	}
	oper.SetNumber1(num1)
	oper.SetNumber2(num2)
	return oper.GetResult()
}
