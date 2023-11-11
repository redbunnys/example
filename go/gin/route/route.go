package route

import (
	"github.com/gin-gonic/gin"
	v1 "go-gin/api/v1"
	"go-gin/middleware"
	"go-gin/middleware/jwt"
)

func RouteInit(r *gin.Engine) {
	r.Use(gin.Logger(), middleware.Cors())
	r.GET("/signin", v1.Home)

	r.Use(jwt.Jwt())
	r.GET("/parse", v1.Parse)
	r.GET("/refresh", v1.Refresh)
}
