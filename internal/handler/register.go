package handler

import (
	"encoding/json"
	"gateway/internal/config"
	"gateway/internal/encrypt"
	"gateway/internal/service"
	"gateway/internal/utils"
	log "gateway/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(reqCtx *gin.Context){
	var registerUser utils.Registeruser

	reqCtx.ShouldBindBodyWithJSON(&registerUser)

	userPayload, err := bind_request_body(&registerUser)
	if err != nil {
		response := utils.Bad(http.StatusBadRequest, err.Error())
		reqCtx.JSON(http.StatusBadRequest, response)
		return 
	}

	if err := service.ValidateRegisterUserPayload(userPayload); err != nil {
		response := utils.Bad(http.StatusBadRequest, err.Error())
		reqCtx.JSON(http.StatusBadRequest, response)
		return 
	}

	service.ProcessNewUserRegistration(reqCtx, userPayload)
}

func bind_request_body(registerUser *utils.Registeruser) (*utils.Data, error) {
	licensekey := config.ChatSettings.Licensekey
	iv := config.ChatSettings.IV
	var user utils.Data

	decryptedVal, err := encrypt.Decrypt(registerUser.RegisterData, licensekey, iv)
	if err != nil {
		log.Error("Err while descrypting register user payload: ", err.Error())
		return &user, err
	}

	if err := json.Unmarshal(decryptedVal, &user); err != nil {
		log.Error("Err while unmarshalling decrypted registration payload to struct: ", err.Error())
		return &user, err
	}

	return &user, nil
}
