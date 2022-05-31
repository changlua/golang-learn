package first_grammer

import (
	"fmt"
	"testing"
)

/**
	一、单变量声明
 */
//4、直接声明并赋值
var username string = "changluchanglu"

func Test_02_01(t *testing.T){
	//1、指定类型
	var name string
	name = "name"

	//2、自动推断(自动根据值来进行推断)
	var name1 = "name1"
	fmt.Println(name)

	//3、省略var（这种方式只能用在函数体中，不能用在全局变量）
	//3.1单个变量赋值时，必须是新变量
	name3 := "name3"
	//3.2多个变量赋值时，必须有一个时新变量
	name1,name2 := "hh","name2"
	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(username)

}

/**
	二、多变量声明
*/
func Test02_02(t *testing.T){
	//1、逗号分割，声明与赋值分开，有默认值为空
	var name1, name2, name3 string
	name1, name2, name3 = "xm", "xh", "ll"   //赋值应该在函数体内

	//2、直接赋值
	var name4, name5, name6 string = "xm", "xh", "ll"

	//3、集合类型
	var (
		name7 string = "hh"
		name8 string = "xx"
	)
	fmt.Println(name1, name2, name3)
	fmt.Println(name4, name5, name6)
	fmt.Println(name7, name8)
}