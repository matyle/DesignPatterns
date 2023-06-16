package ch16

import (
	"fmt"
)

// 不同状态
type State interface {
	String() string
	Handel(c *Context)
}

type Context struct {
	state State
}

func NewContext(s State) *Context {
	return &Context{
		state: s,
	}
}

func (c *Context) Request() {
	c.state.Handel(c)
}

func (c *Context) String() string {
	return c.state.String()
}

type BaseState struct {
	v string
}

func (e *BaseState) Get() string {
	return e.v
}

// 具体状态
type EnergyState struct {
	v    string
	base *BaseState
}

func NewEnergeSate() *EnergyState {
	// t := time.Date(0, 0, 0, 9, 0, 0, 0, time.Local).Format("15:04:05")
	return &EnergyState{
		base: &BaseState{
			v: "精力充沛",
		},
	}
}

func (e *EnergyState) String() string {
	return e.base.v
}
func (e *EnergyState) Handel(c *Context) {
	fmt.Printf("state is %s ", c.state.String())
	c.state = NewRestSate()
	fmt.Printf("nextstate is %s\n", c.state.String())
}

type RestState struct {
	base *BaseState
}

func NewRestSate() *RestState {
	// t := time.Date(0, 0, 0, 9, 0, 0, 0, time.Local).Format("15:04:05")
	return &RestState{
		base: &BaseState{
			v: "休息状态",
		},
	}
}

func (e *RestState) String() string {
	return e.base.v
}
func (e *RestState) Handel(c *Context) {
	//逻辑判断是否转换到下一状态 或者加入参数到下一个状态
	fmt.Printf("state is %s ", c.state.String())
	c.state = NewEnergeSate()
	fmt.Printf("nextstate is %s\n", c.state.String())
}
func stateMain() {
	context := NewContext(NewEnergeSate())
	context.Request()
}
