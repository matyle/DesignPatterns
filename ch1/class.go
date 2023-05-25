package ch1

type Operations[T float64 | int64] interface {
	Add() T
	Sub() T
	Mul() T
	Div() T
}

type Calcultor[T float64 | int64] struct {
	Number1 T
	Number2 T
}

func (c Calcultor[T]) Add() T {
	return c.Number1 + c.Number2
}
func (c Calcultor[T]) Sub() T {
	return c.Number1 - c.Number2
}
