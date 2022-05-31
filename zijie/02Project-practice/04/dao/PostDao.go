package dao

import (
	"sync"
	"time"
)

//回复
type Post struct {
	Id int64 `json:"id"`
	ParentId int64 `json:"parent_id"`
	Content string `json:"content"`
	CreateTime int64 `json:"create_time"`
}


var (
	postDao *PostDao
	postOnce sync.Once
)

type PostDao struct {
}

//NewPostDaoInstance 单例获取帖子
func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

//QueryPostsByParentId 根据id来获取帖子
func (*PostDao) QueryPostsByParentId(id int64) []*Post {
	//return postIndexMap[id]\
	post := &Post{
		Id:         id,
		ParentId:   id,
		Content:    "好文好文1",
		CreateTime: time.Now().Unix(),
	}
	post1 := &Post{
		Id:         id,
		ParentId:   id,
		Content:    "好文好文2",
		CreateTime: time.Now().Unix(),
	}
	post2 := &Post{
		Id:         id,
		ParentId:   id,
		Content:    "好文好文3",
		CreateTime: time.Now().Unix(),
	}
	return []*Post{ post,post1, post2}
}