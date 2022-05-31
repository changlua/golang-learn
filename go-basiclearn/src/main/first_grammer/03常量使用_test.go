package first_grammer

import (
	"fmt"
	"testing"
)

/**
	常量定义与使用
 */
//1、显示定义（指明类型）
const A string = "abc"
//2、隐式定义（不指明类型）
const B = "abc"

func Test_03_01(t *testing.T){
	//多重赋值（函数体中）
	const C, D, E = "c","d","e"
	fmt.Println(C, D, E) //c d e
	const AREA float32 = 5.0
	const WIDTH float32 = 6.0
	//注意：这里相乘必须是float同一类型
	fmt.Println("面积为：", AREA * WIDTH) //面积为： 30

	//常量可作为枚举，没有逗号
	const (
		Unknow = 0
		Femal = 1
		Male = 2
	)
	fmt.Println(Unknow, Femal, Male) //0 1 2

	//常量可以不指定类型以及初始化值，则与上一行非空常量右值相同
	const (
		F string = "ff" //指定类型
		G   //默认与上面类型一致
		H = "abc"   //直接=
	)
	fmt.Println(F, G, H) //ff ff abc
}

/**
  iota特殊常量使用（自增默认是int类型）
 */
//第一个 0，之后会依次递加
const (
	a = iota
	b = iota
	c = iota
)

//第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1
//若是第一个为iota+1，那么默认就是从1开始
const (
	d = iota + 1
	v
	z
)

//这里主要就是熟悉一下之前的，当第二个iota出现时，这个值也就是对应从aa到其次数累加1
//bb连上aa+1  dd与cc一致  gg与ff一致  hh连上之前数个+1，ii继续hh+1
const (
	aa = iota
	bb
	cc = "hh"
	dd
	ff = 100
	gg
	hh = iota
	ii
)

func Test03_02(t *testing.T){
	fmt.Println(a, b, c)                        //0 1 2
	fmt.Println(d, v, z)                        //0 1 2
	fmt.Println(aa, bb, cc, dd, ff, gg, hh, ii) //0 1 hh hh 100 100 6 7
}
