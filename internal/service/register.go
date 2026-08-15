package service

import (
	"errors"
	"gateway/internal/utils"
	"strings"

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

}
