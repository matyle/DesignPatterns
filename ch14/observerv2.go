package ch14

// observer as a obstract or interface
type V2Secretary struct {
	observers []Observer
	action    string
}

type Observer interface {
	Update()
}

type BaseObserver struct {
	name string
	sub  *V2Secretary
}

type StockObserver struct {
	BaseObserver
}

type NBAObserver struct {
	BaseObserver
}

func (s *V2Secretary) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *V2Secretary) Notify() {
	for _, o := range s.observers {
		o.Update()
	}
}

func (s *V2Secretary) SetAction(action string) {
	s.action = action
}

func (o *StockObserver) Update() {
	println(o.name, o.sub.action, "notified to close stock")
}

func (o *NBAObserver) Update() {
	println(o.name, o.sub.action, "notified to close NBA")
}

func NewV2Secretary() *V2Secretary {
	return &V2Secretary{}
}

func NewStockObserver(name string, sub *V2Secretary) *StockObserver {
	return &StockObserver{BaseObserver{name, sub}}
}

func NewNBAObserver(name string, sub *V2Secretary) *NBAObserver {
	return &NBAObserver{BaseObserver{name, sub}}
}

func V2ObserverMain() {
	sub := NewV2Secretary()
	sub.Attach(NewStockObserver("JackMa", sub))
	sub.Attach(NewNBAObserver("PonyMa", sub))
	sub.SetAction("Boss is coming")
	sub.Notify()
}
