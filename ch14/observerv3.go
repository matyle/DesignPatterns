package ch14

type V3Subject interface {
	Notify()
	Attach(V3Observer)
	GetAction() string
	SetAction(string)
	// Detach(Observer)
}

type V3Observer interface {
	Update()
}

type V3SubSecretary struct {
	observers []V3Observer
	action    string
}

func (s *V3SubSecretary) Notify() {
	for _, o := range s.observers {
		o.Update()
	}
}

func (s *V3SubSecretary) Attach(o V3Observer) {
	s.observers = append(s.observers, o)
}

func (s *V3SubSecretary) SetAction(action string) {
	s.action = action
}

func (s *V3SubSecretary) GetAction() string {
	return s.action
}

type V3StockObserver struct {
	name string
	sub  V3Subject
}

func NewV3StockObserver(name string, sub V3Subject) *V3StockObserver {
	return &V3StockObserver{name: name, sub: sub}
}

func (o *V3StockObserver) Update() {
	println(o.name, "close stock, continue work", o.sub.GetAction())
}

type V3NBAObserver struct {
	name string
	sub  V3Subject
}

func NewV3NBAObserver(name string, sub V3Subject) *V3NBAObserver {
	return &V3NBAObserver{name: name, sub: sub}
}

func (o *V3NBAObserver) Update() {
	println(o.name, "close NBA, continue work", o.sub.GetAction())
}

func ObserverV3() {
	var sub V3Subject
	sub = &V3SubSecretary{}
	stockObserver := NewV3StockObserver("stock", sub)
	nbaObserver := NewV3NBAObserver("nba", sub)

	sub.Attach(stockObserver)
	sub.Attach(nbaObserver)

	sub.SetAction("boss is coming")
	sub.Notify()

}
