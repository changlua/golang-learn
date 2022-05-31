//每个文件都要归属一个包
package first_grammer

//引入fmt包
import (
	f "fmt"
	"testing"
)

func sayHello(){
	f.Println("hello,world!")
}

//测试函数
func Test_HelloWorld(t *testing.T) {
	sayHello()
}


