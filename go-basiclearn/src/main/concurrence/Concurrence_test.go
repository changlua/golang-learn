package main

import (
	"fmt"
	"testing"
	"time"
)

func goFunc(i int) {
	i++
	fmt.Println("goroutine ", i, "...")
}

func Test_concurrence_01(t *testing.T){
	var i int = 1
	//总共开启10000个协程
	for i := 0; i < 10000; i++ {
		go goFunc(i) //开启一个并发协程
	}
	time.Sleep(time.Second)
}

