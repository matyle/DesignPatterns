package ch6

import "fmt"

// Component
type ComponentPerson interface {
	Show()
}

type Person struct {
	name string
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) Show() {
	fmt.Println("装扮的", p.name)
}

// Decorator
type Finery struct {
	component ComponentPerson
	Person
}

func (f *Finery) Decorate(component ComponentPerson) {
	f.component = component
}

func (f *Finery) Show() {
	f.component.Show()
}

// ConcreteDecorator
type TShirts struct {
	Finery
}

func (t *TShirts) Show() {
	fmt.Print("大T恤 ")
	t.Finery.Show()
}

// ConcreteDecorator
type BigTrouser struct {
	Finery
}

func (b *BigTrouser) Show() {
	fmt.Print("垮裤 ")
	b.Finery.Show()
}

// ConcreteDecorator
type Sneakers struct {
	Finery
}

func (s *Sneakers) Show() {
	fmt.Print("破球鞋 ")
	s.Finery.Show()
}

// client
func DecoratorMain() {
	person := &Person{name: "小菜"}
	fmt.Println("第一种装扮")
	sneaker := &Sneakers{}
	bt := &BigTrouser{}
	ts := &TShirts{}
	sneaker.Decorate(person)
	bt.Decorate(sneaker)
	ts.Decorate(bt)
	ts.Show()
}
