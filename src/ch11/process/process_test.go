package process

import (
	"fmt"
	"testing"
	"time"
)
func task() {
	i := 0
	for  {
		i++
		fmt.Println("task goroutine", i)
		time.Sleep(1 * time.Second)
	}
}
func TestGoroutine(t *testing.T) {
	i := 0
	go task()
	for {
		i++
		fmt.Println("main goroutine", i)
		time.Sleep(1 * time.Second)
	}
}

func TestMainProcess(t *testing.T) {
	i := 0
	go func() {
		i := 0
		for  {
			i++
			fmt.Println("child process", i)
			time.Sleep(1 * time.Second)
		}
	}()
	for  {
		i++
		fmt.Println("main process", i)
		time.Sleep(1 * time.Second)
		if i >= 3 {
			break
		}
	}
}

func TestChannel(t *testing.T) {
	c := make(chan int)

	go func(c chan <- int) {
		defer close(c)
		for i := 0; i < 4; i++ {
			c <- i
		}
	}(c)
	func(c <- chan int){
		for m := range c{
			fmt.Println(m)
		}
	}(c)
}