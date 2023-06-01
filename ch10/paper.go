package ch9

// immatation of template method
// 模拟多态，go中没有继承，只能通过接口实现多态
type PaperTest interface {
	TestQuestion1()
}
type PaperAnswer interface {
	Answer1() string
	Answer2() string
}

type TestPaper struct {
}

func (this *TestPaper) TestQuestion1() {
	println("杨过得到，后来给了郭靖，炼成倚天剑、屠龙刀的玄铁可能是[ ] a.球磨铸铁 b. 马口铁 c.高速合金钢 d.碳素纤维 ")
}

func (this *TestPaper) Answer1() string {
	return ""
}

func (this *TestPaper) Answer2() string {
	return ""
}

type TestPaperA struct {
	base TestPaper
}

func (this *TestPaperA) TestQuestion1() {
	this.base.TestQuestion1()
	println(this.Answer1())
}

func (this *TestPaperA) Answer1() string {
	return "A"
}

type TestPaperB struct {
	base TestPaper
}

func (this *TestPaperB) TestQuestion1() {
	this.base.TestQuestion1()
	println(this.Answer1())
}

func (this *TestPaperB) Answer1() string {
	return "B"
}

func PaperMain() {
	var p PaperTest = &TestPaperA{}
	p.TestQuestion1()

	p = &TestPaperB{}
	p.TestQuestion1()
}
