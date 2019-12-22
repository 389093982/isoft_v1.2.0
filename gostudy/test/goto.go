package test

import "fmt"

func TestGoto() {

	fmt.Printf("1")
	fmt.Printf("2")
	goto label1
	fmt.Printf("3")
	fmt.Printf("4")
	fmt.Printf("5")
label1:
	fmt.Printf("6")

}
