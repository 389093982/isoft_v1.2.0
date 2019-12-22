package test

import "fmt"

func TestMap() {

	nameMap := make(map[string]string)
	nameMap["name"] = "zhoupeng"
	nameMap["age"] = "27"
	nameMap["address"] = "安徽舒城"
	//delete(nameMap,"age")
	nameMap["name"] = "dier"
	value, flag := nameMap["name"]
	if flag {
		fmt.Printf("有值:%v\n", value)
	} else {
		fmt.Printf("没有值")
	}

	for key, value := range nameMap {
		fmt.Printf("key:%v,value:%v\n", key, value)
	}
	fmt.Printf("nameMap:%v\n", nameMap)

}
