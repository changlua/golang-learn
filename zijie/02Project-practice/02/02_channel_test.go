package _2

import (
	"fmt"
	"sync"
	"testing"
)

func Test02_channle(t *testing.T)  {
	//anli2_NoCacheChannel() //无缓存
	//anli3_CacheChannel() //带缓存通道(容量1个)
	anli3_CacheChannel_upgrade() //带缓存通道(容量10个)
}

//案例1：使用同步锁来实现同步
func anli1_Mutex()  {
	var mu sync.Mutex
	//在main方法中两次执行上锁，第二次加锁时会因为锁已经占用(该锁并非是递归锁)而阻塞
	mu.Lock()
	go func() {
		fmt.Println("changluchanglu")
		mu.Unlock()
	}()
	mu.Lock()
}

//2、使用无缓存通道来实现同步（优化案例1）
func anli2_NoCacheChannel()  {
	done := make(chan int)
	go func() {
		fmt.Println("changlu")
		<- done //表示接收完成
	}()

	done <- 1  //只有接收完成之后，发送操作才可能完成
	fmt.Println("test")
}

//3、缓存通道实现
//说明：案例2的虽然可以正确同步，但是对通道的缓存大小太敏感，如果通道有缓存，就无法保证 main() 函数退出之前后台线程能正常打印了，
//更好的做法是将通道的发送和接收方向调换一下，这样可以避免同步事件受通道缓存大小的影响
func anli3_CacheChannel()  {
	done := make(chan int, 1)
	go func() {
		fmt.Println("changlu")
		done <- 1
	}()
	<-done
	fmt.Println("test")
}

//扩展为10个线程：
func anli3_CacheChannel_upgrade()  {
	done := make(chan int, 10)
	//通过通道容量来确定循环次数
	for i := 0; i < cap(done); i++ {
		go func(j int) {
			fmt.Printf("changlu %v\n",j)
			done <- j
		}(i)
	}

	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

func CalSquare()  {
	src := make(chan int)  //无缓冲通道
	//fmt.Println(reflect.TypeOf(src))  //chan int
	dest := make(chan int, 3)  //有缓冲通道
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i  //将i放置到通道中
		}
	}()

	go func() {
		defer close(dest)
		for i := range dest{
			dest <- i * i
		}
	}()

	for i := range dest {
		//复杂操作
		println(i)
	}
}