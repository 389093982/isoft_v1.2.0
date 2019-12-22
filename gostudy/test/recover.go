package test

import (
	"errors"
	"fmt"
)

func test() {
	defer func() {
		error := recover()
		if error != nil {
			fmt.Printf("遇到错误了:%v\n", error)
			fmt.Printf("发送邮件中...\n")
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Printf("除以0的结果:%v/n", result)
}

func readConfig(configName string) (err error) {
	if configName != "config.ini" {
		return errors.New("读取配置文件出错...")
	} else {
		return nil
	}
}

func test2() {
	err := readConfig("config.ini")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("test2后面的代码逻辑...")
	}

}

func TestRecover() {
	test2()
	fmt.Printf("main后面的逻辑代码...\n")
}
