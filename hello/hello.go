package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpasswd")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	r.Run(":80")
}
