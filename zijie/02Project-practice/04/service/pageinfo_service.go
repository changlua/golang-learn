package service

import "zijie/02Project-practice/04/dao"

type PageInfo struct {
	Topic *dao.Topic
	PostList []*dao.Post
}

type QueryPageInfoFlow struct {
	pageInfo *PageInfo
}

func (f *QueryPageInfoFlow) Do(topicId int64) (*PageInfo, error) {
	//检查参数

	//准备信息

	//查询数据
	topic := dao.NewTopicDaoInstance().QueryTopicById(topicId)
	postList := dao.NewPostDaoInstance().QueryPostsByParentId(topicId)
	f.pageInfo = &PageInfo{
		Topic:    topic,
		PostList: postList,
	}
	return f.pageInfo, nil
}