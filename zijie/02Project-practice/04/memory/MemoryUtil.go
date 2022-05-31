package memory

import (
	"bufio"
	"encoding/json"
	"os"
	"zijie/02Project-practice/04/dao"
)

//定义缓存：实现O(1)复杂度来查询
var (
	TopicIndexMap map[int64] *dao.Topic
	postIndexMap  map[int64][] *dao.Post
)

//初始化话题数据索引
func initTopicIndexMap(filePath string) error  {
	//打开指定文件
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	//创建一个scanner读取对象
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*dao.Topic)
	//遍历每行
	for scanner.Scan() {
		text := scanner.Text()
		var topic dao.Topic
		//进行json序列化
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	TopicIndexMap = topicTmpMap
	return nil
}

