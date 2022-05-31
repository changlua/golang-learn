package main

import (
	"fmt"
	"go-gormLearn/db"
)

func main()  {
	db.Init()

	//1、创建用户（无需填充gorm.model，即可自动进行插入更新值）
	//users := make([]*db.User, 2)
	//users[0] = &db.User{
	//	Model:    gorm.Model{},
	//	Name:     "changlu",
	//	Password: "123456",
	//}
	//users[1] = &db.User{
	//	Model:    gorm.Model{},
	//	Name:     "test",
	//	Password: "123456",
	//}
	//db.CreateUser(nil, users)
	
	//2、更新用户（会自动更新updatetime）
	//方式一：借助struct
	//u := &db.User{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name:     "changlu111",
	//	Password: "12132232",
	//}
	//db.UpdateUser(nil, u)	//通过！
	//方式二：借助map
	//username := "changlu666"
	//password := "122222"
	//db.UpdateUser2(nil, 1, &username, &password)//通过！
	//恢复软删除
	db.RecoverUserDeleted(nil, 1)

	//3、删除用户
	//3.1、软删除
	//db.SofeDeleteUser(nil, 1)
	//3.2、硬删除
	//db.DeleteUser(nil, &db.User{
	//	Model:    gorm.Model{
	//		ID: 1,
	//	},
	//})

	//4、查询用户
	//4.1、查询所有未软删除用户
	//users, _ := db.QueryAllUsers(nil)
	//for _, user := range users {
	//	fmt.Println(user)
	//}
	//4.2、查询所有用户（包含软删除的）
	//users, _ := db.QueryAllIncludeDeletedUsers(nil)
	//for _, user := range users {
	//	fmt.Println(user)
	//}
	//4.3、只查询软删除的用户
	//users, _ := db.QueryAllDeletedUsers(nil)
	//for _, user := range users {
	//	fmt.Println(user)
	//}
	//4.4、分页查询
	users, _ := db.PageQueryUser(nil, 1, 1)
	for _, user := range users {
		fmt.Println(user)
	}

}
