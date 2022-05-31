package main

import (
	"database/sql"  //导入sql驱动
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)



type User struct {
	id int64
	name string
}

func Test_ProtogenesisImpl(t *testing.T) {
	//使用driver（指定mysql） + DSN初始化DB连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/hello")
	//rows, err := db.Query("select id,name from users where id = ?", 1)
	rows, err := db.Query("select id,name from users")
	if err != nil {
		log.Printf("连接异常")
	}
	defer func() {
		err = rows.Close()  //处理完毕，释放连接
	}()
	//定义用户数组
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.id, &user.name)
		if err != nil {
			log.Printf("解析有误")
		}
		users = append(users, user)
	}
	fmt.Println(users)

}