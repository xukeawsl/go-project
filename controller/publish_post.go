package controller

import "github.com/xukeawsl/go-project/service"

func PublishPost(topicId int64, content string) *PageData {
	// 检查是否有对应 topic
	if !service.ExistTopic(topicId) {
		return &PageData{
			Code: -1,
			Msg:  "topic_id is error",
		}
	}

	postId, err := service.PublishPost(topicId, content)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: map[string]int64{
			"post_id": postId,
		},
	}
}
