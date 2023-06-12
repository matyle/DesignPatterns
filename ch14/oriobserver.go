package ch14

import "fmt"

// Secretary watch boss, and notify all observers
// 双向耦合的代码
type OriSecretary struct {
	observers []*OriObserver
	action    string
}

func (sec *OriSecretary) Notify() {
	for _, ob := range sec.observers {
		ob.Update()
	}
}

func (sec *OriSecretary) Add(ob *OriObserver) {
	sec.observers = append(sec.observers, ob)
}

func (sec *OriSecretary) GetAction() string {
	return sec.action
}
func (sec *OriSecretary) SetAction(a string) {
	sec.action = a
}

type OriObserver struct {
	name string
	sub  *OriSecretary
}

func (ob *OriObserver) Update() {
	fmt.Printf("Boss is here,%s, %s you need go back work\n", ob.sub.GetAction(), ob.name)
}

func ObserverV1() {
	secretary := &OriSecretary{
		observers: make([]*OriObserver, 0, 10),
	}
	secretary.SetAction("Stop see the market quotes ")
	ob1 := &OriObserver{
		name: "Jack",
		sub:  secretary,
	}
	ob2 := &OriObserver{
		name: "Mary",
		sub:  secretary,
	}
	secretary.Add(ob1)
	secretary.Add(ob2)
	secretary.Notify()
}
