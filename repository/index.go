package repository

import (
	"bufio"
	"encoding/json"
	"os"
)

var (
	topicIndexMap map[int64]*Topic  // 通过 topic id 索引 Topic
	postIndexMap  map[int64][]*Post // 通过 post id 索引一组 Post
)

func Init(filePath string) error {
	err := initTopicIndexMap(filePath)
	if err != nil {
		return err
	}
	err = initPostIndexMap(filePath)
	if err != nil {
		return err
	}
	return nil
}

func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	defer open.Close()

	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)

	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

// 从指定文件中读取原有数据写入 postIndexMap
func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return nil
	}
	defer open.Close()

	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)

	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		postTmpMap[post.ParentId] = append(postTmpMap[post.ParentId], &post)
	}
	postIndexMap = postTmpMap
	return nil
}
