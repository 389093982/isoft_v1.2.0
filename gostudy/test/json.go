package test

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"monster_name"`
	Age      int64  `json:"monster_age"`
	Birthday string
	Skill    string
}

func TestUnMarshal() {
	var m Monster
	str := "{\"monster_name\":\"白骨精\",\"monster_age\":18,\"Birthday\":\"20191117\",\"Skill\":\"MLOVE\"}"
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Printf("trans error")
		return
	}
	fmt.Printf("m:%v", m)

}

func TestJson() {

	m := Monster{
		Name:     "白骨精",
		Age:      18,
		Birthday: "20191117",
		Skill:    "MLOVE",
	}
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Printf("trans error")
		return
	}
	fmt.Printf("monster data:%v\n", string(data))

	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "zhoupeng"
	a["age"] = 28
	data, err = json.Marshal(a)
	if err != nil {
		fmt.Printf("trans error")
		return
	}
	fmt.Printf("data:%v\n", string(data))

	var slice []map[string]interface{}

	var b map[string]interface{}
	b = make(map[string]interface{})
	b["name"] = "xiaomei"
	b["age"] = 18

	slice = append(slice, a)
	slice = append(slice, b)
	data, err = json.Marshal(slice)
	if err != nil {
		fmt.Printf("trans error")
		return
	}
	fmt.Printf("data:%v\n", string(data))

	var d int = 234
	data, err = json.Marshal(d)
	if err != nil {
		fmt.Printf("trans error")
		return
	}
	fmt.Printf("data:%v\n", string(data))

}
