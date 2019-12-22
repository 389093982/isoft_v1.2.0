package test

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

type studentSlice []Student

func (this studentSlice) Len() int {
	return len(this)
}

func (stu studentSlice) Less(i, j int) bool {
	return stu[i].Score > stu[j].Score
}

func (stu studentSlice) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}

func TestSort() {

	var stuSlice studentSlice
	for i := 0; i < 10; i++ {
		student := Student{
			Name:  fmt.Sprintf("name-%d", i+1),
			Age:   rand.Intn(70),
			Score: float64(rand.Intn(100)),
		}
		stuSlice = append(stuSlice, student)
	}

	fmt.Println("排序前:")
	for _, v := range stuSlice {
		fmt.Println(v)
	}

	fmt.Println("排序后:")
	sort.Sort(stuSlice)
	for _, v := range stuSlice {
		fmt.Println(v)
	}

}
