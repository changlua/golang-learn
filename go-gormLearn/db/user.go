package db

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string     `gorm:"type:varchar(32);not null"`
	Password string `gorm:"type:varchar(32);not null"`
}

func (v User) TableName()string  {
	return "user"
}

//创建用户
func CreateUser(ctx context.Context, users []*User) error {
	//INSERT INTO `user` (`created_at`,`updated_at`,`deleted_at`,`name`,`password`) VALUES ('2022-05-28 19:30:45.339','2022-05-28 19:30:45.339',NULL,'changlu','123456'),()
	return DB.WithContext(ctx).Create(users).Error
}

//修改用户
//方式一：通过struct来进行更新
func UpdateUser(ctx context.Context, user *User) error {
	//UPDATE `user` SET `updated_at`='2022-05-28 19:31:24.468',`name`='changlu111',`password`='12132232' WHERE `user`.`deleted_at` IS NULL AND `id` = 1
	return DB.WithContext(ctx).Model(&user).Select("Name", "Password").Updates(user).Error
}
//方式二：通过map进行更新
func UpdateUser2(ctx context.Context, ID uint, username *string, password *string )error  {
	params := make(map[string]interface{})
	if username != nil {
		params["Name"] = username
	}
	if password != nil {
		params["Password"] = password
	}
	//UPDATE `user` SET `name`='changlu666',`password`='122222',`updated_at`='2022-05-28 19:32:05.675' WHERE id = 1 AND `user`.`deleted_at` IS NULL
	return DB.WithContext(ctx).Model(&User{}).Where("id = ?", ID).Updates(params).Error
}
//恢复软删除
func RecoverUserDeleted(ctx context.Context, ID uint)error  {
	//UPDATE `user` SET `deleted_at`=NULL,`updated_at`='2022-05-28 19:16:19.579' WHERE `ID` = 1
	return DB.WithContext(ctx).Model(&User{}).
				Unscoped(). //针对已软删除的，此时就不会带上`user`.`deleted_at` IS NULL
				Where("ID", ID).Update("deleted_at", nil).Error
}

//删除用户
//方式一：软删除
func SofeDeleteUser(ctx context.Context, ID uint) error {  //uint指的是无符号整数，其大小根据当前的平台来决定
	//UPDATE `user` SET `deleted_at`='2022-05-28 19:32:30.124' WHERE id = 1 AND `user`.`deleted_at` IS NULL
	return DB.WithContext(ctx).Where("id = ?", ID).Delete(&User{}).Error
}
//方式二：硬删除
func DeleteUser(ctx context.Context, user *User) error {
	//调用Unscoped即可表示硬删除
	//DELETE FROM `user` WHERE `user`.`id` = 1
	return DB.WithContext(ctx).Unscoped().Delete(user).Error
}

//查询用户（未被软删除的）
func QueryAllUsers(ctx context.Context)([]*User, error)  {
	var users []*User
	//SELECT * FROM `user` WHERE `user`.`deleted_at` IS NULL
	if err := DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil,err
	}
	return users, nil
}

//查询所有用户（包含软删除的）
func QueryAllIncludeDeletedUsers(ctx context.Context)([]*User, error)  {
	var users []*User
	//SELECT * FROM `user`
	if err := DB.WithContext(ctx).Unscoped().Find(&users).Error; err != nil {
		return nil,err
	}
	return users, nil
}

//只查询软删除用户
func QueryAllDeletedUsers(ctx context.Context)([]*User, error) {
	var users []*User
	//SELECT * FROM `user` WHERE deleted_at <> ''
	if err := DB.WithContext(ctx).Unscoped().Where("deleted_at <> ?", "").Find(&users).Error;err != nil {
		return nil, err
	}
	return users, nil
}

//分页查询
func PageQueryUser(ctx context.Context, pageNo, pageSize int)([]*User, error)  {
	tx := DB.WithContext(nil).Model(&User{})
	//...字段补充  tx.where()...
	var users []*User
	if err := tx.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}





