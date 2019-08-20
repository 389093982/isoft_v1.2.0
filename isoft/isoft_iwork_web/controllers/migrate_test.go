package controllers

import (
	"fmt"
	"regexp"
	"testing"
	"time"
)

func Test_MigrateName(t *testing.T) {
	fmt.Println(regexp.MatchString("^[0-9]{14}_(CREATE|ALTER|DELETE|DROP)_[a-zA-Z]+\\.sql$", "20060102150405_CREATE_HELLOWORLD.sql"))
}

func Test_MigrateName2(t *testing.T) {
	fmt.Println(time.Now().Format("20060102150405"))
}
