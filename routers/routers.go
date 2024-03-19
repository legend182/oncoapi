package routers

import "github.com/gin-gonic/gin"

func SetupRoute(mode string) *gin.Engine{
	if mode == gin.ReleaseMode{
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	v1 :=r.Group("/oncokb/v1")
	v1.POST("getLevel",)
	return r
}