package main

import (
	"fmt"
	"os"
)

func main() {
	// 不支持返回值
	// os.Exit
	fmt.Println("hello world")
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
	os.Exit(1)
}