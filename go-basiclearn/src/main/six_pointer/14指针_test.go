package six_pointer

import (
	"fmt"
	"testing"
)

/**
	一、访问指针值、地址
*/
func Test_14_01(t *testing.T) {
	var a int = 20   /* 声明实际变量 */
	var ip *int        /* 声明指针变量 */

	ip = &a  /* 指针变量的存储地址 */
	//*ip = a  //这种赋值方式是错误的！！！，因为ip是指针，我们只能使用*ip来访问值，而不是来进行赋值
	fmt.Printf("a 变量的地址是: %x\n", &a  )
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量的存储地址: %x\n", ip )
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )
}

/**
	二、指针操作
*/
func Test_14_02(t *testing.T) {
	b:=255
	a := &b
	fmt.Println("address of a:", a) //address of a: 0xc00000a358
	fmt.Println("value of a:", *a) //value of a: 255
	*a = 24  //会对原来b的值也进行改变（因为两者指向的地址都是一个地址）
	fmt.Println("value of a:", *a)  //value of a: 24
	fmt.Println("value of b:", b)  //value of b: 24
}

/**
	三、函数数组修改。①数组指针。②
*/
//方式一：采用指针
func test3(arr *[3]int)  {
	(*arr)[0] = 3
}

//方式二：切片。注意：这里不要指定int[]的个数
func test3_02(slice []int)  {
	slice[0] = 4
}

func Test_14_03(t *testing.T) {
	nums := [3]int{1,2,3}
	//执行方案一
	test3(&nums)
	fmt.Println(nums)  //[3 2 3]

	//执行方案二
	test3_02(nums[:])
	fmt.Println(nums)  //[4 2 3]
}

/**
	四、指针数组
 */
const MAX = 3
func Test_14_04(t *testing.T) {
	a := []int{10,100,200}
	var i int
	var ptr [MAX]*int;

	for  i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for  i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i,*ptr[i] )  //使用*数组[下标]来访问指针数组中的值
	}
}

/**
	五、指针的指针  **
*/
func Test_14_05(t *testing.T) {
	var a int = 3000
	var aa *int
	var aaa **int

	aa = &a
	//给**指针，赋值*指针的地址
	aaa = &aa

	//一级指针：也就是a的地址
	fmt.Println(aa)  //0xc00000a358
	//*aaa：相当于表示aa
	fmt.Println(*aaa)  //0xc00000a358
	//**aaa：表示a
	fmt.Println(**aaa)  //3000
}

/**
	六、使用指针函数来交换两数
*/
func swap(a *int, b *int)  {
	var temp = *a
	*a = *b
	*b = temp
}
func Test_14_06(t *testing.T) {
	a := 10
	b := 15
	swap(&a, &b)
	fmt.Println(a)  //15
	fmt.Println(b)  //10
}