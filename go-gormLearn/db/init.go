package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init()  {
	var err error
	DB, err = gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/hello?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:logger.Default.LogMode(logger.Info),  //打印执行的sql
		},
	)
	if err != nil {
		panic(err)
	}

	//进行表的创建
	m := DB.Migrator()
	if !m.HasTable(&User{}) {
		if err = m.CreateTable(&User{}); err != nil {
			panic(err)
		}
	}
}