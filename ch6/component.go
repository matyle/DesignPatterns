package ch6

import "fmt"

// Component abstract class
type Component[T any] interface {
	Operation() T
}

type ConcreteComponent[T any] struct {
}

func (c *ConcreteComponent[T]) Operation() T {
	var r T
	fmt.Println("ConcreteComponent Operation")
	return r
}

// Decorators contain components, they are aggregation relationships

type DecoratorsBase[T any] struct {
	component Component[T]
}

func (d *DecoratorsBase[T]) SetComponent(cp Component[T]) {
	d.component = cp
}

func (d *DecoratorsBase[T]) Operation() T {
	var r T
	if d.component != nil {
		r = d.component.Operation()
	}
	fmt.Println("DecoratorsBase Operation")
	return r
}

type ConcreteDecoratorsA[T any] struct {
	base       DecoratorsBase[T]
	addedState string
}

func (d *ConcreteDecoratorsA[T]) Operation() T {
	r := d.base.Operation()
	d.addedState = "newState"
	fmt.Println("ConcreteDecoratorsA Operation")
	return r
}

type ConcreteDecoratorsB[T any] struct {
	base DecoratorsBase[T]
}

func (d *ConcreteDecoratorsB[T]) Operation() T {
	r := d.base.Operation()
	fmt.Println("ConcreteDecoratorsB Operation")
	return r
}

func (d *ConcreteDecoratorsB[T]) AddBehavior() {
	fmt.Println("ConcreteDecoratorsB AddBehavior")
}

// client
func DecoratorRun() {
	c := new(ConcreteComponent[int64])
	d1 := new(ConcreteDecoratorsA[int64])
	d2 := new(ConcreteDecoratorsB[int64])

	d1.base.SetComponent(c)
	d2.base.SetComponent(c)

	d1.Operation()
	fmt.Println(".......")
	d2.Operation()

}
