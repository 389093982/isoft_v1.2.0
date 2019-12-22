package test

import "fmt"

type person struct {
	Name string
	age  int
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age >= 0 && age <= 150 {
		p.age = age
	} else {
		fmt.Printf("age设置不合理\n")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func TestPerson() {

	p1 := NewPerson("tom")
	p1.age = 14
	fmt.Println(p1)

}
