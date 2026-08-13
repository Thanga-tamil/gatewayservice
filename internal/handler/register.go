package handler

import (
	"encoding/json"
	"gateway/internal/config"
	"gateway/internal/encrypt"
	log "gateway/internal/utils"

	"github.com/gin-gonic/gin"
)

type user struct{
	Msg string `json:"msg"`
	Not string `json:"not"`
}

type registerPayload struct {
	RegisterData string 
}

func RegisterUser(ginCtx *gin.Context){
	var input registerPayload
	ginCtx.ShouldBindBodyWithJSON(&input)
	licensekey := config.ChatSettings.Licensekey
	iv := config.ChatSettings.IV

	decryptedVal, err := encrypt.Decrypt(input.RegisterData, licensekey, iv)
	if err != nil {
		panic(err)
	}
	log.Info("decryptedVal: ", string(decryptedVal))

	var damn user
	if err := json.Unmarshal(decryptedVal, &damn); err != nil {
		panic(err)
	}
	log.Infof("damn: %#v", damn)
}

