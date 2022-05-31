package five_function

import (
	"fmt"
	"testing"
)

/**
	一、函数的定义
*/
func test01(str string) string  {  //若是有return，那么就一定要有返回类型
	return str
}

// 取得一个最大值
func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func Test_11_01(t *testing.T) {
	fmt.Println(test01("123"))  //"123"
	fmt.Println(max(1,5))  //5
}


/**
	二、函数可变参、参数传递
*/
//1、可变参函数定义
func test2(arg ...int)  {
	//_ , i：这种写法可以忽略掉index索引位置
	for _ , value := range arg {
		fmt.Print(value, " ")
	}
}
func Test_11_02(t *testing.T) {
	//不允许如：nums := [2]int{1,2}
	//只允许一个接着一个逗号
	test2(1,2,3)   //1 2 3

	//2、参数传递
	//2.1、参数传递（值传递）。
	//go语言中可以这样定义函数
	test3 := func(x float64) float64 {
		return x
	}
	fmt.Println(test3(5.0))  //5

	//2.2、引用传递：整型变量也可以传递过去地址
	test4 := func(a *int) int {
		fmt.Println("函数中的a地址：", a)  //函数中的a地址： 0xc00000a360
		*a += 1  //修改a的值
		return *a
	}
	var a = 10
	fmt.Println("&a: ", &a)  //&a:  0xc00000a360
	//调用函数时使用&变量，表示传递地址
	fmt.Println(test4(&a))    //11
	fmt.Println(a)   //11

}

/**
	三、函数的返回值
*/
//1、返回多个值
func test3(A, B int) (add int, Multiplied int) {  //多个返回值要用()包裹
	add = A + B
	Multiplied = A * B
	return add, Multiplied
}

func Test_11_03(t *testing.T) {
	//空白标识符
	add,_ := test3(1,2 )
	fmt.Println(add)   //3
}