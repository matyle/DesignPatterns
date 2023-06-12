package ch14

import "fmt"

type DelegateSubject interface {
	SetAction(action string)
	GetSubjectState() string
	Notify()
}

type EventHandle func()

// concrete subject
type Boss struct {
	update []EventHandle
	action string
}

func NewBoss() *Boss {
	return &Boss{
		update: make([]EventHandle, 0, 10),
	}
}

func (b *Boss) SetAction(action string) {
	b.action = action
}

func (b *Boss) GetSubjectState() string {
	return b.action
}

func (b *Boss) AddEventHandle(handle EventHandle) {
	b.update = append(b.update, handle)
}

func (b *Boss) Notify() {
	for _, handle := range b.update {
		handle()
	}
}

// concrete observer
type DeleStockObserver struct {
	name string
	sub  DelegateSubject
}

func NewDeleStockObserver(name string, sub DelegateSubject) *DeleStockObserver {
	return &DeleStockObserver{
		name: name,
		sub:  sub,
	}
}

func (s *DeleStockObserver) CloseStockMarket() {
	fmt.Printf("%s %s 关闭股票行情，继续工作！\n", s.sub.GetSubjectState(), s.name)
}

func DelegateMain() {
	sub := NewBoss()
	sub.SetAction("我胡汉三回来了！")
	observer1 := NewDeleStockObserver("张三", sub)
	observer2 := NewDeleStockObserver("李四", sub)
	sub.AddEventHandle(observer1.CloseStockMarket)
	sub.AddEventHandle(observer2.CloseStockMarket)
	sub.Notify()
}
