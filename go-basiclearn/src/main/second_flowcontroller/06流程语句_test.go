package second_flowcontroller

import (
	"fmt"
	"testing"
)
/**
  1、if语句
 */
func Test_06_01(t *testing.T){
	//1、常见写法
	var num int = 10
	if num < 10 {
		fmt.Println("<10")
	}
	fmt.Println(">=10")

	//2、变体。（前者相当于进行定义变量初始化，后者进行判断）
	if num1 := 20; num1 % 2 == 10 {
		fmt.Println("num1是偶数")
	}
	fmt.Println("num2是奇数")
}

/**
  2、switch
  2.1：基本使用
*/
func Test_06_02(t *testing.T){
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	//①可传入值来作为筛选条件
	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70 : grade = "C"  //case 后可以由多个数值
		default: grade = "D"
	}

	//2、不传入值，直接case来进行表达式筛选
	switch {
		case grade == "A" :
			fmt.Printf("优秀!\n" )
		case grade == "B", grade == "C" :
			fmt.Printf("良好\n" )
		case grade == "D" :
			fmt.Printf("及格\n" )
		case grade == "F":
			fmt.Printf("不及格\n" )
		default:
			fmt.Printf("差\n" );
	}
	fmt.Printf("你的等级是 %s\n", grade ); //A

}

/**
	2.2、fallthrough使用
 */
func Test_06_02_01(t *testing.T) {
	//3、fallthrouht使用
	switch num := 15; num {
	default:
		fmt.Println("最大")
	case 15:
		fmt.Printf("是15\n")
		fallthrough
	case 20:
		fmt.Printf("fallthrough成功执行\n")
	}
}


/**
	2.3、Type Switch：类型判断
*/
func Test_06_02_02(t *testing.T) {
	var x interface{}  //预先声明为接口

	switch i := x.(type) {
		case nil:  //nil是一个预先声明的标识符，表示指针、通道、函数、接口、映射或切片类型
			fmt.Printf(" x 的类型 :%T",i)
		case int:
			fmt.Printf("x 是 int 型")
		case float64:
			fmt.Printf("x 是 float64 型")
		case func(int) float64:
			fmt.Printf("x 是 func(int) 型")
		case bool, string:
			fmt.Printf("x 是 bool 或 string 型" )
		default:
			fmt.Printf("未知型")
	}
}

/**
  三、for循环
 */
func Test_06_03(t *testing.T)  {
	//1、普通for循环
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	//2、单循环条件
	var a = 5
	for a < 10 {
		fmt.Print(a, " ")
		a++
	}
	fmt.Println()
	//3、迭代器遍历：slice、map、数组、字符串等
	nums := [6]int{1,2,3,4,5}
	for i,x := range nums{
		fmt.Println("第",i,"位置，值为：", x)
	}
}

/**
	五、goto跳转标签
 */
func Test_06_05(t *testing.T)  {
	i := 5
	loop: for ; i < 10; i++ {
		if i == 8 {
			i++
			goto loop  //跳转到loop标签
		}
		fmt.Println(i)
	}
}

