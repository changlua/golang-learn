package third_ArraryAndSlice

import (
	"fmt"
	"testing"
)

/**
	一、数组的创建方式
*/
func Test_07_01(t *testing.T) {
	//1、定义
	var balance [10]float32
	//2、定义长度且进行初始化必须使用=
	var balance1 = [11]float32{100, 20.1}
	//根据元素长度来设置数组大小
	numArr := []int{1, 2, 3, 4, 5}
	numArr2 := [...]int{1, 2, 3, 4, 5}
	//{指定位置: 指定位置值}，其他都为0
	numArr3 := [4]int{2: 100}
	//多组指定位置值
	numArr4 := [5]int{1: 100, 2: 300, 4: 500}
	fmt.Println(balance)
	fmt.Println(balance1)
	fmt.Println(numArr)
	fmt.Println(numArr2)
	fmt.Println(numArr3) //[0 0 100 0]
	fmt.Println(numArr4) //[0 0 100 0]
}

/**
	二、数组的遍历
 */
func Test_07_02(t *testing.T) {
	//1、数组赋值以及遍历
	var nums [10]int
	for i := 0; i < len(nums); i++ {
		nums[i] = i * 10 + i
		fmt.Print(nums[i], " ")  //0 11 22 33 44 55 66 77 88 99
	}
	fmt.Println()
	//2、数组的长度
	var nums1 = [...]int{1,2,3}
	fmt.Println(len(nums1)) //3

	//3、for...range遍历
	for i,value := range nums1{
		fmt.Print("[",i,"] = ", value, "  ")  //[0] = 1  [1] = 2  [2] = 3
	}
}

/**
	三、多维数组
*/
func Test_07_03(t *testing.T) {
	//定义二维数组
	var nums = [2][3]int{
		{1,2,3},
		{4,5,6},
	}
	//遍历二维数组
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			fmt.Print(nums[i][j], " ")
		}
		fmt.Println()
	}
	/**
		1 2 3
		4 5 6
	 */
}

/**
	四、数组是值类型，赋值给其他是值拷贝，修改赋值的数组，原先不会改变
*/
func Test_07_04(t *testing.T) {
	//测试改变赋值后的数组，原数组是否会改变？不会，因为是值类型
	var nums = [3]int{1,2,3}
	nums1 := nums
	nums1[2] = 666
	fmt.Println("nums:", nums)  //nums: [1 2 3]
	fmt.Println("num1:", nums1) //num1 [1 2 666]

	//数组的大小不同，表明类型不同。之后切片可以来进行解决这个问题
	//var nums2 [5]int
	//nums = nums2   //此时就会报编译错误，大小不同不能进行赋值
}