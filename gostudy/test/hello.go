package test

import (
	"fmt"
)

func TestHello() {

	weeks := 97 / 7
	days := 97 % 7
	fmt.Println("weeks:", weeks, "days:", days)

	fmt.Printf("%d weeds and %d days", weeks, days)

	fmt.Println()
	str := fmt.Sprintf("%d", weeks)
	fmt.Println("str:", str)

}
