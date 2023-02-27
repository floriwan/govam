package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	c.HTML(http.StatusOK, "user/user.tmpl", gin.H{
		"title": "User website",
	})
}
