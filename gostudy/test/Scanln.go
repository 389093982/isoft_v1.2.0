package test

import "fmt"

func TestScanln() {

	var name string
	var age int
	var salary float64
	var isPass bool

	//fmt.Print("请输入您的姓名:")
	//fmt.Scanln(&name)
	//fmt.Print("请输入您的年龄:")
	//fmt.Scanln(&age)
	//fmt.Print("请输入您的薪水:")
	//fmt.Scanln(&salary)
	//fmt.Print("请输入是否通过考试:")
	//fmt.Scanln(&isPass)

	fmt.Print("请输入姓名 年龄 薪水 是否通过考试（按空格分离）:")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)

	fmt.Printf("姓名:%v, 年龄:%v, 薪水:%v, 是否通过考试:%v", name, age, salary, isPass)

}
