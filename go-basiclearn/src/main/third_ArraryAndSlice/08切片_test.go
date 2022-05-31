package third_ArraryAndSlice

import "fmt"
import "testing"

/**
	一、切片定义与使用
	切片：数组的抽象（动态数组），Go 数组的长度不可改变，在特定场景中这样的集合就不太适用。
	与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
	核心：现有数组的引用。
	描述数据类型：slice像一个结构体，包含三个元素
		1、指针，指向数组中slice指定的开始位置
		2、长度，即slice的长度
		3、最大长度，也就是slice开始位置到数组的最后位置的长度
 */
func Test_08_01(t *testing.T) {
	//1、创建切片。长度为5，容量为10的切片
	mslice := make([]int, 5, 10)
	fmt.Println(mslice,"，长度为：", len(mslice), ",容量为：", cap(mslice))//[0 0 0 0 0] ，长度为： 5 ,容量为： 10

	//2、拿到数组指定位置的切片
	mslice2 := [5]int{1,2,3,4,5}
	mslice3 := mslice2[2:4]  //[2,4)区间位置
	mslice4 := mslice2[:3]  //[0,3)区间位置
	fmt.Println(mslice3) //[3 4]
	fmt.Println(mslice4) //[1 2 3]

	//3、若是对切片修改，其会对原始数组也进行修改
	mslice3[0] = 5
	fmt.Println()
	fmt.Println("切片修改：", mslice3) //切片修改： [5 4]
	fmt.Println("源数组：", mslice2)  //源数组： [1 2 5 4 5]
}

/**
	二、切片的长度与容量
 */
func Test_08_02(t *testing.T) {
	var nums = make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(nums),cap(nums) ,nums)  //len=3 cap=5 slice=[0 0 0]
}

/**
	三、append()和copy()函数
	append()：append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
		1、append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。
		2、当slice中没有剩余空间（即(cap-len) == 0）时，此时将动态分配新的数组空间。返回的slice数组指针将指向这个空间，而原 数组的内容将保持不变；其它引用此数组的slice则不受影响

	copy()：copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
*/
func Test_08_03(t *testing.T) {
	//1、append()追加返回的数组与源数组无关！！！相当于来了层【深拷贝】
	var nums []int
	nums1 := append(nums, 0)
	printSlice(nums)	//len=0,cap=0,slice=[]
	printSlice(nums1)	//len=1,cap=1,slice=[0]
	//添加多个元素
	nums1 = append(nums1, 2,3,4)
	printSlice(nums)	//len=0,cap=0,slice=[]
	printSlice(nums1)	//len=4,cap=4,slice=[0 2 3 4]

	//创建当前切片的两倍长度、两倍容量。【注意：[]int，[]必须写在前面】
	slices := make([]int, len(nums1), (cap(nums1)) * 2)
	//2、copy()也是相当于深拷贝，复制之后两者无关
	copy(nums1, slices)  //将nums1中的内容拷贝到slices中（新创建的）
	slices[2] = 5
	println(nums1)
	println(slices)
}

//定义函数
func printSlice(x []int)  {
	//注意其中有%d,%v，需要使用printf输出
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(x), cap(x), x)
}