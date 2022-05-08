package service

import (
	"errors"
	"time"
	"unicode/utf16"

	idworker "github.com/gitstliu/go-id-worker"
	"github.com/xukeawsl/go-project/repository"
)

var idGen *idworker.IdWorker

func init() {
	idGen = &idworker.IdWorker{}
	idGen.InitIdWorker(1000, 1)
}

type PublishPostFlow struct {
	topicId int64
	content string
	postId  int64
}

// 得到被分配的 post_id
func PublishPost(topicId int64, content string) (int64, error) {
	return NewPublishPostFlow(topicId, content).Do()
}

func NewPublishPostFlow(topicId int64, content string) *PublishPostFlow {
	return &PublishPostFlow{
		topicId: topicId,
		content: content,
	}
}

func (f *PublishPostFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.postId, nil
}

// 检查长度
func (f *PublishPostFlow) checkParam() error {
	if len(utf16.Encode([]rune(f.content))) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

func (f *PublishPostFlow) publish() error {
	post := &repository.Post{
		ParentId:   f.topicId,
		Content:    f.content,
		CreateTime: time.Now().Unix(),
	}

	// 生成唯一的 postId
	id, err := idGen.NextId()
	if err != nil {
		return err
	}

	post.Id = id

	// 将 post 插入到 map 中
	if err := repository.NewPostDaoInstance().InsertPost(post); err != nil {
		return err
	}
	f.postId = post.Id
	return nil
}

func ExistTopic(topicId int64) bool {
	return repository.NewTopicDaoInstance().Exist(topicId)
}
