package handler

import (
	"net/http"
	"gateway/internal/config"
	"github.com/gin-gonic/gin"
)

func GetLicenseKey(ginCtx *gin.Context){
	result := map[string]any{"status": "200", "licenseKey": config.ChatSettings.Licensekey}
	ginCtx.JSON(http.StatusOK, result) 
}
