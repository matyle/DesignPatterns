package ch2

type Strategy interface {
	AlgorithmInterface()
}

type ConcreteStrategyA struct {
}

func (this *ConcreteStrategyA) AlgorithmInterface() {
	println("this is ConcreteStrategyA")
}

type ConcreteStrategyB struct {
}

func (this *ConcreteStrategyB) AlgorithmInterface() {
	println("this is ConcreteStrategyB")
}

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	context := new(Context)
	context.strategy = strategy
	return context
}

func (this *Context) ContextInterface() {
	this.strategy.AlgorithmInterface()
}

func StrategyRun() {
	sa := &ConcreteStrategyA{}
	context := NewContext(sa)
	context.ContextInterface()

	sb := &ConcreteStrategyB{}
	context = NewContext(sb)
	context.ContextInterface()
}
