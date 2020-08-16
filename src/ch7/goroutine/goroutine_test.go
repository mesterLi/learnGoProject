package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	go say("hello")
	say("world")
}

func say(keyword string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(keyword)
	}
}

func TestChain(t *testing.T) {
	s := []int{1, 2, 3, 4, 45, 23}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <- c, <-c
	t.Log(x,y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}