package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("hello world")
	//os.MkdirAll("./a/b/c/d", 0755)
	fmt.Println(os.Getwd())
	fmt.Println(runtime.Caller(1))

	_, fileStr, _, _ := runtime.Caller(1)
	fmt.Println(fileStr)
}