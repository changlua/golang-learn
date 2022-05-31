package controller

import (
	"strconv"
	"zijie/02Project-practice/04/service"
)

//通用响应值
type PageData struct {
	Code  int64         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限  -1  失败
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据内容
}

//查询页面信息
func QueryPageInfo(topicIdStr string) *PageData {
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64) //转为十进制，大小为64位【int64】
	if err != nil {
		return &PageData{500, "转换有误", nil}
	}
	//调用service来进行查询
	pageInfo, err := (&service.QueryPageInfoFlow{}).Do(topicId)
	if err != nil{
		return &PageData{500, err.Error(), nil}
	}
	return &PageData{200, "查询成功", pageInfo}
}