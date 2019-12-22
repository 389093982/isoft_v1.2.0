package test

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func TestRdeaWriter() {
	filePath1 := "C:\\Users\\zhoupeng\\Desktop\\01.txt"
	filePath2 := "C:\\Users\\zhoupeng\\Desktop\\02.txt"
	reader, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Printf("read error")
		return
	}
	err = ioutil.WriteFile(filePath2, reader, 0666)
	if err != nil {
		fmt.Printf("write error")
	}
	fmt.Printf("write ok")
}

func TestWriterFile() {
	filePath := "C:\\Users\\zhoupeng\\Desktop\\写入文件.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("created or write error !")
		return
	}
	defer file.Close()
	//str := "写入文件，用于测试\n"
	str := "追加内容test\n"
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString(str)
	}
	write.Flush()

}

func TestReadFile() {

	file, err := os.Open("C:\\Users\\zhoupeng\\Desktop\\测试文件.txt")
	if err != nil {
		fmt.Printf("file open error %v", err)
		return
	}
	defer file.Close()

	fmt.Printf("content:%v\n", file)

	reader := bufio.NewReader(file)
	for true {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("read end!\n")
			break
		}
		fmt.Printf("str:%v", line)

	}

	content, err := ioutil.ReadFile("C:\\Users\\zhoupeng\\Desktop\\测试文件.txt")
	if err != nil {
		fmt.Printf("read error!")
	}
	fmt.Printf("content:%v", string(content))

}
