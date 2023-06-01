package ch13

import "fmt"

// Builder abstract interface
type Builder interface {
	BuildHeader()
	BuildBody()
	BuildFooter()
}

// ConcreteBuilder
type TextBuilder struct {
}

func (t *TextBuilder) BuildHeader() {
	fmt.Println("BuildHeader")
}

func (t *TextBuilder) BuildBody() {
	fmt.Println("BuildBody")
}

func (t *TextBuilder) BuildFooter() {
	fmt.Println("BuildFooter")
}

func NewTextBuilder() *TextBuilder {
	return &TextBuilder{}
}

// ConcreteBuilder
type PicBuilder struct {
}

func (p *PicBuilder) BuildHeader() {
	fmt.Println("PicBuilder BuildHeader")
}

func (p *PicBuilder) BuildBody() {
	fmt.Println("PicBuilder BuildBody")
}

func (p *PicBuilder) BuildFooter() {
	fmt.Println("PicBuilder BuildFooter")
}

func NewPicBuilder() *PicBuilder {
	return &PicBuilder{}
}

// Director
type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.BuildHeader()
	d.builder.BuildBody()
	d.builder.BuildFooter()
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

// Client
func BuilderTest() {
	textBuilder := NewTextBuilder()
	director := NewDirector(textBuilder)
	director.Construct()

	picBuilder := NewPicBuilder()
	director = NewDirector(picBuilder)
	director.Construct()
}
