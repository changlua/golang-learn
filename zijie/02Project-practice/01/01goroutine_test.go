package _1

import (
	"fmt"
	"testing"
	"time"
)

func hello(i int)  {
	println("hello goroutine : " + fmt.Sprint(i))
}

func Test_01goroutine(t *testing.T)  {
	for i := 0; i < 5; i++ {
		//开启一个协程
		go func(j int) {
			hello(j)
		}(i)
	}
	//1s足以5个协程跑完
	time.Sleep(time.Second)
}
