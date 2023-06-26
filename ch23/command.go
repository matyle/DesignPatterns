package ch23

import "fmt"

type Command interface {
	Execute()
}

type Receiver interface {
	Action()
}

type Invoker struct {
	command []Command
}

type ConcreteCommand struct {
	receiver Receiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

func (i *Invoker) SetCommand(c Command) {
	i.command = append(i.command, c)
}

func (i *Invoker) RemoveTailCommand() {
	if len(i.command) > 0 {
		i.command = i.command[:len(i.command)-1]
	}
}

func (i *Invoker) ExecuteCommand() {
	for _, c := range i.command {
		c.Execute()
	}
}

type ReceiverA struct {
}

func (r *ReceiverA) Action() {
	println("receiverA action")
}

type ReceiverB struct {
}

func (r *ReceiverB) Action() {
	println("receiverB action")
}

func CommandMain() {
	invoker := new(Invoker)
	receiverA := new(ReceiverA)
	receiverB := new(ReceiverB)
	commandA := &ConcreteCommand{receiverA}
	commandB := &ConcreteCommand{receiverB}
	invoker.SetCommand(commandA)
	invoker.SetCommand(commandB)
	invoker.ExecuteCommand()
	invoker.RemoveTailCommand()
	fmt.Println("remove tail command")
	invoker.ExecuteCommand()
}
