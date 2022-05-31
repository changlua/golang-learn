package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

//默认gorm对struct字段名使用Snake Case命名风格转换成mysql表字段名(需要转换成小写字母)
type Food struct {
	Id         int  //表字段名为：id
	Name       string //表字段名为：name
	Price      float64 //表字段名为：price
	TypeId     int  //表字段名为：type_id
	//我们也可以进行自定义对应映射的字段名：字段定义后面使用两个反引号``包裹起来的字符串部分叫做标签定义，这个是golang的基础语法，不同的库会定义不同的标签，有不同的含义
	CreateTime time.Time `gorm:"column:createtime"`  //表字段名为：createtime  //若是数据库中是DateTime，那么对应这里使用time，之后序列化可以进行转换时间日期
}


//设置表名，可以通过给Food struct类型定义 TableName函数，返回一个字符串作为表名
func (v Food) TableName()string  {
	return "food"
}

func Test_gormImpl(t *testing.T)  {
	db, err := connect()
	if err != nil {
		log.Printf("")
	}
	//v2版本中db没有close()方法：https://www.cnblogs.com/chengqiang521/p/15122102.html
	sqlDB, _  := db.DB()
	var foods []Food
	//查询所有的food列表记录
	db.Find(&foods)
	fmt.Println(foods)
	//延时关闭数据库连接
	defer sqlDB.Close()
}

func connect()  (db *gorm.DB, err error)  {
	//配置MySQL连接参数
	username := "root"  //账号
	password := "123456" //密码
	host := "127.0.0.1" //数据库地址，可以是Ip或者域名
	port := 3306 //数据库端口
	Dbname := "hello" //数据库名
	timeout := "10s" //连接超时，10秒
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db,nil
}