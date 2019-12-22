package test

import "fmt"

func TestSwitch() {

	key := 'a'
	fmt.Print("请输入一个字符（a,b,c,d,e,f,g）:")
	fmt.Scanf("%c", &key)
	switch key {
	case 'a', 'b':
		fmt.Println("周一")
	case 'c':
		fmt.Println("周二")
	case 'd':
		fmt.Println("周三")
	default:
		fmt.Println("无聊...")
	}

}
