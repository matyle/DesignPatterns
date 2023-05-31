package ch9

type Answer interface {
	Answer1() string
	Answer2() string
}

type TestPaper struct {
	ans Answer
}

func (this *TestPaper) TestQuestion1() {
	println("TestQuestion1")
	println(this.ans.Answer1())
}

func (this *TestPaper) TestQuestion2() {
	println("TestQuestion2")
	println(this.ans.Answer2())
}

func (this *TestPaper) Answer1() string {
	return ""
}

func (this *TestPaper) Answer2() string {
	return ""
}

type TestPaperA struct {
	TestPaper
}

func (this *TestPaperA) Answer1() string {
	return "1 A"
}

func (this *TestPaperA) Answer2() string {
	return "2 B"
}

type TestPaperB struct {
	TestPaper
}

func (this *TestPaperB) Answer1() string {
	return "1 A"
}

func (this *TestPaperB) Answer2() string {
	return "2 B"
}

func PaperMain() {
	testPaperA := &TestPaperA{}
	testPaperA.ans = testPaperA
	testPaperA.TestQuestion1()
	testPaperA.TestQuestion2()
}
