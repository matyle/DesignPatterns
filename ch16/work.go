package ch16

import "fmt"

type StateV2 interface {
	WriteProgram(w *Work)
}

type MorningState struct {
}

func NewMorningState() *MorningState {
	return &MorningState{}
}

func (m *MorningState) WriteProgram(w *Work) {
	if w.hour < 12 {
		fmt.Printf("当前时间%d点，上午工作，精神百倍\n", w.hour)
	} else {
		w.SetState(NewAfterNoonState())
		w.WriteProgram()
	}
}

type AfterNoonState struct {
}

func NewAfterNoonState() *AfterNoonState {
	return &AfterNoonState{}
}

func (m *AfterNoonState) WriteProgram(w *Work) {
	if w.hour < 17 {
		fmt.Printf("当前时间%d点，写代码，下午继续努力\n", w.hour)
	} else {
		w.SetState(NewEveningState())
		w.WriteProgram()
	}
}

type EveningState struct {
}

func NewEveningState() *EveningState {
	return &EveningState{}
}
func (m *EveningState) WriteProgram(w *Work) {
	if w.IsFinish() {
		fmt.Println(" 任务已完成，休息")
	}
	if w.hour < 21 {
		fmt.Printf("当前时间%d点，继续加班\n", w.hour)
	} else {
		// w.SetState(NewAfterNoonState())
		// w.SetState(NewMorningState())
		fmt.Println("睡觉")
	}
}

type Work struct {
	currentState StateV2
	hour         int
	finish       bool
}

func NewWork() *Work {
	return &Work{
		currentState: NewMorningState(),
		hour:         10,
		finish:       false,
	}
}

func (w *Work) SetState(s StateV2) {
	w.currentState = s
}

func (w *Work) IsFinish() bool {
	return w.finish
}

func (w *Work) WriteProgram() {
	w.currentState.WriteProgram(w)
}

func WorkMain() {
	w := NewWork()
	w.WriteProgram()
	w.hour = 11
	w.WriteProgram()
	w.hour = 14
	w.WriteProgram()
	w.hour = 18
	w.WriteProgram()
}
