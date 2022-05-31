package _2

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test02_lock_test(t *testing.T)  {
	//开启5个协程来进行操作
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(x)
}

//多个变量快速定义，都是var类型【好处：不用每个都写var】
var (
	x int64
	lock sync.Mutex
)

func addWithLock()  {
	for i := 0; i < 2000; i++ {
		//对于公共变量进行自增操作需要进行上锁，防止出现问题
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}
