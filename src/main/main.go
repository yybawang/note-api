package main

import (
	"cors"
	"log"
	"model/category"
	"model/post"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Cors())
	r.Any("/category", func(c *gin.Context) {
		cate, err := category.List()
		if err != nil {
			jsonError(c, err)
			return
		}
		jsonSuccess(c, cate)
	})

	r.Any("/post", func(c *gin.Context) {
		cateId, err := strconv.Atoi(c.Query("cate_id"))
		if err != nil {
			jsonError(c, err)
			return
		}
		posts, err := post.List(cateId)
		if err != nil {
			jsonError(c, err)
			return
		}
		jsonSuccess(c, posts)
	})

	r.Any("/detail", func(c *gin.Context) {
		postId, err := strconv.Atoi(c.Query("post_id"))
		if err != nil {
			jsonError(c, err)
			return
		}
		posts, err := post.Detail(postId)
		if err != nil {
			jsonError(c, err)
			return
		}
		jsonSuccess(c, posts)
	})

	err := r.Run(":8081") // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Println("Listen run error")
	}
}

func jsonSuccess(c *gin.Context, res interface{}) {
	c.JSON(200, gin.H{"Status": 1, "Message": "OK", "Data": res})
}

func jsonError(c *gin.Context, err error) {
	c.JSON(200, gin.H{"Status": 0, "Message": err, "Data": []string{}})
}
