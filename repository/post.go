package repository

import (
	"fmt"
	"sync"
)

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type PostDao struct{}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

// 根据 topic(parent) id 查询帖子列表
func (*PostDao) QueryPostByParentId(parentId int64) ([]*Post, error) {
	posts, ok := postIndexMap[parentId]
	if !ok {
		return nil, fmt.Errorf("not found %v Posts", parentId)
	}
	return posts, nil
}
