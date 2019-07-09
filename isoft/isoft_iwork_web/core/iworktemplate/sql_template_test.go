package iworktemplate

import (
	"os"
	"testing"
	"text/template"
)

var m map[string]interface{}

func Test_Template(t *testing.T) {
	s := `{{ if eq .hello "abc" }}
			abc
		  {{ else if eq .hello "bcd" }}
			bcd
		  {{ else }}
			def
		  {{ end }}`
	//解析模板
	tmpl, err := template.New("test").Parse(s)
	if err != nil {
		panic(err)
	}
	//数据驱动模板
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"hello": "cde"})
	if err != nil {
		panic(err)
	}
}

func Test_Template2(t *testing.T) {
	s := `{{ if eq .hello "abc" }}
			{{ .abc }}
		  {{ else if eq .hello "bcd" }}
			{{ .bcd }}
		  {{ else }}
			{{ .def }}
		  {{ end }}`
	//解析模板
	tmpl, err := template.New("test").Parse(s)
	if err != nil {
		panic(err)
	}
	//数据驱动模板
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"hello": "cde", "abc": "abc", "bcd": "bcd", "def": "def"})
	if err != nil {
		panic(err)
	}
}

func Test_Template3(t *testing.T) {
	s := `{{ if eq .hello "abc" }}
			select * from abc where a = :a
		  {{ else if eq .hello "bcd" }}
			select * from bcd where b = :b
		  {{ else }}
			select * from cde where c = :c[nil,0]
		  {{ end }}`
	//解析模板
	tmpl, err := template.New("test").Parse(s)
	if err != nil {
		panic(err)
	}
	//数据驱动模板
	err = tmpl.Execute(os.Stdout, map[string]interface{}{"hello": "cde"})
	if err != nil {
		panic(err)
	}
}
