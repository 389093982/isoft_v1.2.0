package test

import "fmt"

type MethodUtils struct {
}

func (mu MethodUtils) printJiuJiu() {
	fmt.Print("请输入一个数：")
	var hight int
	fmt.Scanln(&hight)
	for i := 0; i < hight; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%v * %v = %v  ", j+1, i+1, (j+1)*(i+1))
		}
		fmt.Println()
	}
}

func TestStruct() {
	var mu MethodUtils
	mu.printJiuJiu()

	var m = MethodUtils{}
	m.printJiuJiu()

	m2 := MethodUtils{}
	m2.printJiuJiu()
}
