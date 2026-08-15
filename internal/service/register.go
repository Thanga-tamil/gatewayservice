package service

import (
	"errors"
	"gateway/internal/config"
	"gateway/internal/utils"
	log "gateway/internal/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func ValidateRegisterUserPayload(payload *utils.Data) error {

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


func ProcessNewUserRegistration(reqCtx *gin.Context, payload *utils.Data) {
	user := utils.Users{
		UserId: payload.UserId,
		Nickname: payload.NickName,
		IsDeleted: false,
		ContactPermission: payload.ContactPermission,
		CreatedAt: time.Now(),
		EmailId: payload.Email,
		UserType: "U",
		MobileNumber: payload.MobileNumber,
	}

	if err := config.Psql.Create(user).Error; err != nil {
		log.Error(err.Error())
	}
}



