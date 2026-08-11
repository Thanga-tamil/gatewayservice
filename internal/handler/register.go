package handler

import (
	"fmt"
	"gateway/internal/encrypt"

	"github.com/gin-gonic/gin"
)

type user struct{
	UserId string
	Name string
	mobileNumber int64
}

func RegisterUser(ginCtx *gin.Context){
	var User user
	if err := ginCtx.ShouldBindJSON(&User); err != nil {
		panic(err)
	}
	fmt.Println("user: ", User)
	decryptedData, err := encrypt.AesDecrypt("", "", "")
	if err != nil {
		panic(err)
	}

	fmt.Println("decryptedData: ", string(decryptedData))
}

