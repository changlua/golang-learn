package first_grammer

import (
	"fmt"
	"testing"
)

//1、结构体多行声明
//规范：名字大写
type User struct {
	UserName string
	Email string
}

// 测试结构体
func Test_01struct(t *testing.T)  {
	//多行初始化，有,号
	u := User{
		UserName: "changlu",
		Email: "939974883@qq.com",
	}
	fmt.Println(u)
}

//2、接口定义
//规范：单个函数的结构名以 “er” 作为后缀，例如 Reader , Writer
type Reader interface {
	//定义的接口
	Read(p []byte) (n int ,err error)
}

//3、变量命名
//①遵循驼峰命名法。②变量私有private，也就是小写开头；变量公共使用大写开头(外部包可使用)
var isExists bool
var hasConflict bool
var canManage bool
var allowGitHook bool

//4、常量
//规范：需要全部使用大写开头
const APP_VER = "1.0"
// 枚举常量
type Scheme string
const (
	HTTP  Scheme = "http"
	HTTPS Scheme = "https"
)