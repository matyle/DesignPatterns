package ch8

type Operation[T float64 | int64 | float32 | int] interface {
	SetNumber1(number T)
	SetNumber2(number T)
	GetResult() T
}

type OperationBase[T float64 | int64 | float32 | int] struct {
	Number1 T
	Number2 T
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

// 工厂模式
type IFactory[T float64 | int64 | float32 | int] interface {
	CreateOperation() Operation[T]
}
type AddFactory[T float64 | int64 | float32 | int] struct {
}

func (a *AddFactory[T]) CreateOperation() Operation[T] {
	return &OperationAdd[T]{}
}

type SubFactory[T float64 | int64 | float32 | int] struct {
}

func (a *SubFactory[T]) CreateOperation() Operation[T] {
	return &OperationSub[T]{}
}

type MulFactory[T float64 | int64 | float32 | int] struct {
}

func (a *MulFactory[T]) CreateOperation() Operation[T] {
	return &OperationMul[T]{}
}

type DivFactory[T float64 | int64 | float32 | int] struct {
}

func (a *DivFactory[T]) CreateOperation() Operation[T] {
	return &OperationDiv[T]{}
}

// client
func FactoryRun[T float64 | int64 | float32 | int](num1, num2 T) T {
	factory := &AddFactory[T]{}
	operation := factory.CreateOperation()
	operation.SetNumber1(num1)
	operation.SetNumber2(num2)
	return operation.GetResult()
}
