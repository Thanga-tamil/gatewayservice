package router

import (
	"io"
	"net/http"
	"gateway/internal/config"
	"gateway/internal/encrypt"
	"gateway/internal/handler"
	log "gateway/internal/utils"
	"github.com/gin-gonic/gin"
)

func Route(v1Group *gin.RouterGroup){

	// api/v1 base path routing group
	{
		v1Group.POST("/register", handler.RegisterUser)
		v1Group.GET("/licensekey", handler.GetLicenseKey)

		v1Group.GET("/encrypt", process)
	}

}

func process(ctx *gin.Context){
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}

	licensekey := config.ChatSettings.Licensekey
	iv := config.ChatSettings.IV
	encryptedVal, err := encrypt.Encrypt(string(jsonData), licensekey, iv)
	if err != nil {
		panic(err)
	}
	log.Info("encryptedVal: ", encryptedVal)

	ctx.JSON(http.StatusOK, map[string]string{"data": encryptedVal, "message": "Data Encrypted Successfull"})
}

