package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func TestTime() {

	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("now:%v\n", now)
	dateStr := fmt.Sprintf("now is %v-%v-%v %v:%v:%v\n", year, int(month), day, hour, minute, second)
	fmt.Printf("dateStr:%v\n", dateStr)

	dateStr = now.Format("2006-01-02 15:04:05")
	dateStr = now.Format("20060102150405")
	fmt.Printf("dateStr=%s\n", dateStr)
	dateStr = now.Format("2006-01-02")
	fmt.Printf("dateStr=%s\n", dateStr)
	dateStr = now.Format("15:04:05")
	fmt.Printf("dateStr=%s\n", dateStr)

	dateStr = now.Format("01")
	fmt.Printf("dateStr=%s\n", dateStr)
	dateStr = now.Format("02")
	fmt.Printf("dateStr=%s\n", dateStr)
	dateStr = now.Format("15")
	fmt.Printf("dateStr=%s\n", dateStr)

	i := 0
	for {
		i++
		fmt.Printf("%d\n", i)
		time.Sleep(time.Millisecond * 100)
		if i == 10 {
			break
		}
	}

	fmt.Printf("秒:%v,纳秒:%v\n", now.Unix(), now.UnixNano())

	i = 0
	for true {
		i++
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Printf("第%d次随机数是%d\n", i, n)
		time.Sleep(time.Nanosecond * 1)
		if n == 100 {
			break
		}
	}

	t1 := time.Date(2019, 8, 25, 23, 42, 59, 0, time.Local)
	fmt.Printf("%v\n", t1.Format("2006-01-02 15:04:05"))

	fmt.Printf("%v\n", t1.Add(time.Hour*24).Format("2006-01-02 15:04:05"))

	friday := time.Friday
	fmt.Printf("friday value is %v,type is %T\n", friday, friday)

	love := "i love"
	start := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		love += "" + strconv.Itoa(i)
	}
	end := time.Now().Unix()
	fmt.Printf("start:%v\n", start)
	fmt.Printf("end:%v\n", end)
	fmt.Printf("花费的时间是:%v\n", end-start)

}
