package goPattern

import (
	"fmt"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	tickets, timeout := 3, 3*time.Second

	s := NewSemaphore(tickets, timeout)
	if err := s.Acquire(); err != nil {
		panic(err)
	}
	fmt.Println("信号量减一")

	if err := s.Release(); err != nil {
		panic(err)
	}
	fmt.Println("信号量加一")
}
