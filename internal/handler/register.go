package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gateway/internal/service"
	model "gateway/internal/utils"
)

func RegisterUser(ctx *gin.Context) {
	var registerUser model.Registeruser

	ctx.ShouldBindBodyWithJSON(&registerUser)

	userPayload, err := service.BindRequest(&registerUser)
	if err != nil {
		processErr(ctx, err); return
	}

	if err := service.ValidateRegisterUserPayload(userPayload); err != nil {
		processErr(ctx, err); return
	}

	if err := service.ProcessNewUserRegistration(ctx, userPayload); err != nil {
		processErr(ctx, err); return
	}

	token, err := service.GenerateAccessToken(userPayload.UserId)
	if err != nil {
		processErr(ctx, err); return
	}

	res := map[string]string{"accessToken": token}
	response := model.Success(http.StatusOK, res, "User Registration Success")
	ctx.JSON(http.StatusOK, response)
}

func processErr(ctx *gin.Context, err error) {
	response := model.Bad(http.StatusBadRequest, err.Error())
	ctx.JSON(http.StatusBadRequest, response)
}
