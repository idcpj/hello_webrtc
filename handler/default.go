package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultHomePageHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/web/index.html")
	return
}
