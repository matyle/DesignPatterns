package ch19

import "fmt"

type Company interface {
	Add(c Company)
	Remove()
	Display()
	DoAction()
}
type CompanyBase struct {
	name string
}

// Leaf 结点
type FinanceDepartment struct {
	CompanyBase
	duty string
}

func (c *FinanceDepartment) Add(com Company) {
}

func (c *FinanceDepartment) Remove() {

}

func (c *FinanceDepartment) Display() {
	fmt.Println("this is FinanceDepartment")
}

func (c *FinanceDepartment) DoAction() {
	fmt.Printf("负责%s财务相关", c.name)
}

type HRDepartment struct {
	CompanyBase
	duty string
}

func (c *HRDepartment) Add(com Company) {

}

func (c *HRDepartment) Remove() {

}
func (c *HRDepartment) Display() {
	fmt.Println("this is HRDepartment")
}

func (c *HRDepartment) DoAction() {
	fmt.Printf("负责%s招聘相关", c.name)
}

type ConcreteCompany struct {
	CompanyBase
	children []Company
}
