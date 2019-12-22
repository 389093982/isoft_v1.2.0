package test

import "fmt"

func TestIf() {

	age := 0
	fmt.Print("请输入你的年龄:")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Printf("age:%d", age)
	} else {
		fmt.Printf("小于18")
	}

	fmt.Println()
	if i := age + 1; i > 18 {
		fmt.Printf("你好哦，你大于18岁咯")
	} else {
		fmt.Printf("你仍然是小于18岁")
	}

}
