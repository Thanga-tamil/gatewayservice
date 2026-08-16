package router

import (
	"gateway/internal/encrypt"
	"gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

func Route(v1Group *gin.RouterGroup){

	// api/v1 base path routing group
	{
		v1Group.POST("/register", handler.RegisterUser)
		v1Group.GET("/licensekey", handler.GetLicenseKey)

		v1Group.GET("/encrypt", encrypt.Process)
	}

}

