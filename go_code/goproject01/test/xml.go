package test

import (
	"fmt"
	"encoding/xml"
)

type Dog struct {
	Like string `xml:"like"`
}

type Person struct {
	XMLName xml.Name `xml:"xml"`
	Age int64 `xml:"age"`
	DogS []Dog `xml:"dogs"`
}

func TestXml() {
	d := Dog{}
	d.Like = "like"
	dogSlice := make([]Dog,1)
	dogSlice[0] = d
	p := Person{}
	p.Age = 28
	p.DogS = dogSlice


	xmlstr, err := xml.Marshal(p)
	if err!=nil {
		fmt.Printf("trans error:%v\n",err.Error())
		return
	}
	fmt.Printf("data:%v\n",string(xmlstr))


}
