package test

import (
	"fmt"
	"testing"
	"time"
)

func TestForSelect(t *testing.T) {
	chs := make(chan int, 10)
	go func() {
		for {
			chs <- 1
			break
		}
	}()
	time.Sleep(time.Millisecond * 100)
loop:
	for {
		fmt.Println("begin")
		time.Sleep(time.Second)
		select {
		case <-chs:
			fmt.Println("continue")
			continue
		default:
			fmt.Println("break")
			break loop
		}
	}
}
