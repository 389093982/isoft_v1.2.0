package test

import (
	"fmt"
	"go_code/goproject01/utils"
)

func TestFunc() {

	n1 := 20.0
	n2 := 3.0
	var operator byte = 'a'
	isOk, result := utils.GetOpratorResult(n1, n2, operator)
	fmt.Printf("isOk:%s,result:%f\n", isOk, result)

	utils.GetResult(0)

}
