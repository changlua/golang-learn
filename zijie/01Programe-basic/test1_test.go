package main

import (
	"fmt"
	"testing"
)

func Test_test111(t *testing.T) {
	buf := make([]byte, 5)
	buf[0] = 1
	buf[1] = 2
	buf[2] = 3
	buf[3] = 4
	buf[4] = 5
	fmt.Println(buf[:4])
}