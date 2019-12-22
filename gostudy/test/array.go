package test

import "fmt"

func TestArray() {

	herrs := [3]string{"1", "2", "3"}
	fmt.Printf("herrs:%v\n", herrs)

	charter := [26]byte{}
	for index, _ := range charter {
		charter[index] = 'A' + byte(index)
	}

	for _, v := range charter {
		fmt.Printf("%c ", v)
	}

}
