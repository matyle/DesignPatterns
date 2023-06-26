package ch24

import "fmt"

type Handler interface {
	Request(int)
	SetSuccessor(handler Handler)
}

type ConcreteHandlerLower struct {
	successor Handler
}

func (c *ConcreteHandlerLower) Request(request int) {
	if request > 0 && request < 10 {
		fmt.Println("ConcreteHandlerLower can handle it")
	} else if c.successor != nil {
		fmt.Println("ConcreteHandlerLower cannot handle move to the successor")
		c.successor.Request(request)
	} else {
		fmt.Println("ConcreteHandlerLower cannot handle,other is nil")
	}
}

func (c *ConcreteHandlerLower) SetSuccessor(h Handler) {
	c.successor = h
}

type ConcreteHandlerMiddle struct {
	successor Handler
}

func (c *ConcreteHandlerMiddle) Request(request int) {
	if request >= 10 && request < 20 {
		fmt.Println("ConcreteHandlerMiddle can handle it")
	} else if c.successor != nil {
		fmt.Println("ConcreteHandlerMiddle cannot handle move to the successor")
		c.successor.Request(request)
	}
}

func (c *ConcreteHandlerMiddle) SetSuccessor(h Handler) {
	c.successor = h
}

type ConcreteHandlerUpper struct {
	successor Handler
}

func (c *ConcreteHandlerUpper) Request(request int) {
	if request >= 20 {
		fmt.Println("ConcreteHandlerUpper can handle it")
	} else if c.successor != nil {
		fmt.Println("ConcreteHandlerUpper cannot handle move")
		c.successor.Request(request)
	}
}

func (c *ConcreteHandlerUpper) SetSuccessor(h Handler) {
	c.successor = h
}

func ResMain() {
	h1 := &ConcreteHandlerLower{}
	h2 := &ConcreteHandlerMiddle{}
	h3 := &ConcreteHandlerUpper{}

	h1.SetSuccessor(h2)
	h2.SetSuccessor(h3)

	req := []int{1, 10, 15, 20, 25}

	for _, r := range req {
		h1.Request(r)
	}
}
