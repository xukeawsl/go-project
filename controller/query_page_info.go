package controller

import (
	"strconv"

	"github.com/xukeawsl/go-project/service"
)

// 返回给客户的 http 响应内容
type PageData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func QueryPageInfo(topicIdStr string) *PageData {
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		return &PageData{}
	}
	pageInfo, err := service.QueryPageInfo(topicId)
	if err != nil {
		return &PageData{}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: pageInfo,
	}
}