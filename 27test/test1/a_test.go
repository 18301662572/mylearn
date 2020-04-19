package main

import (
	"fmt"
	"testing"
)

//单元测试

func TestAdd(t *testing.T){
	c:=Add(1,2)
	fmt.Println(c)
}

