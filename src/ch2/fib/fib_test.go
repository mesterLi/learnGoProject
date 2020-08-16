package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	var a int = 1
	var b int = 1
	//var (
	//	a int = 1
	//	b int = 1
	//)
	//a:=1
	//b:=1
	fmt.Println(a)
	for i:=0; i < 5; i++ {
		fmt.Println(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	var (
		a int = 1
		b int = 2
	)
	// 第一种
	//temp := a
	//a = b
	//b = temp
	// 第二种
	a,b = b,a
	t.Log(a, b)
}
