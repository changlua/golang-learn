package fourth_compoundType

import (
	"fmt"
	"testing"
)

/**
	一、定义、创建、查看、获取map长度
 */
func Test_09_01(t *testing.T) {
	//1、定义Map
	var countryCapitalMap map[string]string   //默认定义类型为nil，相当于null
	//创建集合：必须要进行创建才能够添加key、value键值对
	countryCapitalMap = make(map[string]string)
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	//2、使用key来输出map
	for key, value := range countryCapitalMap {
		fmt.Println("key:",key,",values:",value)
	}

	//3、查看map中是否存在元素
	captial, exists := countryCapitalMap["France"]  //这里返回两个值，前者是对应取到的value；后者为true或false，表示是否取到
	if !exists {
		//若是不存在
		fmt.Println("United States is not a city")
	}
	//若是存在
	fmt.Println(captial) // Paris

	//4、获取map的长度
	fmt.Println("map的长度：", len(countryCapitalMap))  //map的长度： 4
}

/**
	二、map删除：delete(map, key)
 */
func Test_09_02(t *testing.T) {
	var myMap map[string]string = make(map[string]string)
	myMap["a"] = "a"
	myMap["b"] = "b"
	myMap["c"] = "c"
	fmt.Println(myMap)   //map[a:a b:b c:c]
	//删除key为a的map，没有返回值
	delete(myMap, "a")
	fmt.Println(myMap)   //map[b:b c:c]
}

/**
	三、ok-idiom：通过取key，来获取到value以及是否取到
*/
func Test_09_03(t *testing.T) {
	m := make(map[string]string)
	m["a"] = "a"
	result,ok := m["a"]
	fmt.Println(result)  //a
	fmt.Println(ok)      //true
}

/**
	四、Map引用类型：引用赋值给其他引用，对赋值后的引用修改实际上也会修改源map
*/
func Test_09_04(t *testing.T) {
	myMap := make(map[string]string)
	myMap["a"] = "a"
	myMap["b"] = "b"
	myMap["c"] = "c"

	//进行引用
	myMap2 := myMap
	myMap2["a"] = "aaa"
	fmt.Println(myMap2)  //map[a:aaa b:b c:c]
	fmt.Println(myMap)   //map[a:aaa b:b c:c]

	//可以来进行判断是否为null
	var myMap3 map[string]string
	if myMap3 == nil {
		fmt.Println("666")
	}
	//666
}