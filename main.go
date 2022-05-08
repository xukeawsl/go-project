package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xukeawsl/go-project/controller"
	"github.com/xukeawsl/go-project/repository"
)

func main() {
	if err := repository.Init("./data/"); err != nil {
		os.Exit(-1)
	}

	r := gin.Default()

	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})

	r.POST("/community/page/post", func(c *gin.Context) {
		var request repository.PostRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data := controller.PublishPost(request.ParentId, request.Content)
		c.JSON(http.StatusOK, data)
	})

	r.Run()
}
