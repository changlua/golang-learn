package dao

import (
	"sync"
	"time"
)

//文章
type Topic struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreateTime int64 `json:"create_time"`  //时间戳
}

var (
	topicDao *TopicDao
	topicOnce sync.Once  //适用于高并发场景下只执行一次的场景，基于once实现的就是单例模式
)

type TopicDao struct {
}

// NewTopicDaoInstance 单例获取文章dao查询
func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}


//QueryTopicById 根据id来查询文章
func (*TopicDao) QueryTopicById(id int64) *Topic {
	//return memory.TopicIndexMap[id]
	//假数据模拟
	return &Topic{
		Id: id,
		Title: "changlu",
		Content: "xxxx",
		CreateTime: time.Now().Unix(),
	}
}


