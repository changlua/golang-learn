package _2

import (
	"fmt"
	"sync"
	"testing"
)

var (
	x1 int
)

func Test_03_WaitGroup(t *testing.T)  {
	var wg sync.WaitGroup
	wg.Add(5) //添加5个数
	//开启5个协程来进行操作
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done() //表示-1
			addWithLock2()
		}()
	}

	//进行阻塞等待，直到结束
	wg.Wait()
	fmt.Println(x1)
}

func addWithLock2()  {
	for i := 0; i < 2000; i++ {
		//对于公共变量进行自增操作需要进行上锁，防止出现问题
		lock.Lock()
		x1 += 1
		lock.Unlock()
	}
}