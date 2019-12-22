package test

import "fmt"

func maopao(arr *[5]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i; j++ {
			if (*arr)[i] < (*arr)[j+i] {
				temp := (*arr)[j+i]
				(*arr)[j+i] = (*arr)[i]
				(*arr)[i] = temp
			}
		}
	}
}

func fbn(n int) []uint {
	fbnSlice := make([]uint, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func TestSlice() {

	intArray := [4]int{2, 3, 4, 5}
	intSlice := intArray[1:]
	fmt.Printf("intSlice:%v\n", intSlice)
	intSlice[0] = 1
	fmt.Printf("intArray:%v\n", intArray)
	fmt.Printf("intSlice:%v\n", intSlice)

	strSlice := []int{1, 2, 3, 4}
	fmt.Printf("strSlice:%v\n", strSlice)

	makeSlice := make([]int, 3)
	copy(makeSlice, intSlice)
	makeSlice = append(makeSlice, 4, 5, 6)
	fmt.Printf("makeSlice:%v\n", makeSlice)

	str := "hello@guigu"
	fmt.Printf("str:%v\n", str)

	arr := []byte(str)
	arr[0] = 'Z'
	str = string(arr)
	fmt.Printf("str:%v\n", str)

	fmt.Printf("fbnSlice:%v\n", fbn(50))

	intArr := [5]int{3, 6, 2, 8, 1}
	maopao(&intArr)
	fmt.Printf("maopao:%v\n", intArr)
}
