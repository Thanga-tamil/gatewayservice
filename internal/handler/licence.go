package handler

import (
	"net/http"
	"gateway/internal/repo"
	"github.com/gin-gonic/gin"
)

func GetLicenseKey(ginCtx *gin.Context){

	licenseKey := repo.FetchLicenseKey()

	result := map[string]any{"status": "200", "licenseKey": licenseKey}

	ginCtx.JSON(http.StatusOK, result) 

}
