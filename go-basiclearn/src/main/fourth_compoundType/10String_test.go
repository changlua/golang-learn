package fourth_compoundType

import (
	"fmt"
	"testing"
)

/**
	一、定义与使用String
*/
func Test_10_01(t *testing.T) {
	//1、定义与使用
	str := "hello,changlu"
	fmt.Println(str)  //hello,changlu

	//2、遍历字符串
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])  //h e l l o , c h a n g l u
	}

	//扩展：Strings包
}