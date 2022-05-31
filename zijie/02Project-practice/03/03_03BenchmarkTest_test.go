package _3

import (
	"github.com/bytedance/gopkg/lang/fastrand"
	"testing"
)

//串行
func BenchmarkSelect(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()  //重新计时
	for i := 0; i < b.N; i++ {
		Select()
	}
}

//并行处理
func BenchmarkSelectParallel(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()
	//并行执行
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Select()
		}
	})
}



var ServerIndex [10]int

func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i + 100
	}
}

func Select() int  {
	//return ServerIndex[rand.Intn(10)]
	//使用字节的fastrand
	return ServerIndex[fastrand.Intn(10)]
}