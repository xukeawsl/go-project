package main

import (
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

	r.Run()
}
