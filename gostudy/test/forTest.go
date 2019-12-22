package test

import (
	"fmt"
	"math/rand"
	"time"
)

func TestFor() {

	for i := 0; i <= 10; i++ {
		fmt.Println("xxxx")
	}

	str := "aaa"
	for index, value := range str {
		fmt.Printf("index:%d,value:%c\n", index, value)
	}

	m := 0
	for {
		if m > 10 {
			break
		}
		fmt.Printf("hello-%d\n", m)
		m++

	}

	//count := 0
	//for{
	//	rand.Seed(time.Now().UnixNano())
	//	i := rand.Intn(100) + 1
	//	count ++
	//	fmt.Printf("count:%d,i=%d\n",count,i)
	//	if i==99{
	//		fmt.Printf("第%d次生成了99\n",count)
	//		break
	//	}
	//}

	fmt.Printf("结束咯\n")

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(999999999)
	strRandom := "000000000" + fmt.Sprintf("%d", random)
	fmt.Printf("strRandom:%v\n", strRandom)

	strRandom = strRandom[len(strRandom)-9 : len(strRandom)]
	fmt.Printf("strRandom:%v\n", strRandom)

}
