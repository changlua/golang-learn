package first_grammer

import (
	"fmt"
	"testing"
)

/**
   一、基本类型
 */
//1、布尔类型：只可以是常量true或false
var bbb bool = true

//2、整数类型
var num int = 100

//3、字符串类型
var str string = "123"

func Test_04_01(t *testing.T){
	//数据类型转换
	//可以自动转型或者手动转型
	var strNum = int(num) //例如字符串"100" => 100
	fmt.Println(strNum)
}

/**
   二、复合类型
	1、指针类型（Pointer）
	2、数组类型
	3、结构化类型(struct)
	4、Channel 类型
	5、函数类型
	6、切片类型
	7、接口类型（interface）
	8、Map 类型
 */