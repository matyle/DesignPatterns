package ch12

import "fmt"

// 外观模式

type SubSystem1 struct {
}

func (s *SubSystem1) Method1() {
	fmt.Println("SubSystem1 Method1")
}

type SubSystem2 struct {
}

func (s *SubSystem2) Method2() {
	fmt.Println("SubSystem2 Method2")
}

type SubSystem3 struct {
}

func (s *SubSystem3) Method3() {
	fmt.Println("SubSystem3 Method3")
}

type SubSystem4 struct {
}

func (s *SubSystem4) Method4() {
	fmt.Println("SubSystem4 Method4")
}

type Facade struct {
	s1 *SubSystem1
	s2 *SubSystem2
	s3 *SubSystem3
	s4 *SubSystem4
}

func NewFacade() *Facade {
	return &Facade{
		s1: &SubSystem1{},
		s2: &SubSystem2{},
		s3: &SubSystem3{},
		s4: &SubSystem4{},
	}
}

func (f *Facade) MethodA() {
	fmt.Println("MethodA")
	f.s1.Method1()
	f.s2.Method2()
	f.s4.Method4()
}

func (f *Facade) MethodB() {
	fmt.Println("MethodB")
	f.s3.Method3()
	f.s4.Method4()
}

func FacadeTest() {
	f := NewFacade()
	f.MethodA()
	f.MethodB()
}
