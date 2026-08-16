package service

import (
	"time"
	"errors"
	"strings"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gateway/internal/config"
	"gateway/internal/encrypt"
	log "gateway/internal/utils"
	model "gateway/internal/utils"
)

func BindRequest(registerUser *model.Registeruser) (*model.Data, error) {
	var user model.Data
	iv := config.ChatSettings.IV
	licensekey := config.ChatSettings.Licensekey

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

func ValidateRegisterUserPayload(payload *model.Data) error {

	if strings.TrimSpace(payload.UserId) == "" { 
		return errors.New("userId can not be null or empty")
	} else if strings.TrimSpace(payload.NickName) == "" { 
		return errors.New("name can not be null or empty")
	} else if strings.TrimSpace(payload.Email) == "" {
		return errors.New("email can not be null or empty")
	} else if strings.TrimSpace(payload.MobileNumber) == "" {
		return errors.New("password can not be null or empty")
	}

	return nil
}


func ProcessNewUserRegistration(ctx *gin.Context, payload *model.Data) error {
	user := model.User{
		UserId: payload.UserId,
		Nickname: payload.NickName,
		IsDeleted: false,
		ContactPermission: payload.ContactPermission,
		CreatedAt: time.Now(),
		EmailId: payload.Email,
		UserType: "U",
		MobileNumber: payload.MobileNumber,
	}

	if err := config.Psql.Table("users").Create(user).Error; err != nil {
		return err
	}
	
	return nil
}



