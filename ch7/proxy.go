package ch7

import "fmt"

//代理模式
// Jason(Pursuit) use  DaiLi(proxy)  to pursuit JiaoJiao

type GiveGifter interface {
	GiveDolls()
	GiveFlowers()
	GiveChocolate()
}

type SchoolGirl struct {
	name string
}

func (s *SchoolGirl) SetName(name string) {
	s.name = name
}

type Pursuit struct {
	sc *SchoolGirl
}

func (p *Pursuit) New(sc *SchoolGirl) {
	p.sc = sc
}

func (p *Pursuit) GiveDolls() {
	fmt.Println("送洋娃娃给", p.sc.name)
}

func (p *Pursuit) GiveFlowers() {
	fmt.Println("送玫瑰花给", p.sc.name)
}

func (p *Pursuit) GiveChocolate() {
	fmt.Println("送巧克力给", p.sc.name)
}

type Proxy struct {
	pursuit *Pursuit
}

func (pr *Proxy) New(pursuit *Pursuit) {
	pr.pursuit = pursuit
}

func (p *Proxy) GiveDolls() {
	p.pursuit.GiveDolls()
}

func (p *Proxy) GiveFlowers() {
	p.pursuit.GiveFlowers()
}

func (p *Proxy) GiveChocolate() {
	p.pursuit.GiveChocolate()
}

func ProxyRun() {
	jiaojiao := new(SchoolGirl)
	jiaojiao.SetName("娇娇")
	pursuit := new(Pursuit)
	pursuit.New(jiaojiao)

	proxy := new(Proxy)
	proxy.New(pursuit)

	proxy.GiveDolls()
	proxy.GiveFlowers()
	proxy.GiveChocolate()

}
