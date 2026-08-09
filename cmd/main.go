package main

import (
	"gateway/internal/app"
	"gateway/internal/router"
	"gateway/internal/middleware"
	log "gateway/internal/utils"

)

func main(){
	// - /register API .. This API should only be used one time for a account.
	log.NewLogger("gatway.log")

	serve := app.App()
	
	serve.Use(middleware.Cors(), middleware.ReqFilter())

	v1Group := serve.Group("/api/v1")

	router.Route(v1Group)

	app.Run(serve, "0.0.0.0:1212")
}
