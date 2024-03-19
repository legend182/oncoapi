package controller

import (
	"oncoapi/logic"
	"oncoapi/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDrugLevel(c *gin.Context){
	var li *models.LevelIn
	if err := c.ShouldBindJSON(&li);err!=nil{
		zap.L().Error("获取药物等级时参数错误:",zap.Error(err))
		ResponseError(c,CodeInvalidParams)
		return
	}
	
		
}