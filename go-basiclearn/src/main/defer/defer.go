package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string   //大写表示public，可访问的
	Age int
}

func (s Student) SayHello(name string)  {   //注意点：①名字大写才能够直接被查看到。②不加指针的方法才能够被查看到
	fmt.Println("hello,", name)
}

//通过反射打印出万能接口实例中的类型与值
func reflectNum(arg interface{})  {
	inputType := reflect.TypeOf(arg)
	fmt.Println(inputType)//main.Student
	inputValue := reflect.ValueOf(arg)
	fmt.Println(inputValue)//{changlu 18}

	//1、获取类中的字段
	//通过类型中的字段数量来进行遍历
	fmt.Println("\n打印类中的属性与值：")
	for i := 0; i < inputType.NumField(); i++ {  //inputType.NumField()：表示类型字段的数量
		field := inputType.Field(i)  //取出第i个filed
		value := inputValue.Field(i).Interface() //取出指定下标的值
		//分别打印：属性的名称、类型、值
		fmt.Printf("%s: %v = %v\n",field.Name, field.Type, value)  //%s：字符串，%v：原样输出
	}

	//2、获取类中的方法
	fmt.Println("\n打印类中的方法：")
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}

	//3、调用指定的SayHello方法
	sayHello := inputValue.MethodByName("SayHello")
	result := sayHello.Call([]reflect.Value{reflect.ValueOf("changlu")})  //传入一个数组对象 or nil
	fmt.Println(result)  //打印返回值
}

func main() {
	student := Student{"changlu", 18}
	reflectNum(student)
}

