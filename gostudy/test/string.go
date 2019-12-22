package test

import (
	"fmt"
	"strconv"
	"strings"
)

func TestString() {
	//1.len
	str := "hello中国"
	fmt.Printf("str的长度是:%d", len(str))

	//2.[]rune
	strRune := []rune(str)
	for i := 0; i < len(strRune); i++ {
		fmt.Printf("字符:%c\n", strRune[i])
	}

	//2.1 range
	for index, value := range str {
		fmt.Printf("字符[%d]是%c\n", index, value)
	}

	//3.Atoi
	strNum := "456"
	num, error := strconv.Atoi(strNum)
	if error == nil {
		fmt.Printf("转换成功,num=%d,type is %T\n", num, num)
	} else {
		fmt.Printf("转换异常:%v", error)
	}

	//4.Itoa
	age := 27
	strAge := strconv.Itoa(age)
	fmt.Printf("strAge:%v,type is %T\n", strAge, strAge)

	//10->2
	fmt.Printf("123对应的2进制%v\n", strconv.FormatInt(123, 2))

	//string.Contains
	fmt.Printf("hello,world 是否包含world:%v\n", strings.Contains("hello,world", "world"))

	//strings.Count
	fmt.Printf("hello中有%v个o\n", strings.Count("hello", "o"))

	//== strings.EqualFold
	fmt.Printf("判断是否相等:%t\n", "aaa" == "Aaa")
	fmt.Printf("判断是否相等:%t\n", strings.EqualFold("aaa", "Aaa"))

	//index of first
	index := strings.Index("hello_world_world_world", "world")
	fmt.Printf("index:%v\n", index)

	//index of last
	index = strings.LastIndex("1234567888", "8")
	fmt.Printf("最后一次出现的位置是:%d\n", index)

	//replace
	str = "i love go go go"
	fmt.Printf("go 替换为go语言:%v\n", strings.Replace(str, "go", "go语言", 2))
	fmt.Printf("go 替换为go语言:%v\n", strings.ReplaceAll(str, "go", "go语言"))
	fmt.Printf("str:%v\n", str)

	//split
	str = "i ,love, you"
	strArray := strings.Split(str, ",")
	for index, value := range strArray {
		fmt.Printf("位置%d,值:%v\n", index, value)
	}

	//大小写
	str = "I Love You"
	fmt.Printf("转化小写:%v\n", strings.ToLower(str))
	fmt.Printf("转化大写:%v\n", strings.ToUpper(str))

	//trim
	str = "    abc i love you ab    "
	fmt.Printf("去空格后:%v\n", strings.TrimSpace(str))
	fmt.Printf("去空格后:%v\n", strings.TrimLeft(str, "ba"))

	//开头和结尾
	fmt.Printf("是否:%t\n", strings.HasPrefix("http:aaaaaa", "http"))
	fmt.Printf("是否:%t\n", strings.HasSuffix("123.jpg", "jpg"))

}
