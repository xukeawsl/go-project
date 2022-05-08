package repository

import (
	"fmt"
	"sync"
)

type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type TopicDao struct{}

// 单例模式创建对象
var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

// 通过 id 查询 Topic 的公开接口
func (*TopicDao) QueryTopicById(id int64) (*Topic, error) {
	topic, ok := topicIndexMap[id]
	if !ok {
		return nil, fmt.Errorf("not found %v Topic", id)
	}
	return topic, nil
}
