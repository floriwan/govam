package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Registration(c *gin.Context) {
	c.HTML(http.StatusOK, "registration.tmpl", gin.H{
		"title": "registration",
	})
}
