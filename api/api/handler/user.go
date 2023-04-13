package handler

import (
	"api/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func GetInformationUsingAPI(logger *zap.SugaredLogger, service service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := service.InsertInformation(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed"})
			logger.Debugf("Error:%s\n", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "successful"})
	}
}

func GetAllUsersOfDB(logger *zap.SugaredLogger, service service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := service.SelectInformation(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed"})
			logger.Debugf("Error:%s\n", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": data})
	}
}
