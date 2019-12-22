package test

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int)
	lock  sync.Mutex
)

type Cat struct {
	Name string
	Age  int
}

func TestChan() {

	var catChan chan interface{}
	catChan = make(chan interface{}, 10)
	catChan <- Cat{Name: "catName1", Age: 18}
	catChan <- Cat{Name: "catName2", Age: 22}
	close(catChan)
	cat1 := <-catChan
	cat2 := <-catChan
	cat3, ok := <-catChan
	fmt.Printf("cat1:%T,%v, - cat2:%T,%v\n", cat1, cat1, cat2, cat2)
	if !ok {
		fmt.Printf("false\n")
	}
	fmt.Printf("cat3:%v\n", cat3)
	//for v:=range catChan{
	//	fmt.Printf("value:%v\n",v)
	//}

	var intChan chan int
	intChan = make(chan int, 22)
	intChan <- 2
	intChan <- 3
	intChan <- -5
	fmt.Printf("intChan len:%v cap:%v\n", len(intChan), cap(intChan))

	n1 := <-intChan
	n2 := <-intChan
	n3 := <-intChan
	fmt.Printf("n1=%v,n2=%v,n3=%v\n", n1, n2, n3)

}

func getSum(n int) {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	lock.Lock()
	myMap[n] = sum
	lock.Unlock()
}

func TestGoroutine() {

	for i := 1; i <= 30; i++ {
		go getSum(i)
	}

	time.Sleep(time.Second * 5)
	for index, value := range myMap {
		fmt.Printf("myMap[%v]=%v\n", index, value)
	}
}
