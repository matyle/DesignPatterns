package ch18

import "fmt"

type Memoer interface {
}
type Roler interface {
}

type BaseRole struct {
	vlt int
	atk int
	def int
}

type RoleStateMemento struct {
	base *BaseRole
}

func NewRoleStateMemento(vlt, atk, def int) (r *RoleStateMemento) {
	r = &RoleStateMemento{
		base: &BaseRole{
			vlt: vlt,
			atk: atk,
			def: def,
		},
	}
	return r
}

func (r *RoleStateMemento) Vitality() int {
	return r.base.vlt
}

func (r *RoleStateMemento) Attack() int {
	return r.base.atk
}

func (r *RoleStateMemento) Defence() int {
	return r.base.def
}

type GameRole struct {
	base *BaseRole
}

func NewGameRole(vlt, atk, def int) *GameRole {
	return &GameRole{
		base: &BaseRole{
			vlt: vlt,
			atk: atk,
			def: def,
		},
	}
}

func (j *GameRole) SaveState() *RoleStateMemento {
	return NewRoleStateMemento(j.base.vlt, j.base.atk, j.base.def)
}

func (j *GameRole) RecoverState(memento *RoleStateMemento) {
	j.base.vlt = memento.Vitality()
	j.base.atk = memento.Attack()
	j.base.def = memento.Defence()
}

func (j *GameRole) StateDisplay() {
	fmt.Println("角色当前状态：")
	fmt.Println("体力：", j.base.vlt)
	fmt.Println("攻击力：", j.base.atk)
	fmt.Println("防御力：", j.base.def)
}

func (j *GameRole) InitState() {
	j.base.vlt = 100
	j.base.atk = 100
	j.base.def = 100
}

func (j *GameRole) Fight() {
	j.base.vlt = 0
	j.base.atk = 0
	j.base.def = 0
}

// adminstate

type RoleStateCaretaker struct {
	memento *RoleStateMemento
}

func NewRoleStateCaretaker() *RoleStateCaretaker {
	return &RoleStateCaretaker{
		memento: &RoleStateMemento{},
	}
}

func (r *RoleStateCaretaker) Memento() *RoleStateMemento {
	return r.memento
}

func (r *RoleStateCaretaker) SetMemento(memento *RoleStateMemento) {
	r.memento = memento
}

func MementoMain() {
	// 大战boss前
	j := NewGameRole(100, 100, 100)
	j.StateDisplay()

	// 保存进度
	stateAdmin := NewRoleStateCaretaker()
	stateAdmin.SetMemento(j.SaveState())

	// 大战boss时，损耗严重
	j.Fight()
	j.StateDisplay()

	// 恢复之前状态
	j.RecoverState(stateAdmin.Memento())
	j.StateDisplay()
}
