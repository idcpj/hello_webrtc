package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Allow() gin.HandlerFunc {
	return func(g *gin.Context) {
		g.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		g.Writer.Header().Set("Access-Control-Max-Age", "86400")
		g.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		g.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		g.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if g.Request.Method == "OPTIONS" {
			g.AbortWithStatus(200)
		} else {
			g.Next()
		}
	}

}
